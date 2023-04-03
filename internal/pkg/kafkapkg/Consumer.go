package kafkapkg

import (
	"KafkaDocker/internal/pkg/metric"
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, mtr metric.Metric) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		//Brokers: []string{broker1Address, broker2Address, broker3Address},
		Brokers: []string{broker1Address},
		Topic:   topic,
		GroupID: "my-group",
	})

	count := 0
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		// Add metrics for a message if message is read successfully

		if count%3 == 0 {
			mtr.KafkaDropMessageCounter.WithLabelValues("dropped message").Inc()
		} else {
			mtr.KafkaMessageCounter.WithLabelValues("message received").Inc()
		}

		start := time.Now()
		dur := float64(time.Since(start).Milliseconds())
		mtr.MessageHistogram.WithLabelValues("message Histogram").Observe(dur)

		mtr.MessageGuage.Add(float64(count))

		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))

		count++
	}
}
