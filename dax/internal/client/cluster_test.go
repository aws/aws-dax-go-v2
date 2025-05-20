/*
  Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.

  Licensed under the Apache License, Version 2.0 (the "License").
  You may not use this file except in compliance with the License.
  A copy of the License is located at

      http://www.apache.org/licenses/LICENSE-2.0

  or in the "license" file accompanying this file. This file is distributed
  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
  express or implied. See the License for the specific language governing
  permissions and limitations under the License.
*/

package client

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	daxTypes "github.com/aws/aws-dax-go-v2/dax/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/metrics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testRetryer struct {
	Attemps     int
	Duration    time.Duration
	CalledCount int
}

var _ aws.Retryer = (*testRetryer)(nil)

func (r *testRetryer) IsErrorRetryable(_ error) bool {
	return true
}
func (r *testRetryer) MaxAttempts() int {
	return r.Attemps
}

func (r *testRetryer) RetryDelay(_ int, _ error) (time.Duration, error) {
	r.CalledCount++
	return r.Duration, nil
}
func (r *testRetryer) GetRetryToken(ctx context.Context, opErr error) (releaseToken func(error) error, err error) {
	return releaseToken, nil
}
func (r *testRetryer) GetInitialToken() (releaseToken func(error) error) {
	return releaseToken
}

func testTaskExecutor(t *testing.T) { // disabled as test is time sensitive
	executor := newExecutor()

	var cnt1, cnt2, cnt3 int32
	executor.start(10*time.Millisecond, func() error {
		atomic.AddInt32(&cnt1, 1)
		return nil
	})
	executor.start(20*time.Millisecond, func() error {
		atomic.AddInt32(&cnt2, 1)
		return nil
	})
	executor.start(50*time.Millisecond, func() error {
		atomic.AddInt32(&cnt3, 1)
		return nil
	})
	<-time.After(105 * time.Millisecond)
	if c := atomic.LoadInt32(&cnt1); c != 10 {
		t.Errorf("expected 10, got %d", c)
	}
	if c := atomic.LoadInt32(&cnt2); c != 5 {
		t.Errorf("expected 10, got %d", c)
	}
	if c := atomic.LoadInt32(&cnt3); c != 2 {
		t.Errorf("expected 10, got %d", c)
	}
	if executor.numTasks() != 3 {
		t.Errorf("expected 3, got %d", executor.numTasks())
	}
	executor.stopAll()
	<-time.After(105 * time.Millisecond)
	if c := atomic.LoadInt32(&cnt1); c != 10 {
		t.Errorf("expected 10, got %d", c)
	}
	if c := atomic.LoadInt32(&cnt2); c != 5 {
		t.Errorf("expected 10, got %d", c)
	}
	if c := atomic.LoadInt32(&cnt3); c != 2 {
		t.Errorf("expected 10, got %d", c)
	}
	if executor.numTasks() != 0 {
		t.Errorf("expected 0, got %d", executor.numTasks())
	}
}

