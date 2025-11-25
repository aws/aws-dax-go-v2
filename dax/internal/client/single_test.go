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
	"errors"
	"fmt"
	"net"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-dax-go-v2/dax/internal/cbor"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var unEncryptedConnConfig = connConfig{isEncrypted: false}

func TestExecuteErrorHandling(t *testing.T) {

	cases := []struct {
		conn *mockConn
		enc  func(writer *cbor.Writer) error
		dec  func(reader *cbor.Reader) error
		ee   error
		ec   map[string]int
	}{
		{ // write error, discard tube
			&mockConn{we: errors.New("io")},
			nil,
			nil,
			errors.New("io"),
			map[string]int{"Write": 1, "Close": 1},
		},
		{ // encoding error, discard tube
			&mockConn{},
			func(writer *cbor.Writer) error { return errors.New("ser") },
			nil,
			errors.New("ser"),
			map[string]int{"Write": 2, "SetDeadline": 1, "Close": 1},
		},
		{ // read error, discard tube
			&mockConn{re: errors.New("IO")},
			func(writer *cbor.Writer) error { return nil },
			nil,
			errors.New("IO"),
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1, "Close": 1},
		},
		{ // serialization error, discard tube
			&mockConn{rd: []byte{cbor.NegInt}},
			func(writer *cbor.Writer) error { return nil },
			nil,
			&smithy.DeserializationError{Err: fmt.Errorf("cbor: expected major type %d, got %d", cbor.Array, cbor.NegInt)},
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1, "Close": 1},
		},
		{ // decode error, discard tube
			&mockConn{rd: []byte{cbor.Array + 0}},
			func(writer *cbor.Writer) error { return nil },
			func(reader *cbor.Reader) error { return errors.New("IO") },
			errors.New("IO"),
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1, "Close": 1},
		},
		{ // dax error, do not discard tube
			&mockConn{rd: []byte{cbor.Array + 3, cbor.PosInt + 4, cbor.PosInt + 0, cbor.PosInt + 0, cbor.Utf, cbor.Nil}},
			func(writer *cbor.Writer) error { return nil },
			nil,
			newDaxRequestFailure([]int{4, 0, 0}, "", "", "", 400, smithy.FaultServer),
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1},
		},
		{ // no error, do not discard tube
			&mockConn{rd: []byte{cbor.Array + 0}},
			func(writer *cbor.Writer) error { return nil },
			func(reader *cbor.Reader) error { return nil },
			nil,
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1},
		},
	}

	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	for i, c := range cases {
		cli, err := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
			return c.conn, nil
		}, nil, om)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		cli.pool.closeTubeImmediately = true

		err = cli.executeWithContext(context.Background(), OpGetItem, c.enc, c.dec, RequestOptions{})
		if !reflect.DeepEqual(c.ee, err) {
			t.Errorf("case[%d] expected error %v, got error %v", i, c.ee, err)
		}
		if !reflect.DeepEqual(c.ec, c.conn.cc) {
			t.Errorf("case[%d] expected %v calls, got %v", i, c.ec, c.conn.cc)
		}
		cli.Close()
	}

	assert.Len(t, tmp.meters, 1)
	s, ok := tmp.meters[daxMeterScope]
	assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, daxMeterScope))
	assert.NotNil(t, s)
	if !ok || s == nil {
		return
	}

	expectCounters(t, om, map[string]int{
		daxConnectionsCreated:                    6,
		daxConnectionsClosedError:                4,
		fmt.Sprintf(daxOpNameSuccess, OpGetItem): 1,
	})
	expectHistograms(t, om, map[string]int{
		fmt.Sprintf(daxOpNameLatencyUs, OpGetItem): 7,
	})
}

func TestRetryPropagatesContextError(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil, om)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	ctx, cancel := context.WithCancel(context.Background())
	requestOptions := RequestOptions{}
	requestOptions.RetryMaxAttempts = 2

	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error { return nil }

	// Cancel context to fail the execution

	cancel()
	err := client.executeWithRetries(ctx, OpGetItem, requestOptions, writer, reader)

	// Context related error should be returned
	cancelErr, ok := err.(*smithy.CanceledError)
	if !ok {
		t.Fatalf("Error type is not smithy.CanceledError, type is %T", err)
	}

	if cancelErr.Err != context.Canceled {
		t.Errorf("aws error doesn't match expected. %v", cancelErr)
	}

	expectCounters(t, om, map[string]int{
		daxConnectionsClosedError:                0,
		fmt.Sprintf(daxOpNameSuccess, OpGetItem): 0,
		fmt.Sprintf(daxOpNameFailure, OpGetItem): 1,
	})
	expectHistograms(t, om, map[string]int{
		fmt.Sprintf(daxOpNameLatencyUs, OpGetItem): 1,
	})
}

