import time
import requests
import datetime
from prometheus_client.core import GaugeMetricFamily, REGISTRY, CounterMetricFamily
from prometheus_client import start_http_server


class CustomCollector(object):
    def collect(self):
        base_metric_name = "covid_19_"
        covid_states = requests.get("https://disease.sh/v3/covid-19/states")
        json_covid_states = covid_states.json()

        # cases
        metric_name = base_metric_name + "cases"
        c = CounterMetricFamily(metric_name, "Help text", labels=["state"])
        for state in json_covid_states:
            c.add_metric([state["state"]], state["cases"])
        yield c

        # todays_cases
        metric_name = base_metric_name + "todayCases"
        c = CounterMetricFamily(metric_name, "Help text", labels=["state"])
        for state in json_covid_states:
            today_cases = state.get("todayCases")
            if today_cases is not None:
                c.add_metric([state["state"]], today_cases)
        yield c

        # deaths
        metric_name = base_metric_name + "deaths"
        c = CounterMetricFamily(metric_name, "Help text", labels=["state"])
        for state in json_covid_states:
            c.add_metric([state["state"]], state["deaths"])
        yield c

        # todayDeaths
        metric_name = base_metric_name + "todayDeaths"
        c = CounterMetricFamily(metric_name, "Help text", labels=["state"])
        for state in json_covid_states:
            today_deaths = state.get("todayDeaths")
            if today_deaths is not None:
                c.add_metric([state["state"]], today_deaths)
        yield c

        # active
        metric_name = base_metric_name + "active"
        c = CounterMetricFamily(metric_name, "Help text", labels=["state"])
        for state in json_covid_states:
            active = state.get("active")
            if active is not None:
                c.add_metric([state["state"]], active)
        yield c

        # tests
        metric_name = base_metric_name + "tests"
        c = CounterMetricFamily(metric_name, "Help text", labels=["state"])
        for state in json_covid_states:
            tests = state.get("tests")
            if tests is not None:
                c.add_metric([state["state"]], tests)
        yield c

        # testsPerOneMillion
        metric_name = base_metric_name + "testsPerOneMillion"
        c = CounterMetricFamily(metric_name, "Help text", labels=["state"])
        for state in json_covid_states:
            tests_per = state.get("testsPerOneMillion")
            if tests_per is not None:
                c.add_metric([state["state"]], tests_per)
        yield c


if __name__ == "__main__":
    start_http_server(8000)
    REGISTRY.register(CustomCollector())
    print("Registered custom metrics", flush=True)
    while True:
        x = datetime.datetime.now()
        print("Pulled new metrics", x, flush=True)
        time.sleep(600)
