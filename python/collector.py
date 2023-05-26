import time
import requests
import sys
import datetime
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor
from prometheus_client.core import GaugeMetricFamily, REGISTRY, CounterMetricFamily
from prometheus_client import start_http_server


class CustomCollector(object):
    def __init__(self):
        self.base_metric_name = "covid_19_"
        covid_states = requests.get('https://disease.sh/v3/covid-19/states')
        self.json_covid_states = covid_states.json()
        #keys = json_covid_states[0].keys()

    def _get_metrics(self, metric_name_arg):
        metric_name = self.base_metric_name + metric_name_arg
        c = CounterMetricFamily(metric_name, 'Help text', labels=['state'])
        for state in self.json_covid_states:
          c.add_metric([state['state']], state[metric_name_arg])
        yield c

    def collect(self):
        """
        metric_name = cases
        """
        metrics_to_collect = ["cases", "todayCases"]
        with ThreadPoolExecutor() as executor:
            results = executor.map(self._get_metrics, metrics_to_collect)
        yield results

    def cases(self):
        # cases
        metric_name = base_metric_name + "cases"
        c = CounterMetricFamily(metric_name, 'Help text', labels=['state'])
        for state in json_covid_states:
          c.add_metric([state['state']], state['cases'])
        yield c

    def todays_cases(self):
        # todays_cases
        metric_name = base_metric_name + "todayCases"
        c = CounterMetricFamily(metric_name, 'Help text', labels=['state'])
        for state in json_covid_states:
          c.add_metric([state['state']], state['todayCases'])
        yield c

    def deaths(self):
        # deaths
        metric_name = base_metric_name + "deaths"
        c = CounterMetricFamily(metric_name, 'Help text', labels=['state'])
        for state in json_covid_states:
          c.add_metric([state['state']], state['deaths'])
        yield c

    def today_deaths(self):
        # todayDeaths
        metric_name = base_metric_name + "todayDeaths"
        c = CounterMetricFamily(metric_name, 'Help text', labels=['state'])
        for state in json_covid_states:
          c.add_metric([state['state']], state['todayDeaths'])
        yield c

    def active(self):
        # active
        metric_name = base_metric_name + "active"
        c = CounterMetricFamily(metric_name, 'Help text', labels=['state'])
        for state in json_covid_states:
          c.add_metric([state['state']], state['active'])
        yield c

    def test(self):
        # tests
        metric_name = base_metric_name + "tests"
        c = CounterMetricFamily(metric_name, 'Help text', labels=['state'])
        for state in json_covid_states:
          c.add_metric([state['state']], state['tests'])
        yield c

    def test_per_one_million(self):
        # testsPerOneMillion
        metric_name = base_metric_name + "testsPerOneMillion"
        c = CounterMetricFamily(metric_name, 'Help text', labels=['state'])
        for state in json_covid_states:
          c.add_metric([state['state']], state['testsPerOneMillion'])
        yield c

if __name__ == '__main__':
    start_http_server(8000)
    custom_collector = CustomCollector()
    metrics_to_collect = ["cases", "todayCases"]
    REGISTRY.register(*custom_collector)
    #for result in results:
    #    #print(*result)
    #    REGISTRY.register(*result)
    print('Registered custom metrics', flush=True)
    while True:
      x = datetime.datetime.now()
      print('Pulled new metrics', x, flush=True)
      time.sleep(600)
