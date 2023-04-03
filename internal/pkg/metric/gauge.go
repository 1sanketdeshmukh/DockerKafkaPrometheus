package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

func messageGuage() prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "client",
		Name:      "message_gauge",
		Help:      "current message",
	})
}
