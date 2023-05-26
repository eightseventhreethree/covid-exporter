package covid

type Endpoints string

const (
	V3States Endpoints = "/v3/covid-19/states"
)

func (e Endpoints) toString() string {
	return string(e)
}
