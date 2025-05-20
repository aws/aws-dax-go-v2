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
	"strings"
	"time"

	"github.com/aws/smithy-go/metrics"
)

const (
	daxMeterScope = "github.com/aws/aws-dax-go-v2"

	daxOpNameSuccess                = "dax.op.%s.success"
	daxOpNameFailure                = "dax.op.%s.failure"
	daxOpNameLatencyUs              = "dax.op.%s.latency_us"     // histogram
	daxConnectionsIdle              = "dax.connections.idle"     // gauge
	daxConcurrentConnectionAttempts = "dax.connections.attempts" // gauge
	daxConnectionsCreated           = "dax.connections.created"
	daxConnectionsClosedError       = "dax.connections.closed.error"
	daxConnectionsClosedIdle        = "dax.connections.closed.idle"
	daxConnectionsClosedSession     = "dax.connections.closed.session"
	daxRouteManagerRoutesAdded      = "dax.route_manager.routes.added"
	daxRouteManagerRoutesRemoved    = "dax.route_manager.routes.removed"
	daxRouteManagerFailOpenEvents   = "dax.route_manager.fail_open.events"
)

type daxSdkMetrics struct {
	counters   map[string]metrics.Int64Counter
	histograms map[string]metrics.Int64Histogram
	gauges     map[string]metrics.Int64Gauge
}

func (m *daxSdkMetrics) counterFor(name string) metrics.Int64Counter {
	return m.counters[name]
}

func (m *daxSdkMetrics) histogramFor(name string) metrics.Int64Histogram {
	return m.histograms[name]
}

func (m *daxSdkMetrics) gaugeFor(name string) metrics.Int64Gauge {
	return m.gauges[name]
}

func buildCounters(meter metrics.Meter, om *daxSdkMetrics, ops []string) (err error) {
	counters := map[string]string{
		daxOpNameSuccess:              "Operations %s success",
		daxOpNameFailure:              "Operations %s failure",
		daxConnectionsCreated:         "Total amount of created connections",
		daxConnectionsClosedError:     "Number of closed connections due to errors",
		daxConnectionsClosedIdle:      "Number of closed connections due to inactivity",
		daxConnectionsClosedSession:   "Number of closed connections due to poll session change",
		daxRouteManagerRoutesAdded:    "The number of routes added back to the active pool.",
		daxRouteManagerRoutesRemoved:  "The number of routes removed from the active pool due to problems.",
		daxRouteManagerFailOpenEvents: `The number of events when the manager enters the "fail-open" state.`,
	}

	for name, description := range counters {
		if strings.Contains(name, "%s") {
			for _, op := range ops {
				metricName := fmt.Sprintf(name, op)
				metricDescription := fmt.Sprintf(description, op)

				om.counters[metricName], err = operationCounter(meter, metricName, metricDescription)

				if err != nil {
					return
				}
			}

			continue
		}

		om.counters[name], err = operationCounter(meter, name, description)

		if err != nil {
			return
		}
	}

	return
}

func buildHistograms(meter metrics.Meter, om *daxSdkMetrics, ops []string) (err error) {
	histograms := map[string]string{
		daxOpNameLatencyUs: "Operations %s latency in microseconds",
	}

	// build histograms
	for name, description := range histograms {
		if strings.Contains(name, "%s") {
			for _, op := range ops {
				metricName := fmt.Sprintf(name, op)
				metricDescription := fmt.Sprintf(description, op)

				om.histograms[metricName], err = operationHistogram(meter, metricName, metricDescription)

				if err != nil {
					return
				}
			}

			continue
		}

		om.histograms[name], err = operationHistogram(meter, name, description)
		if err != nil {
			return
		}
	}

	return
}

func buildGauges(meter metrics.Meter, om *daxSdkMetrics, ops []string) (err error) {
	gauges := map[string]string{
		daxConnectionsIdle:              "Current number of inactive connections in the pool",
		daxConcurrentConnectionAttempts: "Current number of concurrent connection attempts",
	}

	// build gauges
	for name, description := range gauges {
		om.gauges[name], err = operationGauge(meter, name, description)

		if err != nil {
			return
		}
	}

	return
}

func buildDaxSdkMetrics(mp metrics.MeterProvider) (*daxSdkMetrics, error) {
	meter := mp.Meter(daxMeterScope)

	sdkMetrics := &daxSdkMetrics{
		counters:   make(map[string]metrics.Int64Counter),
		histograms: make(map[string]metrics.Int64Histogram),
		gauges:     make(map[string]metrics.Int64Gauge),
	}

	ops := []string{
		OpGetItem,
		OpPutItem,
		OpUpdateItem,
		OpDeleteItem,
		OpBatchGetItem,
		OpBatchWriteItem,
		OpTransactGetItems,
		OpTransactWriteItems,
		OpQuery,
		OpScan,
	}

	if err := buildCounters(meter, sdkMetrics, ops); err != nil {
		return nil, err
	}

	if err := buildHistograms(meter, sdkMetrics, ops); err != nil {
		return nil, err
	}

	if err := buildGauges(meter, sdkMetrics, ops); err != nil {
		return nil, err
	}

	return sdkMetrics, nil
}

func operationCounter(m metrics.Meter, name string, description string) (metrics.Int64Counter, error) {
	return m.Int64Counter(name, func(o *metrics.InstrumentOptions) {
		o.Description = description
	})
}

func operationHistogram(m metrics.Meter, name string, description string) (metrics.Int64Histogram, error) {
	return m.Int64Histogram(name, func(o *metrics.InstrumentOptions) {
		o.UnitLabel = "Microseconds"
		o.Description = description
	})
}

func operationGauge(m metrics.Meter, name string, description string) (metrics.Int64Gauge, error) {
	return m.Int64Gauge(name, func(o *metrics.InstrumentOptions) {
		o.Description = description
	})
}

type metricFunction[T any] func() (T, error)

func countMetricInt64(ctx context.Context, om *daxSdkMetrics, name string, v int64) {
	c := om.counterFor(name)

	if c == nil {
		return
	}

	c.Add(ctx, v)
}

func gaugeInt64(ctx context.Context, om *daxSdkMetrics, name string, v int64) {
	g := om.gaugeFor(name)

	if g == nil {
		return
	}

	g.Sample(ctx, v)
}

func histogramMicrosecondsInt64(ctx context.Context, om *daxSdkMetrics, name string, t time.Time) {
	h := om.histogramFor(name)

	if h == nil {
		return
	}

	h.Record(ctx, time.Since(t).Microseconds())
}

func withMicrosecondHistogramInt64[T any](ctx context.Context, om *daxSdkMetrics, name string, fn metricFunction[T]) (T, error) {
	startTime := time.Now()

	out, err := fn()

	histogramMicrosecondsInt64(ctx, om, name, startTime)

	return out, err
}
