package covid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/eightseventhreethree/covid-exporter/v2/pkg/handlers"
	"github.com/hashicorp/go-retryablehttp"
)

type ClientOpts struct {
	BaseURL    string
	Timeout    time.Duration
	RetryLimit int
	RetryDelay time.Duration
}

type client struct {
	ApiKey     string
	ApiSecret  string
	BaseURL    string
	HTTPClient *retryablehttp.Client
}

type Client interface {
	GetStates() (StatesResponse, error)
}

func NewClient(opts *ClientOpts) Client {
	c := retryablehttp.NewClient()
	c.RetryMax = opts.RetryLimit
	c.RetryWaitMin = opts.RetryDelay
	return &client{
		BaseURL:    opts.BaseURL,
		HTTPClient: c,
	}
}

func (c *client) httpRequest(req *retryablehttp.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return res, err
}

func (c *client) getUmarshalAndStore(u string, s interface{}) error {
	endpoint := fmt.Sprintf("%s%s", c.BaseURL, u)
	req, err := retryablehttp.NewRequest("GET", endpoint, nil)
	handlers.CheckErrLog(err, "Failed to create NewRequest")

	resp, err := c.httpRequest(req)
	handlers.CheckErrPanic(err)
	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	handlers.CheckErrLog(err, "failed to read resp.Body")

	if err := json.Unmarshal(responseBytes, &s); err != nil {
		return err
	}
	handlers.CheckErrLog(err, "failed to Unmarshal responseBytes")

	return err
}

/*
GetStates https://disease.sh/docs/#/COVID-19%3A%20Worldometers/get_v3_covid_19_states
Returns []string of COVID-19 totals for all US States
*/
func (c *client) GetStates() (StatesResponse, error) {
	states := StatesResponse{}
	err := c.getUmarshalAndStore(V3States.toString(), &states)
	handlers.CheckErrLog(err, "failed GetSymbols()")
	return states, err
}