func TestClusterDaxClient_retry(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	retries := 3
	for successfulAttempt := 1; successfulAttempt < retries+5; successfulAttempt++ {
		calls := 0
		action := func(client DaxAPI, o RequestOptions) error {
			calls++
			if calls == successfulAttempt {
				return nil
			}
			// Return a retryable error (code 1 is retryable according to DaxRetryer implementation)
			return newDaxRequestFailure([]int{1}, "RetryableError", "", "", 500, smithy.FaultServer)
		}

		opt := RequestOptions{
			Options: dynamodb.Options{
				RetryMaxAttempts: retries,
			},
			Retryer: DaxRetryer{
				BaseThrottleDelay: time.Millisecond, // Use small delay for testing
				MaxBackoffDelay:   time.Millisecond * 10,
			},
		}

		err := cc.retry(context.Background(), "op", action, opt)
		maxAttempts := retries + 1
		if successfulAttempt <= maxAttempts {
			if calls != successfulAttempt {
				t.Errorf("expected success on %d call, but made %d calls", successfulAttempt, calls)
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		} else {
			expectedCalls := retries + 1
			if calls != expectedCalls {
				t.Errorf("expected %d retries, but made %d", expectedCalls, calls)
			}
			if err == nil {
				t.Errorf("expected error for attempt %d (max attempts: %d)", successfulAttempt, maxAttempts)
			}
		}

		// Add small sleep to ensure we don't overwhelm the system during tests
		time.Sleep(time.Millisecond)
	}
}

func TestClusterDaxClient_retrySleepCycleCount(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	// Track retry delays
	var delays []time.Time

	action := func(client DaxAPI, o RequestOptions) error {
		delays = append(delays, time.Now())
		return &types.ProvisionedThroughputExceededException{
			Message: aws.String("The request rate for the table exceeds the maximum allowed throughput."),
		}
	}

	// Test with no retries
	opt := RequestOptions{
		Retryer: DaxRetryer{
			BaseThrottleDelay: time.Millisecond,
			MaxBackoffDelay:   time.Millisecond * 10,
		},
	}

	delays = nil
	err := cc.retry(context.Background(), "op", action, opt)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	// Should only have one attempt with no retries
	if len(delays) != 1 {
		t.Fatalf("Expected 1 attempt, got %d", len(delays))
	}

	// Test with retries
	delays = nil
	opt = RequestOptions{
		Options: dynamodb.Options{
			RetryMaxAttempts: 3,
		},
		Retryer: DaxRetryer{
			BaseThrottleDelay: time.Millisecond,
			MaxBackoffDelay:   time.Millisecond * 10,
		},
	}

	err = cc.retry(context.Background(), "op", action, opt)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	// Should have initial attempt plus retries
	expectedAttempts := opt.Options.RetryMaxAttempts + 1
	if len(delays) != expectedAttempts {
		t.Fatalf("Expected %d attempts, got %d", expectedAttempts, len(delays))
	}

	// Verify that there were delays between attempts
	for i := 1; i < len(delays); i++ {
		delay := delays[i].Sub(delays[i-1])
		if delay < opt.Retryer.BaseThrottleDelay {
			t.Errorf("Delay between attempts %d and %d was too short: %v",
				i-1, i, delay)
		}
	}

	// Test that delays increase with each retry
	var lastDelay time.Duration
	for i := 1; i < len(delays); i++ {
		currentDelay := delays[i].Sub(delays[i-1])
		if i > 1 && currentDelay < lastDelay {
			t.Logf("Warning: Expected increasing delays, got %v after %v",
				currentDelay, lastDelay)
		}
		lastDelay = currentDelay
	}
}

func TestClusterDaxClient_throttleRetry(t *testing.T) {
	cluster, _ := newTestCluster([]string{"*********:8111"})
	cluster.update([]serviceEndpoint{{hostname: "*********", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	attempts := 0
	action := func(client DaxAPI, o RequestOptions) error {
		attempts++
		if attempts <= 2 { // First two calls return throttle error
			return &types.ProvisionedThroughputExceededException{
				Message: aws.String("Throttled request"),
			}
		}
		return nil // Success on third attempt
	}

	opt := RequestOptions{
		Options: dynamodb.Options{
			RetryMaxAttempts: 3,
		},
		Retryer: DaxRetryer{
			BaseThrottleDelay: time.Millisecond,
			MaxBackoffDelay:   time.Millisecond * 10,
		},
	}

	err := cc.retry(context.Background(), "op", action, opt)
	if err != nil {
		t.Fatalf("Expected success after retries, got error: %v", err)
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

func TestClusterDaxClient_retryReturnsLastError(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	callCount := 0
	action := func(client DaxAPI, o RequestOptions) error {
		callCount++
		return fmt.Errorf("Error_%d", callCount)
	}

	opt := RequestOptions{}
	opt.RetryMaxAttempts = 2

	err := cc.retry(context.Background(), "op", action, opt)
	expectedError := fmt.Errorf("Error_%d", callCount)
	if err.Error() != expectedError.Error() {
		t.Fatalf("Wrong error. Expected %v, but got %v", expectedError, err)
	}
}

func TestClusterDaxClient_retryReturnsCorrectErrorType(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8111"})
	cluster.update([]serviceEndpoint{{hostname: "localhost", port: 8121}})
	cc := ClusterDaxClient{config: DefaultConfig(), cluster: cluster}

	message := "Message"
	statusCode := 0
	requestID := "RequestID"
	defaultErrCode := "empty"

	cases := []struct {
		// input
		codes []int

		// output
		errCode string
		class   reflect.Type
	}{
		{
			codes:   []int{4, 23, 24},
			errCode: (&types.ResourceNotFoundException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ResourceNotFoundException{}),
		},
		{
			codes:   []int{4, 23, 35},
			errCode: (&types.ResourceInUseException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ResourceInUseException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 40},
			errCode: (&types.ProvisionedThroughputExceededException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ProvisionedThroughputExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 40},
			errCode: (&types.ProvisionedThroughputExceededException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ProvisionedThroughputExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 41},
			errCode: (&types.ResourceNotFoundException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ResourceNotFoundException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 43},
			errCode: (&types.ConditionalCheckFailedException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ConditionalCheckFailedException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 45},
			errCode: (&types.ResourceInUseException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ResourceInUseException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 46},
			errCode: ErrCodeValidationException,
			class:   reflect.TypeOf(&smithy.GenericAPIError{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 47},
			errCode: (&types.InternalServerError{}).ErrorCode(),
			class:   reflect.TypeOf(&types.InternalServerError{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 48},
			errCode: (&types.ItemCollectionSizeLimitExceededException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.ItemCollectionSizeLimitExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 49},
			errCode: (&types.LimitExceededException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.LimitExceededException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 50},
			errCode: ErrCodeThrottlingException,
			class:   reflect.TypeOf(&smithy.GenericAPIError{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 57},
			errCode: (&types.TransactionConflictException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.TransactionConflictException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 58},
			errCode: (&types.TransactionCanceledException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.TransactionCanceledException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 59},
			errCode: (&types.TransactionInProgressException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.TransactionInProgressException{}),
		},
		{
			codes:   []int{4, 37, 38, 39, 60},
			errCode: (&types.IdempotentParameterMismatchException{}).ErrorCode(),
			class:   reflect.TypeOf(&types.IdempotentParameterMismatchException{}),
		},
		{
			codes:   []int{4, 37, 38, 44},
			errCode: ErrCodeNotImplemented,
			class:   reflect.TypeOf(&smithy.GenericAPIError{}),
		},
	}

	for _, c := range cases {
		action := func(client DaxAPI, o RequestOptions) error {
			if c.errCode == (&types.TransactionCanceledException{}).ErrorCode() {
				return newDaxTransactionCanceledFailure(c.codes, defaultErrCode, message, requestID, statusCode, nil, nil, nil)
			}
			return newDaxRequestFailure(c.codes, defaultErrCode, message, requestID, statusCode, smithy.FaultServer)
		}

		opt := RequestOptions{}

		err := cc.retry(context.Background(), "op", action, opt)
		actualClass := reflect.TypeOf(err)
		if actualClass != c.class {
			t.Errorf("conversion of code sequence %v failed: expected %s, but got %s", c.codes, c.class.String(), actualClass.String())
		}
		f, _ := err.(smithy.APIError)
		require.NotNilf(t, f, "conversion of code sequence %v failed: expected implement smithy.APIError", c.codes)
		assert.Equal(t, c.errCode, f.ErrorCode())
	}
}

func TestCluster_parseHostPorts(t *testing.T) {
	endpoints := []string{"dax.us-east-1.amazonaws.com:8111"}
	hostPorts, _, _, err := getHostPorts(endpoints)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(hostPorts) != len(endpoints) {
		t.Errorf("expected %v, got %v", len(endpoints), len(hostPorts))
	}
	if hostPorts[0].host != "dax.us-east-1.amazonaws.com" {
		t.Errorf("expected %v, got %v", "dax.us-east-1.amazonaws.com", hostPorts[0].host)
	}
	if hostPorts[0].port != 8111 {
		t.Errorf("expected %v, got %v", 8111, hostPorts[0].port)
	}
}

func TestCluster_pullFromNextSeed(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"non-existent-host:8888", "127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8121}})

	if err := cluster.refresh(false); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}
	client := clientBuilder.clients[0]
	assertDiscoveryClient(client, t)
	assertActiveClient(clientBuilder.clients[1], t)
	expected := hostPort{"127.0.0.1", 8111}
	if expected != client.hp {
		t.Errorf("expected %v, got %v", expected, client.hp)
	}
}

func TestCluster_refreshEmpty(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{})

	if err := cluster.refresh(false); err != nil {
		t.Errorf("unexpected error %v", err)
	}

	assertNumRoutes(cluster, 0, t)
	if _, err := cluster.client(nil, "op"); err == nil {
		t.Errorf("expected err, got nil")
	}
	if len(clientBuilder.clients) != 1 {
		t.Errorf("expected 1, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)
}

func TestCluster_refreshThreshold(t *testing.T) {
	cfg := DefaultConfig()
	cfg.ClusterUpdateThreshold = time.Millisecond * 100
	cfg.HostPorts = []string{"127.0.0.1:8111"}
	cfg.Region = "us-west-2"

	cluster, clientBuilder := newTestClusterWithConfig(cfg)
	for i := 0; i < 10; i++ {
		if err := cluster.refresh(false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	if 1 != len(clientBuilder.clients) {
		t.Errorf("expected 1, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)

	<-time.After(cfg.ClusterUpdateThreshold)
	for i := 0; i < 10; i++ {
		if err := cluster.refresh(false); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	if 2 != len(clientBuilder.clients) {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[1], t)
}

func TestCluster_refreshDup(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8121}})

	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %v", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)
	assertActiveClient(clientBuilder.clients[1], t)

	oldActive := cluster.active
	oldRoutes := cluster.getAllRoutes()
	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if fmt.Sprintf("%p", cluster.active) != fmt.Sprintf("%p", oldActive) {
		t.Errorf("unexpected updation to active")
	}
	if fmt.Sprintf("%p", cluster.getAllRoutes()) != fmt.Sprintf("%p", oldRoutes) {
		t.Errorf("unexpected updation to routes")
	}
	if len(clientBuilder.clients) != 3 {
		t.Errorf("expected 3, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[2], t)
}

func TestCluster_refreshUpdate(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8111"})

	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8121}})

	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}
	assertDiscoveryClient(clientBuilder.clients[0], t)
	assertActiveClient(clientBuilder.clients[1], t)

	setExpectation(cluster, []serviceEndpoint{
		{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8121},
		{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8122}})
	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 2, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if len(clientBuilder.clients) != 4 {
		t.Errorf("expected 4, got %d", len(clientBuilder.clients))
	}

	assertDiscoveryClient(clientBuilder.clients[2], t)
	assertActiveClient(clientBuilder.clients[3], t)
}

func TestCluster_update(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8888"})

	first := []serviceEndpoint{{hostname: "localhost", port: 8121}}
	if !cluster.hasChanged(first) {
		t.Errorf("expected config change")
	}
	cluster.update(first)
	assertNumRoutes(cluster, 1, t)
	assertConnections(cluster, first, t)
	assertHealthCheckCalls(cluster, t)

	// add new hosts
	second := []serviceEndpoint{
		{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8121},
		{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8122},
		{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8123}}
	if !cluster.hasChanged(second) {
		t.Errorf("expected config change")
	}
	cluster.update(second)
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, second, t)
	assertHealthCheckCalls(cluster, t)

	// replace host
	third := []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8124}}
	if !cluster.hasChanged(third) {
		t.Errorf("expected config change")
	}
	cluster.update(third)
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, third, t)
	assertHealthCheckCalls(cluster, t)

	// remove host
	fourth := []serviceEndpoint{{hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8124}}
	if !cluster.hasChanged(fourth) {
		t.Errorf("expected config change")
	}
	cluster.update(fourth)
	assertNumRoutes(cluster, 2, t)
	assertConnections(cluster, fourth, t)
	assertHealthCheckCalls(cluster, t)

	// no change
	fifth := []serviceEndpoint{{hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8124}}
	if cluster.hasChanged(fifth) {
		t.Errorf("unexpected config change")
	}
	cluster.update(fifth)
	assertNumRoutes(cluster, 2, t)
	assertConnections(cluster, fifth, t)
	assertHealthCheckCalls(cluster, t)
}

func TestCluster_onHealthCheckFailed(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8888"})
	endpoint := serviceEndpoint{hostname: "localhost", port: 8123}
	first := []serviceEndpoint{endpoint, {hostname: "localhost", port: 8124}, {hostname: "localhost", port: 8125}}
	cluster.update(first)

	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, first, t)
	assertHealthCheckCalls(cluster, t)
	// Replace old instance of client with new one. Total client instances: 3 + 0
	assert.Equal(t, 3, len(clientBuilder.clients))
	assertCloseCalls(cluster, 0, t)

	cluster.onHealthCheckFailed(endpoint.hostPort())
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, first, t)
	assertHealthCheckCalls(cluster, t)
	// Replace old instance of client with new one. Total client instances: 3 + 1
	assert.Equal(t, 4, len(clientBuilder.clients))
	assertCloseCalls(cluster, 1, t)

	// Another failure
	cluster.onHealthCheckFailed(endpoint.hostPort())
	assertNumRoutes(cluster, 3, t)
	assertConnections(cluster, first, t)
	assertHealthCheckCalls(cluster, t)
	// Replace old instance of client with new one. Total client instances: 3 + 2
	assert.Equal(t, 5, len(clientBuilder.clients))
	assertCloseCalls(cluster, 2, t)
}

func TestCluster_client(t *testing.T) {
	cluster, _ := newTestCluster([]string{"127.0.0.1:8888"})
	endpoints := []serviceEndpoint{{hostname: "localhost", port: 8121}, {hostname: "localhost", port: 8122}, {hostname: "localhost", port: 8123}}

	cluster.update(endpoints)
	assertNumRoutes(cluster, 3, t)
	prev, err := cluster.client(nil, "op")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	for i := 0; i < 100; i++ {
		next, err := cluster.client(prev, "op")
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if next == prev {
			t.Errorf("expected next != prev")
		}
		prev = next
	}
}

func TestCluster_Close(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8121}})

	if err := cluster.refreshNow(); err != nil {
		t.Errorf("unpexected error %v", err)
	}
	assertNumRoutes(cluster, 1, t)
	if _, err := cluster.client(nil, "op"); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(clientBuilder.clients) != 2 {
		t.Errorf("expected 2, got %d", len(clientBuilder.clients))
	}

	cluster.Close()
	for _, c := range clientBuilder.clients {
		if c.closeCalls != 1 {
			t.Errorf("expected 1, got %d", c.closeCalls)
		}
	}
}

