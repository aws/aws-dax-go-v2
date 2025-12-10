/*
  Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCountMetricInt64(t *testing.T) {
	mp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(mp)

	countMetricInt64(context.TODO(), om, fmt.Sprintf(daxOpNameFailure, OpGetItem), 1)

	assert.Len(t, mp.meters, 1)
	s, ok := mp.meters[daxMeterScope]
	assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, daxMeterScope))
	assert.NotNil(t, s)
	if !ok || s == nil {
		return
	}

	tm, ok := s.(*testMeter)
	assert.True(t, ok)
	if !ok {
		return
	}

	keys := map[string]int{
		fmt.Sprintf(daxOpNameFailure, OpGetItem): 1,
	}

	for k, v := range keys {
		i, ok := tm.i64s[k]
		assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, k))
		assert.Equal(t, v, len(i.data), k)
	}
}

func TestGaugeInt64(t *testing.T) {
	mp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(mp)

	gaugeInt64(context.TODO(), om, daxConnectionsIdle, 1)

	assert.Len(t, mp.meters, 1)
	s, ok := mp.meters[daxMeterScope]
	assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, daxMeterScope))
	assert.NotNil(t, s)
	if !ok || s == nil {
		return
	}

	tm, ok := s.(*testMeter)
	assert.True(t, ok)
	if !ok {
		return
	}

	keys := map[string]int{
		daxConnectionsIdle: 1,
	}

	for k, v := range keys {
		i, ok := tm.i64s[k]
		assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, k))
		assert.Equal(t, v, len(i.data), k)
	}
}

func TestHistogramMicrosecondsInt64(t *testing.T) {
	mp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(mp)

	startTime := time.Now()
	histogramMicrosecondsInt64(context.TODO(), om, fmt.Sprintf(daxOpNameLatencyUs, OpGetItem), startTime)

	assert.Len(t, mp.meters, 1)
	s, ok := mp.meters[daxMeterScope]
	assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, daxMeterScope))
	assert.NotNil(t, s)
	if !ok || s == nil {
		return
	}

	tm, ok := s.(*testMeter)
	assert.True(t, ok)
	if !ok {
		return
	}

	keys := map[string]int{
		fmt.Sprintf(daxOpNameLatencyUs, OpGetItem): 1,
	}

	for k, v := range keys {
		i, ok := tm.i64s[k]
		assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, k))
		assert.Equal(t, v, len(i.data), k)
	}

}

func TestWithMicrosecondHistogramInt64(t *testing.T) {
	mp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(mp)

	_, _ = withMicrosecondHistogramInt64(
		context.TODO(),
		om,
		fmt.Sprintf(daxOpNameLatencyUs, OpGetItem),
		func() (any, error) {
			return nil, nil
		},
	)

	assert.Len(t, mp.meters, 1)
	s, ok := mp.meters[daxMeterScope]
	assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, daxMeterScope))
	assert.NotNil(t, s)
	if !ok || s == nil {
		return
	}

	tm, ok := s.(*testMeter)
	assert.True(t, ok)
	if !ok {
		return
	}

	keys := map[string]int{
		fmt.Sprintf(daxOpNameLatencyUs, OpGetItem): 1,
	}

	for k, v := range keys {
		i, ok := tm.i64s[k]
		assert.True(t, ok, fmt.Sprintf(`expected key "%s" to exist in meters map`, k))
		assert.Equal(t, v, len(i.data), k)
	}
}

func BenchmarkCountMetricInt64Nop(b *testing.B) {
	mp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(mp)
	ctx := context.TODO()

	for range b.N {
		countMetricInt64(ctx, om, daxConnectionsCreated, 1)
	}
}

func BenchmarkCountMetricInt64Test(b *testing.B) {
	mp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(mp)
	ctx := context.TODO()

	for range b.N {
		countMetricInt64(ctx, om, daxConnectionsCreated, 1)
	}
}

func BenchmarkGaugeInt64Nop(b *testing.B) {
	mp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(mp)
	ctx := context.TODO()

	for range b.N {
		gaugeInt64(ctx, om, daxConnectionsIdle, 1)
	}
}

func BenchmarkGaugeInt64Test(b *testing.B) {
	mp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(mp)
	ctx := context.TODO()

	for range b.N {
		gaugeInt64(ctx, om, daxConnectionsIdle, 1)
	}
}

func BenchmarkHistogramMicrosecondsInt64Nop(b *testing.B) {
	mp := &testMeterProvider{}
	om, err := buildDaxSdkMetrics(mp)
	if err != nil {
		b.Fatal(err)
	}
	ctx := context.TODO()
	startTime := time.Now()
	name := fmt.Sprintf(daxOpNameLatencyUs, OpGetItem)

	for range b.N {
		histogramMicrosecondsInt64(ctx, om, name, startTime)
	}
}

func BenchmarkHistogramMicrosecondsInt64Test(b *testing.B) {
	mp := &testMeterProvider{}
	om, _ := buildDaxSdkMetrics(mp)
	ctx := context.TODO()
	startTime := time.Now()
	name := fmt.Sprintf(daxOpNameLatencyUs, OpGetItem)

	for range b.N {
		histogramMicrosecondsInt64(ctx, om, name, startTime)
	}
}
