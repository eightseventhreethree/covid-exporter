package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eightseventhreethree/covid-exporter/v2/pkg/covid"
	"github.com/eightseventhreethree/covid-exporter/v2/pkg/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	client := covid.NewClient(&covid.ClientOpts{
		BaseURL:    "https://disease.sh",
		Timeout:    time.Second.Round(30),
		RetryLimit: 3,
		RetryDelay: time.Second.Round(3),
	})

	gauges := covid.NewGauges()
	go func() {
		for {
			states, err := client.GetStates()
			handlers.CheckErrLog(err, "Failed to get states")
			for _, v := range states {
				gauges.Active.With(prometheus.Labels{"state": v.State}).Set(float64(v.Active))
				gauges.Cases.With(prometheus.Labels{"state": v.State}).Set(float64(v.Cases))
				gauges.CasesPerOneMillion.With(prometheus.Labels{"state": v.State}).Set(float64(v.CasesPerOneMillion))
				gauges.Deaths.With(prometheus.Labels{"state": v.State}).Set(float64(v.Deaths))
				gauges.DeathsPerOneMillion.With(prometheus.Labels{"state": v.State}).Set(float64(v.DeathsPerOneMillion))
				gauges.Population.With(prometheus.Labels{"state": v.State}).Set(float64(v.Population))
				gauges.Tests.With(prometheus.Labels{"state": v.State}).Set(float64(v.Tests))
				gauges.TestsPerOneMillion.With(prometheus.Labels{"state": v.State}).Set(float64(v.TestsPerOneMillion))
				gauges.TodayCases.With(prometheus.Labels{"state": v.State}).Set(float64(v.TodayCases))
				gauges.TodayDeaths.With(prometheus.Labels{"state": v.State}).Set(float64(v.TodayDeaths))
			}
			fmt.Printf("%s [INFO] Pulled new metrics.\n", time.Now().Format("2006/01/02 15:04:05"))
			time.Sleep(time.Duration(600 * time.Second))
		}
	}()
	registry := covid.NewRegistry(gauges)

	// Expose /metrics HTTP endpoint using the created custom registry.
	http.Handle(
		"/metrics", promhttp.HandlerFor(
			registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			}),
	)
	// To test: curl -H 'Accept: application/openmetrics-text' localhost:8080/metrics
	log.Fatalln(http.ListenAndServe(":8000", nil))

}
