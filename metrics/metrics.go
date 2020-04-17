package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//IncUniqueIDsGenerated - Increment number of unique IDs generated so far
func IncUniqueIDsGenerated() {
	opsProcessed.Inc()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "num_unique_ids_generated",
		Help: "The total number of unique ids generated",
	})
)
