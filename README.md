# covid-exporters for Prometheus

# Go version was written to improve performance
# Python is still functional

Uses public API available here: https://disease.sh

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

Docker container images are prefixed in tag name with language:
i.e:
```
docker pull rushsimonson/covid-exporter:go-${VERSION}
docker pull rushsimonson/covid-exporter:py-${VERSION}
```

An example Grafana dashboard for import is available in grafana folder.