func Test_CorrectHostPortUrlFormat(t *testing.T) {
	hostPort := "dax://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com:1234"
	host, port, scheme, _ := parseHostPort(hostPort)
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 1234, port, "")
	assertEqual(t, "dax", scheme, "")
}

func Test_MissingScheme(t *testing.T) {
	hostPort := "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com:8111"
	host, port, scheme, _ := parseHostPort(hostPort)
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 8111, port, "")
	assertEqual(t, "dax", scheme, "")
}

func Test_MissingPortForDax(t *testing.T) {
	hostPort := "dax://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com"
	host, port, scheme, _ := parseHostPort(hostPort)
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 8111, port, "")
	assertEqual(t, "dax", scheme, "")
}

func Test_MissingPortForDaxs(t *testing.T) {
	hostPort := "daxs://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com"
	host, port, scheme, _ := parseHostPort(hostPort)
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 9111, port, "")
	assertEqual(t, "daxs", scheme, "")
}

func Test_UnsupportedScheme(t *testing.T) {
	hostPort := "sample://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com"
	_, _, _, err := parseHostPort(hostPort)
	assertEqual(t, reflect.TypeOf(err), reflect.TypeOf(&smithy.GenericAPIError{}), "")
}

func Test_DaxsCorrectUrlFormat(t *testing.T) {
	hostPort := "daxs://test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com:1234"
	host, port, scheme, _ := parseHostPort(hostPort)
	assertEqual(t, "test.nds.clustercfg.dax.usw2integ.cache.amazonaws.com", host, "")
	assertEqual(t, 1234, port, "")
	assertEqual(t, "daxs", scheme, "")
}

