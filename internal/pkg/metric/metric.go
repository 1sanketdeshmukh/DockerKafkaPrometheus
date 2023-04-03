package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	KafkaMessageCounter     *prometheus.CounterVec
	KafkaDropMessageCounter *prometheus.CounterVec
	MessageGuage            prometheus.Gauge
	MessageHistogram        *prometheus.HistogramVec
}

func New(registry *prometheus.Registry) Metric {
	m := &Metric{}

	m.KafkaMessageCounter = kafkaMessageCounter()
	registry.MustRegister(m.KafkaMessageCounter)

	m.KafkaDropMessageCounter = kafkaDropMessageCounter()
	registry.MustRegister(m.KafkaDropMessageCounter)

	m.MessageGuage = messageGuage()
	registry.MustRegister(m.MessageGuage)

	m.MessageHistogram = messageHistogram()
	registry.MustRegister(m.MessageHistogram)

	return *m
}
