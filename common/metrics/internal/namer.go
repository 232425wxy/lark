package internal

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/232425wxy/lark/common/metrics"
)

var (
	formatRegexp = regexp.MustCompile(`%{([#?[:alnum:]_]+)}`)
	invalidLabelValueRegexp = regexp.MustCompile(`[.|:\s]`)
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

// 将{k1,v1,k2,v2,k3,v3...}转换为{k1:v1,k2:v2,k3:v3...}
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

func (n *Namer) Format(labelValues ...string) string {
	labels := n.labelsToMap(labelValues)
	cursor := 0
	var segments []string
	matches := formatRegexp.FindAllStringSubmatchIndex(n.nameFormat, -1)
	for _, m := range matches {
		start, end := m[0], m[1]
		labelStart, labelEnd := m[2], m[3]
		if start > cursor {
			segments = append(segments, n.nameFormat[cursor:start])
		}
		key := n.nameFormat[labelStart:labelEnd]
		var value string
		switch key {
		case "#namespace":
			value = n.namespace
		case "#subsystem":
			value = n.subsystem
		case "#name":
			value = n.name
		case "#fqname":
			value = n.FullyQualifiedName()
		default:
			var ok bool
			value, ok = labels[key]
			if !ok {
				panic(fmt.Sprintf("invalid label in name format: %s", key))
			}
			value = invalidLabelValueRegexp.ReplaceAllString(value, "_")
		}
		segments = append(segments, value)
		cursor = end
	}

	if cursor != len(n.nameFormat) {
		segments = append(segments, n.nameFormat[cursor:])
	}

	return strings.Join(segments, "")
}

func sliceToSet(slice []string) map[string]struct{} {
	labelsSet := map[string]struct{}{}
	for _, s := range slice {
		labelsSet[s] = struct{}{}
	}
	return labelsSet
}

