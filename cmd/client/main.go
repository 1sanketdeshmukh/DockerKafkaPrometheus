package main

import (
	"KafkaDocker/internal/pkg/kafkapkg"
	"KafkaDocker/internal/pkg/metric"
	"KafkaDocker/internal/pkg/prometheus"
	"context"
)

func main() {

	// Prometheus
	pro := prometheus.New(true)

	// Metric
	mtr := metric.New(pro.Registry())

	// Kafka
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go kafkapkg.Produce(ctx)
	kafkapkg.Consume(ctx)

}
