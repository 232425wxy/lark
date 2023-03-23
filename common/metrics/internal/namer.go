package internal

import (
	"strings"

	"github.com/232425wxy/lark/common/metrics"
)

type Namer struct {
	namespace  string
	subsystem  string
	name       string
	nameFormat string
	labelNames map[string]struct{}
}

func NewCounterNamer(opts metrics.CounterOpts) *Namer {
	return &Namer{
		namespace:  opts.Namespace,
		subsystem:  opts.Subsystem,
		name:       opts.Name,
		nameFormat: opts.StatsFormat,
		labelNames: sliceToSet(opts.LabelNames),
	}
}

func NewGaugeNamer(opts metrics.GaugeOpts) *Namer {
	return &Namer{
		namespace:  opts.Namespace,
		subsystem:  opts.Subsystem,
		name:       opts.Name,
		nameFormat: opts.Help,
		labelNames: sliceToSet(opts.LabelNames),
	}
}

func NewHistogramNamer(opts metrics.HistogramOpts) *Namer {
	return &Namer{
		namespace:  opts.Namespace,
		subsystem:  opts.Subsystem,
		name:       opts.Name,
		nameFormat: opts.StatsFormat,
		labelNames: sliceToSet(opts.LabelNames),
	}
}

func (n *Namer) validateKey(name string) {
	if _, ok := n.labelNames[name]; !ok {
		panic("invalid label name: " + name)
	}
}

func (n *Namer) labelsToMap(labelValues []string) map[string]string {
	labels := map[string]string{}
	for i := 0; i < len(labelValues); i += 2 {
		key := labelValues[i]
		n.validateKey(key)
		if i == len(labelValues)-1 {
			labels[key] = "unknown"
		} else {
			labels[key] = labelValues[i+1]
		}
	}
	return labels
}

func (n *Namer) FullyQualifiedName() string {
	switch {
	case n.namespace != "" && n.subsystem != "":
		return strings.Join([]string{n.namespace, n.subsystem, n.name}, ".")
	case n.namespace != "":
		return strings.Join([]string{n.namespace, n.name}, ".")
	case n.subsystem != "":
		return strings.Join([]string{n.subsystem, n.name}, ".")
	default:
		return n.name
	}
}

func sliceToSet(slice []string) map[string]struct{} {
	labelsSet := map[string]struct{}{}
	for _, s := range slice {
		labelsSet[s] = struct{}{}
	}
	return labelsSet
}
