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
	"testing"

	"github.com/aws/smithy-go/metrics"
	"github.com/stretchr/testify/assert"
)

// testMeterProvider is naive metrics implementation for testing purposes
type testMeterProvider struct {
	meters map[string]metrics.Meter
}

var _ metrics.MeterProvider = (*testMeterProvider)(nil)

// Meter returns a meter which creates no-op instruments.
func (t *testMeterProvider) Meter(name string, _ ...metrics.MeterOption) metrics.Meter {
	if t.meters == nil {
		t.meters = map[string]metrics.Meter{}
	}

	if out, ok := t.meters[name]; ok {
		return out
	}

	out := &testMeter{}

	t.meters[name] = out

	return out
}

type testMeter struct {
	i64s map[string]*testInstrument[int64]
	f64s map[string]*testInstrument[float64]
}

var _ metrics.Meter = (*testMeter)(nil)

func (t *testMeter) instrumentInt64(name string) *testInstrument[int64] {
	if t.i64s == nil {
		t.i64s = map[string]*testInstrument[int64]{}
	}

	if out, ok := t.i64s[name]; ok {
		return out
	}

	out := &testInstrument[int64]{}

	t.i64s[name] = out

	return out
}

func (t *testMeter) instrumentFloat64(name string) *testInstrument[float64] {
	if t.f64s == nil {
		t.f64s = map[string]*testInstrument[float64]{}
	}

	if out, ok := t.f64s[name]; ok {
		return out
	}

	out := &testInstrument[float64]{}

	t.f64s[name] = out

	return out
}

func (t *testMeter) Int64Counter(name string, _ ...metrics.InstrumentOption) (metrics.Int64Counter, error) {
	return t.instrumentInt64(name), nil
}
func (t *testMeter) Int64UpDownCounter(name string, _ ...metrics.InstrumentOption) (metrics.Int64UpDownCounter, error) {
	return t.instrumentInt64(name), nil
}
func (t *testMeter) Int64Gauge(name string, _ ...metrics.InstrumentOption) (metrics.Int64Gauge, error) {
	return t.instrumentInt64(name), nil
}
func (t *testMeter) Int64Histogram(name string, _ ...metrics.InstrumentOption) (metrics.Int64Histogram, error) {
	return t.instrumentInt64(name), nil
}
func (t *testMeter) Int64AsyncCounter(name string, _ metrics.Int64Callback, _ ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	return t.instrumentInt64(name), nil
}
func (t *testMeter) Int64AsyncUpDownCounter(name string, _ metrics.Int64Callback, _ ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	return t.instrumentInt64(name), nil
}
func (t *testMeter) Int64AsyncGauge(name string, _ metrics.Int64Callback, _ ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	return t.instrumentInt64(name), nil
}
func (t *testMeter) Float64Counter(name string, _ ...metrics.InstrumentOption) (metrics.Float64Counter, error) {
	return t.instrumentFloat64(name), nil
}
func (t *testMeter) Float64UpDownCounter(name string, _ ...metrics.InstrumentOption) (metrics.Float64UpDownCounter, error) {
	return t.instrumentFloat64(name), nil
}
func (t *testMeter) Float64Gauge(name string, _ ...metrics.InstrumentOption) (metrics.Float64Gauge, error) {
	return t.instrumentFloat64(name), nil
}
func (t *testMeter) Float64Histogram(name string, _ ...metrics.InstrumentOption) (metrics.Float64Histogram, error) {
	return t.instrumentFloat64(name), nil
}
func (t *testMeter) Float64AsyncCounter(name string, _ metrics.Float64Callback, _ ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	return t.instrumentFloat64(name), nil
}
func (t *testMeter) Float64AsyncUpDownCounter(name string, _ metrics.Float64Callback, _ ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	return t.instrumentFloat64(name), nil
}
func (t *testMeter) Float64AsyncGauge(name string, _ metrics.Float64Callback, _ ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	return t.instrumentFloat64(name), nil
}

type testInstrument[N int64 | float64] struct {
	data      []N
	callbacks []any
	stopCh    chan bool
}

func (t *testInstrument[N]) Add(_ context.Context, n N, _ ...metrics.RecordMetricOption) {
	if len(t.data) == 0 {
		t.data = append(t.data, n)
	} else {
		t.data[0] += n
	}
}

func (t *testInstrument[N]) Sample(_ context.Context, n N, _ ...metrics.RecordMetricOption) {
	t.data = []N{n}
}

func (t *testInstrument[N]) Record(_ context.Context, n N, _ ...metrics.RecordMetricOption) {
	t.data = append(t.data, n)
}

func (testInstrument[_]) Stop() {}

func counter(om *daxSdkMetrics, name string) (metrics.Int64Counter, bool, int) {
	c, ok := om.counters[name]
	val := 0
	if c != nil {
		if i, iOk := c.(*testInstrument[int64]); iOk && len(i.data) > 0 {
			val = int(i.data[len(i.data)-1])
		}
	}
	return c, ok, val
}

func gauge(om *daxSdkMetrics, name string) (metrics.Int64Gauge, bool, int) {
	g, ok := om.gauges[name]
	val := 0
	if g != nil {
		if i, iOk := g.(*testInstrument[int64]); iOk && len(i.data) > 0 {
			val = int(i.data[0])
		}
	}
	return g, ok, val
}

func histogram(om *daxSdkMetrics, name string) (metrics.Int64Histogram, bool, int) {
	h, ok := om.histograms[name]
	val := 0
	if h != nil {
		if i, iOk := h.(*testInstrument[int64]); iOk && len(i.data) > 0 {
			val = len(i.data)
		}
	}
	return h, ok, val
}

func expectCounters(t *testing.T, om *daxSdkMetrics, counters map[string]int) {
	for k, v := range counters {
		c, exists, cnt := counter(om, k)
		assert.NotNil(t, c, k)
		assert.True(t, exists)
		if v != -1 {
			assert.Equal(t, v, cnt, k)
		} else {
			assert.Greater(t, cnt, 0, k)
		}
	}
}

func expectGauges(t *testing.T, om *daxSdkMetrics, gauges map[string]int) {
	for k, v := range gauges {
		c, exists, cnt := gauge(om, k)
		assert.NotNil(t, c, k)
		assert.True(t, exists)
		if v != -1 {
			assert.Equal(t, v, cnt, k)
		} else {
			assert.Greater(t, cnt, 0, k)
		}
	}
}

func expectHistograms(t *testing.T, om *daxSdkMetrics, histograms map[string]int) {
	for k, v := range histograms {
		c, exists, cnt := histogram(om, k)
		assert.NotNil(t, c, k)
		assert.True(t, exists)
		if v != -1 {
			assert.Equal(t, v, cnt, k)
		} else {
			assert.Greater(t, cnt, 0, k)
		}
	}
}

func expectMetrics(
	t *testing.T,
	om *daxSdkMetrics,
	counters map[string]int,
	gauges map[string]int,
	histograms map[string]int,
) {
}
