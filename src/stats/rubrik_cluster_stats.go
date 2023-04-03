package stats

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
)

var (
	rubrikClusterStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rubrik_cluster_status",
			Help: "Rubrik cluster status.",
		},
		[]string{
			"clusterName",
		},
	)
)

func init() {
	prometheus.MustRegister(rubrikClusterStatus)
}

func GetClusterStats(rubrik *rubrikcdm.Credentials, clusterName string) {
	clusterDetails, err := rubrik.Get("v1", "/cluster/me", 60)
	if err != nil {
		return
	}

	clusterStatus := clusterDetails.(map[string]interface{})["status"].(string)

	var statusValue float64
	if clusterStatus == "OK" {
		statusValue = 1
	} else {
		statusValue = 0
	}

	rubrikClusterStatus.WithLabelValues(clusterName).Set(statusValue)
}