var nonEncEp = "dax://cluster.random.alpha-dax-clusters.us-east-1.amazonaws.com"
var nonEncNodeEp = "cluster-a.random.nodes.alpha-dax-clusters.us-east-1.amazonaws.com:8111"
var encEp = "daxs://cluster2.random.alpha-dax-clusters.us-east-1.amazonaws.com"
var encNodeEp = "daxs://cluster2-a.random.nodes.alpha-dax-clusters.us-east-1.amazonaws.com:9111"

func Test_InconsistentScheme(t *testing.T) {
	_, _, _, err := getHostPorts([]string{nonEncEp, encEp})
	assertEqual(t, reflect.TypeOf(err), reflect.TypeOf(&smithy.GenericAPIError{}), "")
}

func Test_MultipleUnEncryptedEndpoints(t *testing.T) {
	hps, _, _, _ := getHostPorts([]string{nonEncEp, nonEncNodeEp})
	assert.Contains(t, hps, hostPort{"cluster.random.alpha-dax-clusters.us-east-1.amazonaws.com", 8111})
	assert.Contains(t, hps, hostPort{"cluster-a.random.nodes.alpha-dax-clusters.us-east-1.amazonaws.com", 8111})
}

func Test_MultipleEncryptedEndpoints(t *testing.T) {
	_, _, _, err := getHostPorts([]string{encEp, encNodeEp})
	assertEqual(t, reflect.TypeOf(err), reflect.TypeOf(&smithy.GenericAPIError{}), "")
}

