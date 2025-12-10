# AWS DAX SDK for Go v2

aws-dax-go-v2 is the official AWS DAX SDK for the Go programming language. https://aws.amazon.com/dynamodb/dax

Checkout our [release notes](https://github.com/aws/aws-dax-go-v2/releases) for
information about the latest bug fixes, updates, and features added to the SDK.

## Getting started

The best way to get started working with the SDK is to use go get to add the SDK
to your Go Workspace manually.

    go get github.com/aws/aws-dax-go-v2

## Making API requests

This example shows how you can use the AWS DAX SDK to make an API request.

```go
package main

import (
	"context"
	"fmt"
	"net"

	"github.com/aws/aws-dax-go-v2/dax"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	ctx := context.Background()

	cfg := dax.DefaultConfig()
	cfg.HostPorts = []string{"dax://mycluster.frfx8h.clustercfg.dax.usw2.amazonaws.com:8111"}
	cfg.Region = "us-west-2"
	client, err := dax.New(cfg)

	if err != nil {
		panic(fmt.Errorf("unable to initialize client %v", err))
	}

	//Connecion to a secure cluster
	secureEndpoint := "daxs://mycluster.frfx8h.clustercfg.dax.usw2.amazonaws.com"
	secureCfg := dax.DefaultConfig()
	secureCfg.HostPorts = []string{secureEndpoint}
	secureCfg.Region = "us-west-2"

	//WARN: Skip hostname verification of TLS connections. 
	//The default is to perform hostname verification, setting this to True will skip verification. 
	//Be sure you understand the implication of doing so, which is the inability to authenticate
	//the cluster that you are connecting to.
	secureCfg.SkipHostnameVerification = false

	// DialContext is an optional field in Config.
	// If DialContext is being set in Config for a secure/ encrypted cluster, then use dax.SecureDialContext to 
	// return DialContext. An example of how DailContext can be set using dax.SecureDialContext is shown below.
	secureCfg.DialContext = func(ctx context.Context, network string, address string) (net.Conn, error) {
		// fmt.Println("Write your custom logic here")
		dialCon, err := dax.SecureDialContext(secureEndpoint, secureCfg.SkipHostnameVerification)
		if err != nil {
			panic(fmt.Errorf("secure dialcontext creation failed %v", err))
		}
		return dialCon(ctx, network, address)
	}
	secureClient, err := dax.New(secureCfg)
	if err != nil {
		panic(fmt.Errorf("unable to initialize secure client %v", err))
	}
	fmt.Println("secure client created", secureClient)

	// Marshal the values
	pkVal, err := attributevalue.Marshal("mykey")
	if err != nil {
		return fmt.Errorf("error marshaling pk value: %v", err)
	}

	skVal, err := attributevalue.Marshal("0")
	if err != nil {
		return fmt.Errorf("error marshaling sk value: %v", err)
	}

	valueVal, err := attributevalue.Marshal("myvalue")
	if err != nil {
		return fmt.Errorf("error marshaling value: %v", err)
	}

	// PutItem operation
	input := &dynamodb.PutItemInput{
		TableName: aws.String("TryDaxGoTable"),
		Item: map[string]types.AttributeValue{
			"pk":    pkVal,
			"sk":    skVal,
			"value": valueVal,
		},
	}

	output, err := client.PutItem(ctx, input)
	if err != nil {
		panic(fmt.Errorf("unable to make request %v", err))
	}

	fmt.Println("Output: ", output)

	// Scan with pagination example using DAX paginator
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String("TryDaxGoTable"),
		Limit:     aws.Int32(10), // Items per page
	}

	// Create paginator
	paginator := dax.NewScanPaginator(client, scanInput)

	// Iterate through pages
	pageNum := 0
	for paginator.HasMorePages() {
		pageNum++
		fmt.Printf("Processing page %d\n", pageNum)

		// Get next page
		page, err := paginator.NextPage(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to get page %d: %v", pageNum, err))
		}

		// Process items in the page
		for _, item := range page.Items {
			fmt.Printf("Item: %v\n", item)
		}
	}
}
```

## Metrics

The Dax SDK produces a number of metrics which can be sent to CloudWatch or any other logging platform.
You simply have to provide a meter provider which satisfies
the [MeterProvider](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#MeterProvider) interface
from [smithy-go](https://github.com/aws/smithy-go).

See the [example](#example-with-meter-provider) below.

### List of metrics

| Type                  | Metric Name                            | Metric Type                                                                                  | Description                                                         |
|-----------------------|----------------------------------------|----------------------------------------------------------------------------------------------|---------------------------------------------------------------------|
| Operation Metrics     | `dax.op.API_OPERATION_NAME.success`    | [Int64Counter](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Counter)     | The number of successful calls for each operation                   |
| Operation Metrics     | `dax.op.API_OPERATION_NAME.failure`    | [Int64Counter](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Counter)     | The number of failed calls for each operation                       |
| Operation Metrics     | `dax.op.API_OPERATION_NAME.latency_us` | [Int64Histogram](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Histogram) | The latency in microseconds for each operation                      |
| Connection Metrics    | `dax.connections.created`              | [Int64Counter](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Counter)     | Total amount of created connections                                 |
| Connection Metrics    | `dax.connections.closed.error`         | [Int64Counter](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Counter)     | Number of closed connections due to errors                          |
| Connection Metrics    | `dax.connections.closed.idle`          | [Int64Counter](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Counter)     | Number of closed connections due to inactivity                      |
| Connection Metrics    | `dax.connections.closed.session`       | [Int64Counter](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Counter)     | Number of closed connections due to poll session change             |
| Connection Metrics    | `dax.connections.attempts`             | [Int64Gauge](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Gauge)         | Current number of concurrent connection attempts                    |
| Connection Metrics    | `dax.connections.idle`                 | [Int64Gauge](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Gauge)         | Current number of inactive connections in the pool                  |
| Route Manager Metrics | `dax.route_manager.routes.added`       | [Int64Counter](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Counter)     | The number of routes added back to the active pool.                 |              
| Route Manager Metrics | `dax.route_manager.routes.removed`     | [Int64Counter](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Counter)     | The number of routes removed from the active pool due to problems.  |  
| Route Manager Metrics | `dax.route_manager.fail_open.events`   | [Int64Counter](https://pkg.go.dev/github.com/aws/smithy-go@v1.22.3/metrics#Int64Counter)     | The number of events when the manager enters the "fail-open" state. |

| `API_OPERATION_NAME` |
|----------------------|
| `BatchGetItem`       |
| `BatchWriteItem`     |
| `DeleteItem`         |
| `GetItem`            |
| `PutItem`            |
| `Query`              |
| `Scan`               |
| `TransactGetItems`   |
| `TransactWriteItems` |
| `UpdateItem`         |

### Example with Meter Provider:

```go
package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-dax-go-v2/dax"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/metrics"
)

func main() {
	region := "eu-west-1"
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
	)
	if err != nil {
		panic(err)
	}

	cw := cloudwatch.NewFromConfig(cfg)
	if cw == nil {
		panic("unable to create cloudwatch client")
	}

	daxCfg := dax.NewConfig(cfg, "dax://mycluster.frfx8h.clustercfg.dax.usw2.amazonaws.com:8111")
	// enable route manager health check (will generate metrics if enabled)
	daxCfg.RouteManagerEnabled = true
	daxCfg.MeterProvider = &MyMeterProvider{cw: *cw}

	dax, err := dax.New(daxCfg)
	if err != nil {
		panic(err)
	}

	_, _ = dax.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String("mytable"),
		Key: map[string]ddbtypes.AttributeValue{
			"pk": &ddbtypes.AttributeValueMemberN{Value: "251"},
			"sk": &ddbtypes.AttributeValueMemberN{Value: "1"},
		},
	})
}

type MyMeterProvider struct {
	cw cloudwatch.Client

	meters map[string]metrics.Meter
}

func (m *MyMeterProvider) Meter(scope string, _ ...metrics.MeterOption) metrics.Meter {
	if m.meters == nil {
		m.meters = make(map[string]metrics.Meter)
	}

	mtr := m.meters[scope]
	if mtr != nil {
		return mtr
	}

	mtr = &MyMeter{scope: scope, cw: m.cw}
	m.meters[scope] = mtr

	return mtr
}

type MyMeter struct {
	scope string
	cw    cloudwatch.Client

	counters   map[string]metrics.Int64Counter
	histograms map[string]metrics.Int64Histogram
	gauges     map[string]metrics.Int64Gauge
}

func (m *MyMeter) Int64Counter(name string, _ ...metrics.InstrumentOption) (metrics.Int64Counter, error) {
	if m.counters == nil {
		m.counters = make(map[string]metrics.Int64Counter)
	}

	c := m.counters[name]
	if c != nil {
		return c, nil
	}

	c = &MyInstrument{scope: m.scope, name: name, cw: m.cw}
	m.counters[name] = c

	return c, nil
}

func (m *MyMeter) Int64UpDownCounter(_ string, _ ...metrics.InstrumentOption) (metrics.Int64UpDownCounter, error) {
	panic("not used")
}

func (m *MyMeter) Int64Gauge(name string, _ ...metrics.InstrumentOption) (metrics.Int64Gauge, error) {
	if m.gauges == nil {
		m.gauges = make(map[string]metrics.Int64Gauge)
	}

	g := m.gauges[name]
	if g != nil {
		return g, nil
	}

	g = &MyInstrument{scope: m.scope, name: name, cw: m.cw}
	m.gauges[name] = g

	return g, nil
}

func (m *MyMeter) Int64Histogram(name string, _ ...metrics.InstrumentOption) (metrics.Int64Histogram, error) {
	if m.histograms == nil {
		m.histograms = make(map[string]metrics.Int64Histogram)
	}

	h := m.histograms[name]
	if h != nil {
		return h, nil
	}

	h = &MyInstrument{scope: m.scope, name: name, cw: m.cw}
	m.histograms[name] = h

	return h, nil
}

func (m *MyMeter) Int64AsyncCounter(_ string, _ metrics.Int64Callback, _ ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	panic("not used")
}

func (m *MyMeter) Int64AsyncUpDownCounter(_ string, _ metrics.Int64Callback, _ ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	panic("not used")
}

func (m *MyMeter) Int64AsyncGauge(_ string, _ metrics.Int64Callback, _ ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	panic("not used")
}

func (m *MyMeter) Float64Counter(_ string, _ ...metrics.InstrumentOption) (metrics.Float64Counter, error) {
	panic("not used")
}

func (m *MyMeter) Float64UpDownCounter(_ string, _ ...metrics.InstrumentOption) (metrics.Float64UpDownCounter, error) {
	panic("not used")
}

func (m *MyMeter) Float64Gauge(_ string, _ ...metrics.InstrumentOption) (metrics.Float64Gauge, error) {
	panic("not used")
}

func (m *MyMeter) Float64Histogram(_ string, _ ...metrics.InstrumentOption) (metrics.Float64Histogram, error) {
	panic("not used")
}

func (m *MyMeter) Float64AsyncCounter(name string, callback metrics.Float64Callback, opts ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	panic("not used")
}

func (m *MyMeter) Float64AsyncUpDownCounter(name string, callback metrics.Float64Callback, opts ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	panic("not used")
}

func (m *MyMeter) Float64AsyncGauge(name string, callback metrics.Float64Callback, opts ...metrics.InstrumentOption) (metrics.AsyncInstrument, error) {
	panic("not used")
}

type MyInstrument struct {
	scope string
	name  string
	cw    cloudwatch.Client
}

// Sample - gauge
func (m *MyInstrument) Sample(ctx context.Context, i int64, _ ...metrics.RecordMetricOption) {
	_, _ = m.cw.PutMetricData(ctx, &cloudwatch.PutMetricDataInput{
		Namespace: aws.String(m.scope),
		MetricData: []types.MetricDatum{
			{
				MetricName: aws.String(m.name),
				Values:     []float64{float64(i)},
			},
		},
	})
}

// Record - histogram
func (m *MyInstrument) Record(ctx context.Context, i int64, option ...metrics.RecordMetricOption) {
	_, _ = m.cw.PutMetricData(ctx, &cloudwatch.PutMetricDataInput{
		Namespace: aws.String(m.scope),
		MetricData: []types.MetricDatum{
			{
				MetricName: aws.String(m.name),
				Values:     []float64{float64(i)},
			},
		},
	})
}

// Add - Counter
func (m *MyInstrument) Add(ctx context.Context, i int64, option ...metrics.RecordMetricOption) {
	_, _ = m.cw.PutMetricData(ctx, &cloudwatch.PutMetricDataInput{
		Namespace: aws.String(m.scope),
		MetricData: []types.MetricDatum{
			{
				MetricName: aws.String(m.name),
				Values:     []float64{float64(i)},
			},
		},
	})
}
```

## Feedback and contributing

**GitHub issues:** To provide feedback or report bugs, file GitHub
[Issues](https://github.com/aws/aws-dax-go-v2/issues) on the SDK.
This is the preferred mechanism to give feedback so that other users can engage in
the conversation, +1 issues, etc. Issues you open will be evaluated, and included
in our roadmap.

**Contributing:** You can open pull requests for fixes or additions to the
AWS DAX SDK for Go v2. All pull requests must be submitted under the Apache 2.0
license and will be reviewed by an SDK team member before being merged in.
Accompanying unit tests, where possible, are appreciated.

## License

This library is licensed under the Apache 2.0 License. 