func TestRetryPropagatesOtherErrors(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil, om)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	requestOptions := RequestOptions{
		Options: dynamodb.Options{
			RetryMaxAttempts: 1,
		},
	}
	expectedError := errors.New("IO")

	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error { return expectedError }

	err := client.executeWithRetries(context.Background(), OpGetItem, requestOptions, writer, reader)

	// Should be wrapped as a daxRequestFailure
	daxErr, ok := err.(*daxRequestFailure)
	if !ok {
		t.Fatalf("Error type is not daxRequestFailure. type is %T", err)
	}

	if daxErr == nil {
		t.Fatal("Error is nil")
	}

	// Verify error properties
	if !strings.Contains(daxErr.Error(), expectedError.Error()) {
		t.Errorf("error doesn't contain expected message. got: %v, want to contain: %v",
			daxErr.Error(), expectedError.Error())
	}

	// For unknown errors, the code should be ErrCodeUnknown
	if daxErr.Code != ErrCodeUnknown {
		t.Errorf("Expected error code %s, got %s", ErrCodeUnknown, daxErr.Code)
	}

	// For unknown errors, code sequence should be [0] (unretryable)
	codes := daxErr.CodeSequence()
	expectedCodes := []int{0}
	if !reflect.DeepEqual(codes, expectedCodes) {
		t.Errorf("Expected code sequence %v, got %v", expectedCodes, codes)
	}

	// Verify status code is 400 for unknown errors
	if daxErr.StatusCode() != 400 {
		t.Errorf("Expected status code 400, got %d", daxErr.StatusCode())
	}

	expectCounters(t, om, map[string]int{
		daxConnectionsClosedError:                2,
		daxConnectionsCreated:                    2,
		fmt.Sprintf(daxOpNameSuccess, OpGetItem): 0,
		fmt.Sprintf(daxOpNameFailure, OpGetItem): 2,
	})
	expectHistograms(t, om, map[string]int{
		fmt.Sprintf(daxOpNameLatencyUs, OpGetItem): 2,
	})
}

func TestRetryPropagatesOtherErrorsWithDelay(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil, om)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	requestOptions := RequestOptions{}
	requestOptions.RetryMaxAttempts = 1
	expectedError := errors.New("IO")

	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error {
		// Create a daxRequestFailure that represents an IO error
		return newDaxRequestFailure(
			[]int{2}, // Code 2 indicates recoverable failure
			ErrCodeInternalServerError,
			expectedError.Error(),
			"",  // requestID
			400, // statusCode
			smithy.FaultClient,
		)
	}

	err := client.executeWithRetries(context.Background(), OpGetItem, requestOptions, writer, reader)

	// Check that we get a daxRequestFailure
	daxErr, ok := err.(*daxRequestFailure)
	if !ok {
		t.Fatalf("Error type is not daxRequestFailure. type is %T", err)
	}

	if daxErr == nil {
		t.Fatal("Error is nil")
	}

	// Check the error details
	if !strings.Contains(daxErr.Error(), expectedError.Error()) {
		t.Errorf("error message doesn't contain expected text. got: %v, want to contain: %v",
			daxErr.Error(), expectedError.Error())
	}

	// Verify this is marked as a recoverable error
	if !daxErr.recoverable() {
		t.Error("Expected error to be recoverable")
	}

	assert.Len(t, tmp.meters, 1)
	s, ok := tmp.meters[daxMeterScope]
	assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, daxMeterScope))
	assert.NotNil(t, s)
	if !ok || s == nil {
		return
	}

	expectCounters(t, om, map[string]int{
		daxConnectionsClosedError:                2,
		daxConnectionsCreated:                    2,
		fmt.Sprintf(daxOpNameSuccess, OpGetItem): 0,
		fmt.Sprintf(daxOpNameFailure, OpGetItem): 2,
	})
	expectHistograms(t, om, map[string]int{
		fmt.Sprintf(daxOpNameLatencyUs, OpGetItem): 2,
	})
}

