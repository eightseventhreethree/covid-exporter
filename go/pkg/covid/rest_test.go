package covid

import (
	"fmt"
	"testing"
	"time"

	"github.com/eightseventhreethree/covid-exporter/v2/pkg/handlers"
	"github.com/stretchr/testify/require"
)

var CLIENT = NewClient(&ClientOpts{
	BaseURL:    "https://disease.sh",
	Timeout:    time.Second.Round(30),
	RetryLimit: 3,
	RetryDelay: time.Second.Round(3),
})

func TestGetStates(t *testing.T) {
	resp, err := CLIENT.GetStates()
	handlers.CheckErrLogT(t, err, "TestGetStates")
	count := len(resp)
	// more than just US states appear in this list i.e "Diamond Princess Ship"
	require.GreaterOrEqual(t, count, 50, "failed to get expected number of states")
	if err != nil {
		fmt.Printf("%+v\n", resp)
	}
}
