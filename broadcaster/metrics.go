package broadcaster

import "github.com/prometheus/client_golang/prometheus"

var (
	failedSubscribeNewHeadCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "nerif_app",
		Subsystem: "broadcaster",
		Name:      "failed_subscribe_new_head",
		Help:      "The total number of failed subscriptions for new heads",
	}, []string{"chain_id"})

	failedHealthcheckCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "nerif_app",
		Subsystem: "broadcaster",
		Name:      "failed_healthcheck",
		Help:      "The total number of failed healthchecks",
	}, []string{"chain_id"})

	resubscribeNewHeadsSubscriptionCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "nerif_app",
		Subsystem: "broadcaster",
		Name:      "reinit_new_heads",
		Help:      "The total number of re-initializing a new heads subscription",
	}, []string{"chain_id"})
)

func init() {
	prometheus.MustRegister(failedSubscribeNewHeadCounter)
	prometheus.MustRegister(failedHealthcheckCounter)
	prometheus.MustRegister(resubscribeNewHeadsSubscriptionCounter)
}
