package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

func messageHistogram() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "client",
		Name:      "message_histogram",
		Help:      "message duration (ms)",
		Buckets:   []float64{10, 50, 90, 130, 170, 210, 250, 290, 330},
		// This is same as prometheus.LinearBuckets(10, 40, 9)
		// 9 buckets starting from 10 increased by 40
	}, []string{"event"})
}

// Another example:
// 4 buckets starting from 1 multiplied by 3 between. e.g. 1, 3, 9, 27
// prometheus.ExponentialBuckets(1, 3, 9)