func TestCluster_RouteManagerDisabled(t *testing.T) {
	cluster, clientBuilder := newTestCluster([]string{"non-existent-host:8888", "127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8121}})

	if cluster.isRouteManagerEnabled() {
		t.Errorf("Route manager should be disabled!")
	}

	oldRoutes := cluster.getAllRoutes()
	route, _ := clientBuilder.newClient(net.IP{}, 8111, connConfig{}, "dummy", nil, 10, nil, nil, nil)
	cluster.addRoute("dummy", route)
	newRoutes := cluster.getAllRoutes()

	if len(newRoutes) != len(oldRoutes) {
		t.Errorf("Route added with disabled route manager")
	}

	cluster.removeRoute("dummy", route)
	newRoutes = cluster.getAllRoutes()
	if len(newRoutes) != len(oldRoutes) {
		t.Errorf("Route removed with disabled route manager")
	}
}

func TestCluster_RouteManagerEnabled(t *testing.T) {
	cluster, clientBuilder := newTestClusterWithRouteManagerEnabled([]string{"non-existent-host:8888", "127.0.0.1:8111"})
	setExpectation(cluster, []serviceEndpoint{{hostname: "localhost", address: net.ParseIP("127.0.0.1"), port: 8121}})
	if !cluster.isRouteManagerEnabled() {
		t.Errorf("Route manager should be enabled!")
	}
	oldRoutes := cluster.getAllRoutes()
	route, _ := clientBuilder.newClient(net.IP{}, 8111, connConfig{}, "dummy", nil, 10, nil, nil, nil)
	cluster.addRoute("dummy", route)
	newRoutes := cluster.getAllRoutes()

	if len(newRoutes) != len(oldRoutes)+1 {
		t.Errorf("Route not added with enabled route manager")
	}

	cluster.removeRoute("dummy", route)
	newRoutes = cluster.getAllRoutes()
	if len(newRoutes) != len(oldRoutes) {
		t.Errorf("Route not removed with enabled route manager")
	}
}

func assertConnections(cluster *cluster, endpoints []serviceEndpoint, t *testing.T) {
	if len(cluster.active) != len(endpoints) {
		t.Errorf("expected %d, got %d", len(cluster.active), len(endpoints))
	}
	for _, ep := range endpoints {
		hp := ep.hostPort()
		c, ok := cluster.active[hp]
		if !ok {
			t.Errorf("missing client %v", hp)
		}
		if tc, ok := c.client.(*testClient); ok {
			if tc.hp != hp {
				t.Errorf("expected %v, got %v", hp, tc.hp)
			}
		}
	}
	return
}

func assertNumRoutes(cluster *cluster, num int, t *testing.T) {
	t.Helper()
	if len(cluster.active) != num {
		t.Errorf("expected %d, got %d", num, len(cluster.active))
	}
	if len(cluster.getAllRoutes()) != num {
		t.Errorf("expected %d, got %d", num, len(cluster.getAllRoutes()))
	}
}

func assertHealthCheckCalls(cluster *cluster, t *testing.T) {
	t.Helper()
	for _, cliAndCfg := range cluster.active {
		healtCheckCalls := cliAndCfg.client.(*testClient).healthCheckCalls
		if healtCheckCalls != 1 {
			t.Errorf("expected 1 healthcheck call, got %d", healtCheckCalls)
		}
	}
}

func assertCloseCalls(cluster *cluster, num int, t *testing.T) {
	t.Helper()
	cnt := 0
	for _, client := range cluster.clientBuilder.(*testClientBuilder).clients {
		if client.closeCalls == 1 {
			cnt++
		}
	}
	assert.Equal(t, num, cnt)
}

func assertDiscoveryClient(client *testClient, t *testing.T) {
	t.Helper()
	if client.endpointsCalls != 1 {
		t.Errorf("expected 1, got %d", client.endpointsCalls)
	}
	if client.closeCalls != 1 {
		t.Errorf("expected 1, got %d", client.closeCalls)
	}
}

func assertActiveClient(client *testClient, t *testing.T) {
	t.Helper()
	if client.endpointsCalls != 0 {
		t.Errorf("expected 0, got %d", client.endpointsCalls)
	}
	if client.closeCalls != 0 {
		t.Errorf("expected 0, got %d", client.closeCalls)
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	t.Helper()
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func newTestCluster(seeds []string) (*cluster, *testClientBuilder) {
	cfg := DefaultConfig()
	cfg.HostPorts = seeds
	cfg.Region = "us-west-2"
	return newTestClusterWithConfig(cfg)
}

func newTestClusterWithRouteManagerEnabled(seeds []string) (*cluster, *testClientBuilder) {
	cfg := DefaultConfig()
	cfg.HostPorts = seeds
	cfg.Region = "us-west-2"
	cfg.RouteManagerEnabled = true
	return newTestClusterWithConfig(cfg)
}

func newTestClusterWithConfig(config Config) (*cluster, *testClientBuilder) {
	cluster, _ := newCluster(config)
	if cluster == nil {
		return nil, nil
	} else {
		b := &testClientBuilder{}
		cluster.clientBuilder = b
		return cluster, b
	}
}

func setExpectation(cluster *cluster, ep []serviceEndpoint) {
	cluster.clientBuilder.(*testClientBuilder).ep = ep
}

func TestCluster_customDialer(t *testing.T) {
	ours, theirs := net.Pipe()
	var wg sync.WaitGroup
	var result []byte
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			buf := make([]byte, 4096)
			n, _ := ours.Read(buf)
			result = buf[:n]
			ours.Close()
			return
		}
	}()

	var dialContextFn dialContext = func(ctx context.Context, address string, network string) (net.Conn, error) {
		return theirs, nil
	}
	cfg := Config{
		MaxPendingConnectionsPerHost: 1,
		ClusterUpdateInterval:        1 * time.Second,
		Credentials:                  &testCredentialProvider{},
		DialContext:                  dialContextFn,
		Region:                       "us-west-2",
		HostPorts:                    []string{"localhost:9121"},
		logger:                       &logging.Nop{},
		IdleConnectionReapDelay:      30 * time.Second,
		MeterProvider:                &metrics.NopMeterProvider{},
	}
	cc, err := New(cfg)
	require.NoError(t, err)
	cc.GetItemWithOptions(context.Background(), &dynamodb.GetItemInput{TableName: aws.String("MyTable")}, &dynamodb.GetItemOutput{}, RequestOptions{})

	wg.Wait()

	assert.Equal(t, magic, string(result[1:8]), "expected the ClusterClient to write to the connection provided by the custom dialer")
}

func TestCluster_filterAndSelectAddress(t *testing.T) {
	ipv4Seeds := []net.IP{
		net.ParseIP("192.0.2.1"),
		net.ParseIP("198.51.100.42"),
		net.ParseIP("203.0.113.17"),
		net.ParseIP("198.51.100.88"),
		net.ParseIP("192.0.2.99"),
	}

	ipv6Seeds := []net.IP{
		net.ParseIP("2001:db8::1"),
		net.ParseIP("2001:db8:0:1::abcd"),
		net.ParseIP("2001:db8:abcd:0012::1"),
		net.ParseIP("2001:db8:ffff::42"),
		net.ParseIP("2001:db8::dead:beef"),
	}

	dual_stack_seeds := append(ipv4Seeds, ipv6Seeds...)

	// Define a struct to encapsulate the config, that embeds IpDiscovery, and the desired seeds used for determining cluster type
	type testInput struct {
		IpDiscovery daxTypes.IpDiscovery
		seeds       []net.IP
	}

	apiErrIpDiscoveryIPv6 := smithy.GenericAPIError{
		Code:    ErrCodeValidationException,
		Message: "ipDiscovery ipv6 does not match the SupportedNetworkType ipv4.",
		Fault:   smithy.FaultClient,
	}

	apiErrIpDiscoveryIPv4 := apiErrIpDiscoveryIPv6
	apiErrIpDiscoveryIPv4.Message = "ipDiscovery ipv4 does not match the SupportedNetworkType ipv6."

	// Define test cases inputs and expected outputs
	cases := []struct {
		clusterConfigType string
		input             testInput
		output            []net.IP
		err               error
	}{
		{"IPv4", testInput{"", ipv4Seeds}, ipv4Seeds, nil},
		{"IPv4", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv4, seeds: ipv4Seeds}, ipv4Seeds, nil},
		{"IPv4", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv6, seeds: ipv4Seeds}, nil, &apiErrIpDiscoveryIPv6},
		{"IPv6", testInput{IpDiscovery: "", seeds: ipv6Seeds}, ipv6Seeds, nil},
		{"IPv6", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv6, seeds: ipv6Seeds}, ipv6Seeds, nil},
		{"IPv6", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv4, seeds: ipv6Seeds}, nil, &apiErrIpDiscoveryIPv4},
		{"Dual stack", testInput{IpDiscovery: "", seeds: dual_stack_seeds}, ipv4Seeds, nil},
		{"Dual stack", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv4, seeds: dual_stack_seeds}, ipv4Seeds, nil},
		{"Dual stack", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv6, seeds: dual_stack_seeds}, ipv6Seeds, nil},
	}

	clusterConfig := DefaultConfig()
	clusterConfig.Region = "us-east-1"

	// Execute a test for each case
	for _, testCase := range cases {
		// Define a convetional name for each test case

		testScenario := " ConfigType: " + testCase.clusterConfigType + " IpDiscovery: " + testCase.input.IpDiscovery.String()

		// Run the test
		t.Run(testScenario, func(t *testing.T) {
			// Set clusterConfig fields with data specific to the current test case
			clusterConfig.IpDiscovery = testCase.input.IpDiscovery
			// HostPorts must be set in clusterConfig, otherwise the config is not valid in cfg.validate()
			for _, ip := range testCase.input.seeds {
				clusterConfig.HostPorts = append(clusterConfig.HostPorts, ip.String()+":8111")
			}
			// Create a cluster instance using the config specific to the current test case
			cluster, _ := newCluster(clusterConfig)

			if cluster == nil {
				t.Errorf("Couldn't create cluster instance")
			} else {
				expectedOutput, err := filterAndSelectAddress(testCase.input.seeds, cluster.IpDiscovery)

				assert.Equal(t, expectedOutput, testCase.output)
				assert.Equal(t, err, testCase.err)
			}
		})
	}
}

