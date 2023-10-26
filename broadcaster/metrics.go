package broadcaster

import "github.com/prometheus/client_golang/prometheus"

var (
	failedSubscribeNewHeadCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "nerif_app",
		Subsystem: "broadcaster",
		Name:      "failed_subscribe_new_head",
		Help:      "The total number of failed subscriptions for new heads",
	}, []string{"chain_id"})
)

func init() {
	prometheus.MustRegister(failedSubscribeNewHeadCounter)
}
