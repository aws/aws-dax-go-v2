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
	"net"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-dax-go-v2/dax/utils"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type mockDaxAPI struct {
	id int
}

func (m mockDaxAPI) PutItemWithOptions(ctx context.Context, input *dynamodb.PutItemInput, output *dynamodb.PutItemOutput, opt RequestOptions) (*dynamodb.PutItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) DeleteItemWithOptions(ctx context.Context, input *dynamodb.DeleteItemInput, output *dynamodb.DeleteItemOutput, opt RequestOptions) (*dynamodb.DeleteItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) UpdateItemWithOptions(ctx context.Context, input *dynamodb.UpdateItemInput, output *dynamodb.UpdateItemOutput, opt RequestOptions) (*dynamodb.UpdateItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) GetItemWithOptions(ctx context.Context, input *dynamodb.GetItemInput, output *dynamodb.GetItemOutput, opt RequestOptions) (*dynamodb.GetItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) ScanWithOptions(ctx context.Context, input *dynamodb.ScanInput, output *dynamodb.ScanOutput, opt RequestOptions) (*dynamodb.ScanOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) QueryWithOptions(ctx context.Context, input *dynamodb.QueryInput, output *dynamodb.QueryOutput, opt RequestOptions) (*dynamodb.QueryOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) BatchWriteItemWithOptions(ctx context.Context, input *dynamodb.BatchWriteItemInput, output *dynamodb.BatchWriteItemOutput, opt RequestOptions) (*dynamodb.BatchWriteItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) BatchGetItemWithOptions(ctx context.Context, input *dynamodb.BatchGetItemInput, output *dynamodb.BatchGetItemOutput, opt RequestOptions) (*dynamodb.BatchGetItemOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) TransactWriteItemsWithOptions(ctx context.Context, input *dynamodb.TransactWriteItemsInput, output *dynamodb.TransactWriteItemsOutput, opt RequestOptions) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) TransactGetItemsWithOptions(ctx context.Context, input *dynamodb.TransactGetItemsInput, output *dynamodb.TransactGetItemsOutput, opt RequestOptions) (*dynamodb.TransactGetItemsOutput, error) {
	panic("implement me")
}

func (m mockDaxAPI) endpoints(ctx context.Context, opt RequestOptions) ([]serviceEndpoint, error) {
	panic("implement me")
}

func Test_disabledRouteManager(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	rm := newRouteManager(false, time.Second, nil, utils.LogOff, om)
	defer rm.close()
	if rm.isEnabled {
		t.Errorf("Expected route manager to be disabled")
	}

	rm.addRoute("dummy", mockDaxAPI{})
	if len(rm.routes) != 0 {
		t.Errorf("addRoute getting called even with routeManager disabled")
	}

	rm.removeRoute("dummy", mockDaxAPI{}, map[hostPort]clientAndConfig{hostPort{"dummy", 9111}: {client: mockDaxAPI{}}})
	if len(rm.routes) != 0 {
		t.Errorf("addRoute getting called even with routeManager disabled")
	}

	expectCounters(t, om, map[string]int{
		daxRouteManagerRoutesAdded:   0,
		daxRouteManagerRoutesRemoved: 0,
	})
}

func Test_setRoutes(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	rm := newRouteManager(true, time.Second, nil, utils.LogOff, om)
	defer rm.close()
	if len(rm.routes) != 0 {
		t.Errorf("Expected empty routes list, got %v", rm.routes)
	}
	rm.setRoutes(append([]DaxAPI{}, mockDaxAPI{}))
	if len(rm.routes) != 1 {
		t.Errorf("Expected one route but got %v", rm.routes)
	}
}
func Test_getRoute(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	rm := newRouteManager(true, time.Second, nil, utils.LogOff, om)
	defer rm.close()
	if rm.getRoute(nil) != nil {
		t.Errorf("Expected nil route, got other")
	}

	daxAPI1 := mockDaxAPI{1}
	daxAPI2 := mockDaxAPI{2}
	rm.setRoutes(append([]DaxAPI{}, daxAPI1, daxAPI2))

	if rm.getRoute(daxAPI1) != daxAPI2 {
		t.Errorf("Expected route to be daxAPI2, got other")
	}

	if rm.getRoute(daxAPI1) == daxAPI1 {
		t.Errorf("Expected route to be daxAPI2, got daxAPI1")
	}
}