func TestCluster_filterAndSelectEndpointAddress(t *testing.T) {
	const (
		nodeNo int64  = 1708865638
		portNo int    = 8111
		region string = "us-east-1"
	)

	createServiceEndpointFromIP := func(ip string) serviceEndpoint {
		return serviceEndpoint{nodeId: nodeNo, address: net.ParseIP(ip), port: portNo, availabilityZone: region}
	}

	ipv4Endpoints := []serviceEndpoint{
		createServiceEndpointFromIP("192.0.2.1"),
		createServiceEndpointFromIP("198.51.100.42"),
		createServiceEndpointFromIP("203.0.113.17"),
		createServiceEndpointFromIP("198.51.100.88"),
		createServiceEndpointFromIP("192.0.2.99"),
	}

	ipv6Endpoints := []serviceEndpoint{
		createServiceEndpointFromIP("2001:db8::1"),
		createServiceEndpointFromIP("2001:db8:0:1::abcd"),
		createServiceEndpointFromIP("2001:db8:abcd:0012::1"),
		createServiceEndpointFromIP("2001:db8:ffff::42"),
		createServiceEndpointFromIP("2001:db8::dead:beef"),
	}

	dual_stack_endpoints := append(ipv4Endpoints, ipv6Endpoints...)

	// Define a struct to encapsulate the config, that embeds IpDiscovery, and the desired seeds used for determining cluster type
	type testInput struct {
		IpDiscovery daxTypes.IpDiscovery
		endpoints   []serviceEndpoint
	}

	apiErrIpDiscoveryIPv6 := smithy.GenericAPIError{
		Code:    ErrCodeValidationException,
		Message: "ipDiscovery ipv6 does not match the SupportedNetworkType ipv4.",
		Fault:   smithy.FaultClient,
	}

	apiErrIpDiscoveryIPv4 := apiErrIpDiscoveryIPv6
	apiErrIpDiscoveryIPv4.Message = "ipDiscovery ipv4 does not match the SupportedNetworkType ipv6."

	// Define test cases inputs and expected outputs
	cases := []struct {
		clusterConfigType string
		input             testInput
		output            []serviceEndpoint
		err               error
	}{
		{"IPv4", testInput{"", ipv4Endpoints}, ipv4Endpoints, nil},
		{"IPv4", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv4, endpoints: ipv4Endpoints}, ipv4Endpoints, nil},
		{"IPv4", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv6, endpoints: ipv4Endpoints}, nil, &apiErrIpDiscoveryIPv6},
		{"IPv6", testInput{IpDiscovery: "", endpoints: ipv6Endpoints}, ipv6Endpoints, nil},
		{"IPv6", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv6, endpoints: ipv6Endpoints}, ipv6Endpoints, nil},
		{"IPv6", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv4, endpoints: ipv6Endpoints}, nil, &apiErrIpDiscoveryIPv4},
		{"Dual stack", testInput{IpDiscovery: "", endpoints: dual_stack_endpoints}, ipv4Endpoints, nil},
		{"Dual stack", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv4, endpoints: dual_stack_endpoints}, ipv4Endpoints, nil},
		{"Dual stack", testInput{IpDiscovery: daxTypes.IpDiscoveryIPv6, endpoints: dual_stack_endpoints}, ipv6Endpoints, nil},
	}

	clusterConfig := DefaultConfig()
	clusterConfig.Region = "us-east-1"

	// Execute a test for each case
	for _, testCase := range cases {
		// Define a convetional name for each test case

		testScenario := " ConfigType: " + testCase.clusterConfigType + " IpDiscovery: " + testCase.input.IpDiscovery.String()

		// Run the test
		t.Run(testScenario, func(t *testing.T) {
			// Set clusterConfig fields with data specific to the current test case
			clusterConfig.IpDiscovery = testCase.input.IpDiscovery
			// HostPorts must be set in clusterConfig, otherwise the config is not valid in cfg.validate()
			for _, endpoint := range testCase.input.endpoints {
				clusterConfig.HostPorts = append(clusterConfig.HostPorts, (net.IP(endpoint.address)).String()+":"+strconv.Itoa(endpoint.port))
			}
			// Create a cluster instance using the config specific to the current test case
			cluster, _ := newCluster(clusterConfig)

			if cluster == nil {
				t.Errorf("Couldn't create cluster instance")
			} else {
				expectedOutput, err := filterAndSelectAddress(testCase.input.endpoints, cluster.IpDiscovery)

				assert.Equal(t, expectedOutput, testCase.output)
				assert.Equal(t, err, testCase.err)
			}
		})
	}
}

