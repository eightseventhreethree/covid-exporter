package covid

type StatesResponse []struct {
	State               string `json:"state"`
	Updated             int    `json:"updated"`
	Cases               int    `json:"cases"`
	TodayCases          int    `json:"todayCases"`
	Deaths              int    `json:"deaths"`
	TodayDeaths         int    `json:"todayDeaths"`
	Active              int    `json:"active"`
	CasesPerOneMillion  int    `json:"casesPerOneMillion"`
	DeathsPerOneMillion int    `json:"deathsPerOneMillion"`
	Tests               int    `json:"tests"`
	TestsPerOneMillion  int    `json:"testsPerOneMillion"`
	Population          int    `json:"population"`
}