func Test_addRoute(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	rm := newRouteManager(true, time.Second, nil, utils.LogOff, om)
	defer rm.close()
	if len(rm.routes) != 0 {
		t.Errorf("Expected empty routes list, got %v", rm.routes)
	}
	rm.addRoute("dummy", mockDaxAPI{})
	if len(rm.routes) != 1 {
		t.Errorf("Expected one route but got %v", rm.routes)
	}

	expectCounters(t, om, map[string]int{
		daxRouteManagerRoutesAdded:   1,
		daxRouteManagerRoutesRemoved: 0,
	})
}

func Test_removeRoute(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	rm := newRouteManager(true, time.Second, nil, utils.LogOff, om)
	defer rm.close()
	if len(rm.routes) != 0 {
		t.Errorf("Expected empty routes list, got %v", rm.routes)
	}
	daxAPI1 := mockDaxAPI{}
	daxAPI2 := mockDaxAPI{}
	daxAPI3 := mockDaxAPI{}
	dummyHostClientMap := map[hostPort]clientAndConfig{
		hostPort{"dummy.1", 9111}: {client: daxAPI1},
		hostPort{"dummy.2", 9111}: {client: daxAPI2},
		hostPort{"dummy.3", 9111}: {client: daxAPI3},
	}
	rm.setRoutes(append([]DaxAPI{}, daxAPI1, daxAPI2, daxAPI3))
	if len(rm.routes) != 3 {
		t.Errorf("Expected three routes but got %v", rm.routes)
	}

	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	if len(rm.routes) != 2 {
		t.Errorf("Expected two routes but got %v", rm.routes)
	}

	// removing same route again should do nothing
	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	if len(rm.routes) != len(dummyHostClientMap) {
		t.Errorf("Expected two routes but got %v", rm.routes)
	}

	expectCounters(t, om, map[string]int{
		daxRouteManagerRoutesAdded:    0,
		daxRouteManagerRoutesRemoved:  1,
		daxRouteManagerFailOpenEvents: 1,
	})
}

func Test_removeRouteFailOpen(t *testing.T) {
	daxAPI1 := mockDaxAPI{}
	daxAPI2 := mockDaxAPI{}
	daxAPI3 := mockDaxAPI{}
	dummyHostClientMap := map[hostPort]clientAndConfig{
		hostPort{"dummy.1", 9111}: {client: daxAPI1},
		hostPort{"dummy.2", 9111}: {client: daxAPI2},
		hostPort{"dummy.3", 9111}: {client: daxAPI3},
	}
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	rm := newRouteManager(true, time.Second, nil, utils.LogOff, om)
	defer rm.close()
	if len(rm.routes) != 0 {
		t.Errorf("Expected empty routes list, got %v", rm.routes)
	}
	rm.setRoutes(append([]DaxAPI{}, daxAPI1, daxAPI2, daxAPI3))
	if len(rm.routes) != 3 {
		t.Errorf("Expected three routes but got %v", rm.routes)
	}

	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	rm.removeRoute("dummy.2:9111", daxAPI2, dummyHostClientMap)
	if len(rm.routes) != len(dummyHostClientMap) {
		t.Errorf("Fail Open didn't work as expected")
	}

	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	rm.removeRoute("dummy.2:9111", daxAPI2, dummyHostClientMap)
	if len(rm.routes) != len(dummyHostClientMap) {
		t.Errorf("Fail Open didn't work as expected")
	}

	rm.removeRoute("dummy.1:9111", daxAPI1, dummyHostClientMap)
	rm.removeRoute("dummy.2:9111", daxAPI2, dummyHostClientMap)
	if rm.isEnabled {
		t.Errorf("Fail Open didn't work as expected")
	}

	expectCounters(t, om, map[string]int{
		daxRouteManagerRoutesAdded:    0,
		daxRouteManagerRoutesRemoved:  3,
		daxRouteManagerFailOpenEvents: 3,
	})
}