func TestCluster_refreshNowForIpDiscovery(t *testing.T) {
	ipv4ClusterSeeds := []string{
		"192.0.2.1:8080",
		"198.51.100.42:8080",
		"203.0.113.17:8080"}

	expectedEndpointsIPv4 := []serviceEndpoint{
		{hostname: "localhost", address: net.ParseIP("192.0.2.1"), port: 8121},
		{hostname: "localhost", address: net.ParseIP("198.51.100.42"), port: 8122},
		{hostname: "localhost", address: net.ParseIP("203.0.113.17"), port: 8123}}

	ipv6ClusterSeeds := []string{
		"2001:db8::1:8081",
		"2001:db8:0:1::abcd:8081",
		"2001:db8:abcd:0012::1:8081",
		"2600:1f18:3782:de01:11ed:c16a:5644:4d2e:8081",
	}

	expectedEndpointsIPv6 := []serviceEndpoint{
		{hostname: "localhost", address: net.ParseIP("2001:db8::1"), port: 8124},
		{hostname: "localhost", address: net.ParseIP("2001:db8:0:1::abcd"), port: 8125},
		{hostname: "localhost", address: net.ParseIP("2001:db8:abcd:0012::1"), port: 8126},
		{hostname: "localhost", address: net.ParseIP("2600:1f18:3782:de01:11ed:c16a:5644:4d2e"), port: 8127}}

	apiErrIpDiscoveryIPv6 := smithy.GenericAPIError{
		Code:    ErrCodeValidationException,
		Message: "ipDiscovery ipv6 does not match the SupportedNetworkType ipv4.",
		Fault:   smithy.FaultClient,
	}

	apiErrIpDiscoveryIPv4 := apiErrIpDiscoveryIPv6
	apiErrIpDiscoveryIPv4.Message = "ipDiscovery ipv4 does not match the SupportedNetworkType ipv6."

	// Define test cases inputs and expected outputs
	cases := []struct {
		clusterConfigType string
		IpDiscovery       daxTypes.IpDiscovery
		seeds             []string
		expectedEndpoints []serviceEndpoint
		err               error
	}{
		{"IPv4", "", ipv4ClusterSeeds, expectedEndpointsIPv4, nil},
		{"IPv4", daxTypes.IpDiscoveryIPv4, ipv4ClusterSeeds, expectedEndpointsIPv4, nil},
		{"IPv4", daxTypes.IpDiscoveryIPv6, ipv4ClusterSeeds, nil, &apiErrIpDiscoveryIPv6},
		{"IPv6", "", ipv6ClusterSeeds, expectedEndpointsIPv6, nil},
		{"IPv6", daxTypes.IpDiscoveryIPv6, ipv6ClusterSeeds, expectedEndpointsIPv6, nil},
		{"IPv6", daxTypes.IpDiscoveryIPv4, ipv6ClusterSeeds, nil, &apiErrIpDiscoveryIPv4},
		{"Dual stack", "", append(ipv4ClusterSeeds, ipv6ClusterSeeds...), expectedEndpointsIPv4, nil},
		{"Dual stack", daxTypes.IpDiscoveryIPv4, append(ipv4ClusterSeeds, ipv6ClusterSeeds...), expectedEndpointsIPv4, nil},
		{"Dual stack", daxTypes.IpDiscoveryIPv6, append(ipv6ClusterSeeds, ipv4ClusterSeeds...), expectedEndpointsIPv6, nil},
	}

	// Execute a test for each case
	for _, testCase := range cases {
		// Define a convetional name for each test case
		testScenario := " ConfigType: " + testCase.clusterConfigType + " IpDiscovery: " + testCase.IpDiscovery.String()

		//each test case uses a dfistinct config for creating the cluster instance, hence a distinct cluster instance
		clusterConfig := DefaultConfig()
		clusterConfig.Region = "us-west-2"
		clusterConfig.IpDiscovery = testCase.IpDiscovery
		clusterConfig.HostPorts = testCase.seeds
		// Create a cluster instance using the config specific to the current test case
		cluster, clientBuilder := newTestClusterWithConfig(clusterConfig)

		if cluster == nil {
			t.Errorf("Couldn't create cluster instance")
		} else if clientBuilder == nil {
			t.Errorf("couldn't create clientBuilder")
		} else {
			t.Run(testScenario, func(t *testing.T) {
				setExpectation(cluster, testCase.expectedEndpoints)

				err := cluster.refreshNow()

				assert.Equal(t, err, testCase.err)
				assertNumRoutes(cluster, len(testCase.expectedEndpoints), t)

				if len(testCase.expectedEndpoints) > 0 && len(clientBuilder.clients) != len(testCase.expectedEndpoints)+1 {
					t.Errorf("expected %d, got %d", len(testCase.expectedEndpoints)+1, len(clientBuilder.clients))
				} else if len(testCase.expectedEndpoints) == 0 {
					assert.Equal(t, len(testCase.expectedEndpoints), len(clientBuilder.clients))
				}

				assertConnections(cluster, testCase.expectedEndpoints, t)

				for idx := range clientBuilder.clients {
					if idx == 0 {
						assertDiscoveryClient(clientBuilder.clients[idx], t)
					} else {
						assertActiveClient(clientBuilder.clients[idx], t)
					}
				}
			})
		}
	}
}

