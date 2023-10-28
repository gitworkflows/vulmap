package main

import vulmap "github.com/khulnasoft-lab/vulmap/v3/lib"

func main() {
	ne, err := vulmap.NewVulmapEngine(
		vulmap.WithTemplateFilters(vulmap.TemplateFilters{Tags: []string{"oast"}}),
		vulmap.EnableStatsWithOpts(vulmap.StatsOptions{MetricServerPort: 6064}), // optionally enable metrics server for better observability
	)
	if err != nil {
		panic(err)
	}
	// load targets and optionally probe non http/https targets
	ne.LoadTargets([]string{"http://honey.scanme.sh"}, false)
	err = ne.ExecuteWithCallback(nil)
	if err != nil {
		panic(err)
	}
	defer ne.Close()
}
