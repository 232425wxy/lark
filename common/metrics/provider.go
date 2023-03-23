package metrics

// Provider 提供了统计功能的接口。
type Provider interface {
	// NewCounter 创建一个计数器。
	NewCounter(CounterOpts) Counter
}

type Counter interface {
	With(labelValues ...string) Counter
	Add(delta float64)
}

type CounterOpts struct {
	Namespace string
	Subsystem string
	Name string

	Help string // 提供关于此metric的帮助信息。

	LabelNames []string
	LabelHelp map[string]string // 提供关于标签的帮助信息

	StatsFormat string
}

type Gauge interface {
	With(labelValues... string) Gauge
	Add(delta float64) // TODO Hyperledger官方考虑移除此功能。
	Set(value float64)
}

type GaugeOpts struct {
	Namespace string
	Subsystem string
	Name string

	Help string // 提供关于此metric的帮助信息。

	LabelNames []string
	LabelHelp map[string]string // 提供关于标签的帮助信息

	StatsFormat string
}

type Histogram interface {
	With(labelValues... string) Histogram
	Observe(value float64)
}

type HistogramOpts struct {
	Namespace string
	Subsystem string
	Name string

	Help string // 提供关于此metric的帮助信息。

	Buckets []float64 // Buckets 可以用来为Prometheus提供桶的边界。当省略时，将使用默认的Prometheus桶值。

	LabelNames []string
	LabelHelp map[string]string // 提供关于标签的帮助信息

	StatsFormat string
}