func TestRetrySleepCycleCount(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil, om)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	// Create base options
	requestOptions := RequestOptions{
		Options: dynamodb.Options{},
		Retryer: DaxRetryer{
			BaseThrottleDelay: DefaultBaseRetryDelay,
			MaxBackoffDelay:   DefaultMaxBackoffDelay,
		},
	}

	// Test with no retries
	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error { return errors.New("IO") }

	mockRetryer := DaxRetryer{
		BaseThrottleDelay: time.Millisecond,
		MaxBackoffDelay:   time.Millisecond * 100,
	}
	requestOptions.Retryer = mockRetryer

	err := client.executeWithRetries(context.Background(), OpGetItem, requestOptions, writer, reader)
	if err == nil {
		t.Fatal("Expected error but got nil")
	}

	// Test with retries
	requestOptions.Options.RetryMaxAttempts = 3
	err = client.executeWithRetries(context.Background(), OpGetItem, requestOptions, writer, reader)
	if err == nil {
		t.Fatal("Expected error but got nil")
	}

	assert.Len(t, tmp.meters, 1)
	s, ok := tmp.meters[daxMeterScope]
	assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, daxMeterScope))
	assert.NotNil(t, s)
	if !ok || s == nil {
		return
	}

	expectCounters(t, om, map[string]int{
		daxConnectionsClosedError:                5,
		daxConnectionsCreated:                    5,
		fmt.Sprintf(daxOpNameSuccess, OpGetItem): 0,
		fmt.Sprintf(daxOpNameFailure, OpGetItem): 5,
	})
	expectHistograms(t, om, map[string]int{
		fmt.Sprintf(daxOpNameLatencyUs, OpGetItem): 5,
	})
}

func TestRetryLastError(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil, om)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	var callCount int
	requestOptions := RequestOptions{
		Options: dynamodb.Options{
			RetryMaxAttempts: 2,
		},
		Retryer: DaxRetryer{
			BaseThrottleDelay: DefaultBaseRetryDelay,
			MaxBackoffDelay:   DefaultMaxBackoffDelay,
		},
	}

	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error {
		callCount++
		if callCount < 3 { // Return retryable error for first two calls
			// Return a retryable error (code 2 indicates recoverable failure)
			return newDaxRequestFailure(
				[]int{2},
				ErrCodeThrottlingException,
				"Retryable error",
				"",  // requestID
				500, // statusCode
				smithy.FaultServer,
			)
		}
		// Third call returns the final error (code 0 indicates unretryable error)
		return newDaxRequestFailure(
			[]int{0},
			ErrCodeInternalServerError,
			"Last error",
			"",  // requestID
			400, // statusCode
			smithy.FaultClient,
		)
	}

	err := client.executeWithRetries(context.Background(), OpGetItem, requestOptions, writer, reader)

	// Verify error is not nil
	if err == nil {
		t.Fatal("Expected error but got nil")
	}

	// Verify the error is a daxRequestFailure
	daxErr, ok := err.(*daxRequestFailure)
	if !ok {
		t.Fatalf("Expected daxRequestFailure, got %T", err)
	}

	// Verify error details
	if !strings.Contains(daxErr.Error(), "Last error") {
		t.Errorf("Expected error to contain 'Last error', got: %v", daxErr.Error())
	}

	// Verify the error is not recoverable (code sequence should be [0])
	if daxErr.recoverable() {
		t.Error("Expected final error to be non-recoverable")
	}

	// Verify status code
	if daxErr.StatusCode() != 400 {
		t.Errorf("Expected status code 400, got %d", daxErr.StatusCode())
	}

	// Verify retry count
	expectedCalls := 3 // Initial attempt + 2 retries
	if callCount != expectedCalls {
		t.Fatalf("Expected %d calls, got %d", expectedCalls, callCount)
	}

	expectCounters(t, om, map[string]int{
		daxConnectionsClosedError:                3,
		daxConnectionsCreated:                    3,
		fmt.Sprintf(daxOpNameSuccess, OpGetItem): 0,
		fmt.Sprintf(daxOpNameFailure, OpGetItem): 3,
	})
	expectHistograms(t, om, map[string]int{
		fmt.Sprintf(daxOpNameLatencyUs, OpGetItem): 3,
	})
}

func TestSingleClient_customDialer(t *testing.T) {
	conn := &mockConn{}
	var dialContextFn dialContext = func(ctx context.Context, address string, network string) (net.Conn, error) {
		return conn, nil
	}
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	client, err := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, dialContextFn, nil, om)
	require.NoError(t, err)
	defer client.Close()

	c, _ := client.pool.dialContext(context.TODO(), "address", "network")
	assert.Equal(t, conn, c)
}