func Test_verifyAndDisable(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	rm := newRouteManager(true, time.Second, nil, utils.LogOff, om)
	defer rm.close()
	rm.disableDuration = 100 * time.Millisecond
	rm.failOpenTimeList = []time.Time{time.Now(), time.Now(), time.Now()}
	rm.verifyAndDisable(time.Now())
	if rm.isEnabled {
		t.Errorf("Expected isRouteManagerEnabled false but got true")
	}

	// this part tests the timer function
	time.Sleep(105 * time.Millisecond)
	if !rm.isEnabled {
		t.Errorf("Fail Open Callback didn't re-open the routeManager")
	}

	rm.failOpenTimeList = []time.Time{time.Now(), time.Now().Add(-5 * time.Second), time.Now().Add(-5 * time.Second)}
	curTime := time.Now()
	rm.verifyAndDisable(curTime)
	if !rm.isEnabled {
		t.Errorf("Fail Open are not continuous so, it shouldn't disable routeManager")
	}
}

func Test_rebuildRoutes(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	rm := newRouteManager(true, time.Second, nil, utils.LogOff, om)
	defer rm.close()
	daxAPI1 := mockDaxAPI{}
	daxAPI2 := mockDaxAPI{}
	daxAPI3 := mockDaxAPI{}
	dummyHostClientMap := map[hostPort]clientAndConfig{
		hostPort{"dummy.1", 9111}: {client: daxAPI1},
		hostPort{"dummy.2", 9111}: {client: daxAPI2},
		hostPort{"dummy.3", 9111}: {client: daxAPI3},
	}
	if len(rm.routes) != 0 {
		t.Errorf("Expected zero routes but got %v", rm.routes)
	}
	rm.rebuildRoutes(dummyHostClientMap)
	if len(rm.routes) != len(dummyHostClientMap) {
		t.Errorf("Expected %v routes but got %v", len(dummyHostClientMap), len(rm.routes))
	}
}

func Test_stopTimer(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	rm := newRouteManager(true, time.Second, nil, utils.LogOff, om)
	defer rm.close()
	timer := time.AfterFunc(rm.disableDuration, func() { rm.isEnabled = true })
	rm.timer = timer
	rm.stopTimer()
	if rm.timer != nil {
		t.Errorf("stopTimer didn't set timer to nil")
	}
	if timer.Stop() {
		t.Errorf("stopTimer didn't stop the timer")
	}
}

func Test_daxConnectionsIdle(t *testing.T) {
	tmp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(tmp)

	tp := newTubePoolWithOptions(
		"test",
		tubePoolOptions{
			maxConcurrentConnAttempts: 20,
			dialContext: func(ctx context.Context, network string, address string) (net.Conn, error) {
				<-time.After(time.Second)
				return nil, os.ErrClosed
			},
		},
		connConfig{},
		om,
	)

	for range 10 {
		go func() {
			tp.get()
		}()
	}

	<-time.After(time.Millisecond * 10)

	expectGauges(t, om, map[string]int{
		daxConcurrentConnectionAttempts: 10,
	})

	for range 10 {
		go func() {
			tp.get()
		}()
	}
	<-time.After(time.Millisecond * 10)

	expectGauges(t, om, map[string]int{
		daxConcurrentConnectionAttempts: 20,
	})

	<-time.After(time.Second)

	expectGauges(t, om, map[string]int{
		daxConcurrentConnectionAttempts: 0,
	})
}
