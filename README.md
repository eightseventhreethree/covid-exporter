# covid-exporter for Prometheus

Uses public API available here: https://corona.lmao.ninja

Values:
```
# TYPE covid_19_active_total counter
# TYPE covid_19_cases_total counter
# TYPE covid_19_deaths_total counter
# TYPE covid_19_testsPerOneMillion_total counter
# TYPE covid_19_tests_total counter
# TYPE covid_19_todayCases_total counter
# TYPE covid_19_todayDeaths_total counter
```

Label example:
```
{state="Nebraska"}
```

An example Grafana dashboard for import is available in grafana folder.