type mockConn struct {
	net.Conn
	we, re error
	wd, rd []byte
	cc     map[string]int
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	m.register()
	if m.re != nil {
		return 0, m.re
	}
	if len(m.rd) > 0 {
		l := copy(b, m.rd)
		m.rd = m.rd[l:]
		return l, nil
	}
	return 0, nil
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	m.register()
	if m.we != nil {
		return 0, m.we
	}
	if len(m.wd) > 0 {
		l := copy(m.wd, b)
		m.wd = m.wd[l:]
		return l, nil
	}
	return len(b), nil
}

func (m *mockConn) Close() error {
	m.register()
	return nil
}

func (m *mockConn) SetDeadline(t time.Time) error {
	m.register()
	return nil
}

func (m *mockConn) register() {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	s := strings.Split(fn.Name(), ".")
	n := s[len(s)-1]
	if m.cc == nil {
		m.cc = make(map[string]int)
	}
	m.cc[n]++
}

func (m *mockConn) LocalAddr() net.Addr {
	return nil
}

func (m *mockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

// Test daxRequestFailure where authError() returns true
// Test daxTransactionCanceledFailure where authError() returns true
// Encapsulate both cases when the (mocked) tube session matches and don't match the pool's session,
// as this impacts the tubepool.put method control flow, not recycleTube's
func TestRecycleTube(t *testing.T) {
	mockedConn := mockConn{rd: []byte{cbor.Array}}
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockedConn, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	// Create daxRequestFailure and simulate authError returning true
	drfErr := newDaxRequestFailure([]int{1, 23, 31, 32}, "ec", "msg", "rid", 500, smithy.FaultServer)
	// Create daxTransactionCanceledFailure error type and simulate authError returning true
	dtcfErr := newDaxTransactionCanceledFailure([]int{1, 23, 31, 32}, "ec", "msg", "rid", 500, nil, nil, nil)
	var sessionId int64 = 1
	mockedTube := new(mockTube)
	nextTube, nextTubeErr := newTube(&mockedConn, sessionId)
	if nextTubeErr != nil {
		t.Fatalf("unexpected error %v", nextTubeErr)
	}

	client.pool.session = sessionId
	client.pool.closed = false

	// Identical tubepool session id
	mockedTube.On("Session", mock.Anything).Return(sessionId)
	mockedTube.On("SetNext", mock.Anything).Return()
	mockedTube.On("Next").Return(nextTube)
	mockedTube.On("Close").Return(nil)
	timeNow := time.Now().Unix()
	mockedTube.On("SetAuthExpiryUnix", timeNow).Return()

	client.recycleTube(mockedTube, drfErr)
	mockedTube.AssertCalled(t, "SetAuthExpiryUnix", timeNow)
	client.recycleTube(mockedTube, dtcfErr)
	mockedTube.AssertCalled(t, "SetAuthExpiryUnix", timeNow)

	// Distinct tubepool session id
	mockedTube.ExpectedCalls = nil // Clear previous expectations
	mockedTube.On("Session", mock.Anything).Return(sessionId + 1)
	mockedTube.On("SetNext", mock.Anything).Return()
	mockedTube.On("Next").Return(nextTube)
	mockedTube.On("Close").Return(nil)
	timeNow = time.Now().Unix()
	mockedTube.On("SetAuthExpiryUnix", timeNow).Return()

	client.recycleTube(mockedTube, drfErr)
	mockedTube.AssertCalled(t, "SetAuthExpiryUnix", timeNow)
	client.recycleTube(mockedTube, dtcfErr)
	mockedTube.AssertCalled(t, "SetAuthExpiryUnix", timeNow)
}

// Nil tube and nil error passed as input to recycleTube
func TestRecycleTube_WithNilTube(t *testing.T) {
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	client.recycleTube(nil, nil) // Should do nothing
}

// Non-nil tube, but nil error passed as args to recycleTube
func TestRecycleTube_WithNilError(t *testing.T) {
	mockedConn := mockConn{rd: []byte{cbor.Array}}
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockedConn, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	var sessionId int64 = 1
	client.pool.session = sessionId
	client.pool.closed = false

	mockedTube := new(mockTube)
	nextTube, nextTubeErr := newTube(&mockedConn, sessionId)
	if nextTubeErr != nil {
		t.Fatalf("unexpected error %v", nextTubeErr)
	}

	mockedTube.On("Session", mock.Anything).Return(sessionId)
	mockedTube.On("SetNext", mock.Anything).Return()
	mockedTube.On("Next").Return(nextTube)
	mockedTube.On("Close").Return(nil)

	client.recycleTube(mockedTube, nil)
	mockedTube.AssertNotCalled(t, "SetAuthExpiryUnix")
}

// Test daxRequestFailure where authError() returns false
// Test daxTransactionCanceledFailure where authError() returns false
func TestRecycleTube_NonAuthErrors(t *testing.T) {
	mockedConn := mockConn{rd: []byte{cbor.Array}}
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockedConn, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	// Create daxRequestFailure and simulate authError returning false
	drfErr := newDaxRequestFailure([]int{1, 2, 3, 4}, "ec", "msg", "rid", 500, smithy.FaultServer)
	// Create daxTransactionCanceledFailure error type and simulate authError returning false
	dtcfErr := newDaxTransactionCanceledFailure([]int{1, 2, 3, 4}, "ec", "msg", "rid", 500, nil, nil, nil)
	var sessionId int64 = 1
	mockedTube := new(mockTube)
	nextTube, nextTubeErr := newTube(&mockedConn, sessionId)
	if nextTubeErr != nil {
		t.Fatalf("unexpected error %v", nextTubeErr)
	}

	client.pool.session = sessionId
	client.pool.closed = false

	// Identical tubepool session id
	mockedTube.On("Session", mock.Anything).Return(sessionId)
	mockedTube.On("SetNext", mock.Anything).Return()
	mockedTube.On("Next").Return(nextTube)
	mockedTube.On("Close").Return(nil)

	client.recycleTube(mockedTube, drfErr)
	mockedTube.AssertNotCalled(t, "SetAuthExpiryUnix")
	client.recycleTube(mockedTube, dtcfErr)
	mockedTube.AssertNotCalled(t, "SetAuthExpiryUnix")
}

// Test daxTransactionCanceledFailure with nil embedded daxRequestFailure
func TestRecycleTube_NilDaxRequestFailure(t *testing.T) {
	mockedConn := mockConn{rd: []byte{cbor.Array}}
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockedConn, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	// Create daxTransactionCanceledFailure error type without setting values for any field
	dtcfErr := &daxTransactionCanceledFailure{
		daxRequestFailure:       nil,
		cancellationReasonCodes: nil,
		cancellationReasonMsgs:  nil,
		cancellationReasonItems: nil,
	}
	if dtcfErr.daxRequestFailure != nil {
		t.Fatal("daxRequestFailure is not nil")
	}

	var sessionId int64 = 1
	mockedTube := new(mockTube)
	nextTube, nextTubeErr := newTube(&mockedConn, sessionId)
	if nextTubeErr != nil {
		t.Fatalf("unexpected error %v", nextTubeErr)
	}

	client.pool.session = sessionId
	client.pool.closed = false

	// Identical tubepool session id
	mockedTube.On("Session", mock.Anything).Return(sessionId)
	mockedTube.On("SetNext", mock.Anything).Return()
	mockedTube.On("Next").Return(nextTube)
	mockedTube.On("Close").Return(nil)

	client.recycleTube(mockedTube, dtcfErr)
	mockedTube.AssertNotCalled(t, "SetAuthExpiryUnix")
}

func TestRecycleTube_NonDaxErrors(t *testing.T) {
	// Test with regular errors that don't match the switch cases
	mockedConn := mockConn{rd: []byte{cbor.Array}}
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", &testCredentialProvider{}, 1, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockedConn, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	// Create daxTransactionCanceledFailure error type without setting values for any field
	err := new(error)
	var sessionId int64 = 1
	mockedTube := new(mockTube)
	nextTube, nextTubeErr := newTube(&mockedConn, sessionId)
	if nextTubeErr != nil {
		t.Fatalf("unexpected error %v", nextTubeErr)
	}

	client.pool.session = sessionId
	client.pool.closed = false

	// Identical tubepool session id
	mockedTube.On("Session", mock.Anything).Return(sessionId)
	mockedTube.On("SetNext", mock.Anything).Return()
	mockedTube.On("Next").Return(nextTube)
	mockedTube.On("Close").Return(nil)

	client.recycleTube(mockedTube, *err)
	mockedTube.AssertNotCalled(t, "SetAuthExpiryUnix")
}