func getEndPointResolver(url string) aws.EndpointResolverWithOptions {
	return aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: url,
		}, nil
	})
}

type testClientBuilder struct {
	ep      []serviceEndpoint
	clients []*testClient
}

var _ clientBuilder = (*testClientBuilder)(nil)

func (b *testClientBuilder) newClient(ip net.IP, port int, _ connConfig, _ string, _ aws.CredentialsProvider, _ int, _ dialContext, _ RouteListener, _ *daxSdkMetrics) (DaxAPI, error) {
	t := &testClient{ep: b.ep, hp: hostPort{ip.String(), port}}
	b.clients = append(b.clients, []*testClient{t}...)
	return t, nil
}

type testClient struct {
	hp                                           hostPort
	ep                                           []serviceEndpoint
	endpointsCalls, closeCalls, healthCheckCalls int
}

var _ DaxAPI = (*testClient)(nil)

func (c *testClient) startHealthChecks(_ *cluster, _ hostPort) {
	c.healthCheckCalls++
}

func (c *testClient) endpoints(_ context.Context, _ RequestOptions) ([]serviceEndpoint, error) {
	c.endpointsCalls++
	return c.ep, nil
}

func (c *testClient) Close() error {
	c.closeCalls++
	return nil
}

func (c *testClient) PutItemWithOptions(_ context.Context, _ *dynamodb.PutItemInput, _ *dynamodb.PutItemOutput, _ RequestOptions) (*dynamodb.PutItemOutput, error) {
	panic("not implemented")
}

func (c *testClient) DeleteItemWithOptions(_ context.Context, _ *dynamodb.DeleteItemInput, _ *dynamodb.DeleteItemOutput, _ RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	panic("not implemented")
}

func (c *testClient) UpdateItemWithOptions(_ context.Context, _ *dynamodb.UpdateItemInput, _ *dynamodb.UpdateItemOutput, _ RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	panic("not implemented")
}

func (c *testClient) GetItemWithOptions(_ context.Context, _ *dynamodb.GetItemInput, _ *dynamodb.GetItemOutput, _ RequestOptions) (*dynamodb.GetItemOutput, error) {
	panic("not implemented")
}

func (c *testClient) ScanWithOptions(_ context.Context, _ *dynamodb.ScanInput, _ *dynamodb.ScanOutput, _ RequestOptions) (*dynamodb.ScanOutput, error) {
	panic("not implemented")
}

func (c *testClient) QueryWithOptions(_ context.Context, _ *dynamodb.QueryInput, _ *dynamodb.QueryOutput, _ RequestOptions) (*dynamodb.QueryOutput, error) {
	panic("not implemented")
}

func (c *testClient) BatchWriteItemWithOptions(_ context.Context, _ *dynamodb.BatchWriteItemInput, _ *dynamodb.BatchWriteItemOutput, _ RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	panic("not implemented")
}

func (c *testClient) BatchGetItemWithOptions(_ context.Context, _ *dynamodb.BatchGetItemInput, _ *dynamodb.BatchGetItemOutput, _ RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	panic("not implemented")
}

func (c *testClient) TransactWriteItemsWithOptions(_ context.Context, _ *dynamodb.TransactWriteItemsInput, _ *dynamodb.TransactWriteItemsOutput, _ RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("not implemented")
}

func (c *testClient) TransactGetItemsWithOptions(_ context.Context, _ *dynamodb.TransactGetItemsInput, _ *dynamodb.TransactGetItemsOutput, _ RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	panic("not implemented")
}

type testCredentialProvider struct {
}

func (p *testCredentialProvider) Retrieve(_ context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     "id",
		SecretAccessKey: "secret",
		SessionToken:    "token",
	}, nil
}
