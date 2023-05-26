package covid

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

//type Client interface{}

type Gauges struct {
	Cases               prometheus.GaugeVec
	TodayCases          prometheus.GaugeVec
	Deaths              prometheus.GaugeVec
	TodayDeaths         prometheus.GaugeVec
	Active              prometheus.GaugeVec
	CasesPerOneMillion  prometheus.GaugeVec
	DeathsPerOneMillion prometheus.GaugeVec
	Tests               prometheus.GaugeVec
	TestsPerOneMillion  prometheus.GaugeVec
	Population          prometheus.GaugeVec
}

func NewRegistry(gauges Gauges) *prometheus.Registry {
	registry := prometheus.NewRegistry()
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		gauges.Cases,
		gauges.TodayCases,
		gauges.Deaths,
		gauges.TodayDeaths,
		gauges.Active,
		gauges.CasesPerOneMillion,
		gauges.DeathsPerOneMillion,
		gauges.Tests,
		gauges.TestsPerOneMillion,
		gauges.Population,
	)
	return registry
}

func NewGauges() Gauges {
	var labelNames = []string{"state"}
	return Gauges{
		Cases: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "Cases",
			Help:      "Cases",
		},
			labelNames),
		TodayCases: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "TodayCases",
			Help:      "Today cases",
		}, labelNames),
		Deaths: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "Deaths",
			Help:      "Deaths",
		}, labelNames),
		TodayDeaths: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "TodayDeaths",
			Help:      "TodayDeaths",
		}, labelNames),
		Active: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "Active",
			Help:      "Active",
		}, labelNames),
		CasesPerOneMillion: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "CasesPerOneMillion",
			Help:      "CasesPerOneMillion",
		}, labelNames),
		DeathsPerOneMillion: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "DeathsPerOneMillion",
			Help:      "DeathsPerOneMillion",
		}, labelNames),
		Tests: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "Tests",
			Help:      "Tests",
		}, labelNames),
		TestsPerOneMillion: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "TestsPerOneMillion",
			Help:      "TestsPerOneMillion",
		}, labelNames),
		Population: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "covid_19",
			Name:      "Population",
			Help:      "Population",
		}, labelNames),
	}
}
