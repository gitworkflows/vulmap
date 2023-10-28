package main

import (
	vulmap "github.com/khulnasoft-lab/vulmap/v3/lib"
	"github.com/remeh/sizedwaitgroup"
)

func main() {
	// create vulmap engine with options
	ne, err := vulmap.NewThreadSafeVulmapEngine()
	if err != nil {
		panic(err)
	}
	// setup sizedWaitgroup to handle concurrency
	sg := sizedwaitgroup.New(10)

	// scan 1 = run dns templates on scanme.sh
	sg.Add()
	go func() {
		defer sg.Done()
		err = ne.ExecuteVulmapWithOpts([]string{"scanme.sh"},
			vulmap.WithTemplateFilters(vulmap.TemplateFilters{ProtocolTypes: "dns"}),
		)
		if err != nil {
			panic(err)
		}
	}()

	// scan 2 = run templates with oast tags on honey.scanme.sh
	sg.Add()
	go func() {
		defer sg.Done()
		err = ne.ExecuteVulmapWithOpts([]string{"http://honey.scanme.sh"}, vulmap.WithTemplateFilters(vulmap.TemplateFilters{Tags: []string{"oast"}}))
		if err != nil {
			panic(err)
		}
	}()

	// wait for all scans to finish
	sg.Wait()
	defer ne.Close()

	// Output:
	// [dns-saas-service-detection] scanme.sh
	// [nameserver-fingerprint] scanme.sh
	// [dns-saas-service-detection] honey.scanme.sh
}
