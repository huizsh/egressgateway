// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spidernet-io/egressgateway/pkg/iptables"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

func RegisterMetricCollectors() {
	var metricCollectors []prometheus.Collector
	metricCollectors = append(metricCollectors, iptables.MetricCollectors()...)
	for _, collector := range metricCollectors {
		metrics.Registry.MustRegister(collector)
	}
}
