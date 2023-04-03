package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

func kafkaMessageCounter() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Name:      "message_counter",
		Help:      "Number of Kafka messages",
	}, []string{"event"})
}

func kafkaDropMessageCounter() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Name:      "dropped_message_counter",
		Help:      "Number of dropped messages",
	}, []string{"event"})
}
