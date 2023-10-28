package main

import (
	"github.com/khulnasoft-lab/vulmap/v3/pkg/testutils"
)

var whoisTestCases = []TestCaseInfo{
	{Path: "protocols/whois/basic.yaml", TestCase: &whoisBasic{}},
}

type whoisBasic struct{}

// Execute executes a test case and returns an error if occurred
func (h *whoisBasic) Execute(filePath string) error {
	results, err := testutils.RunVulmapTemplateAndGetResults(filePath, "https://example.com", debug)
	if err != nil {
		return err
	}
	return expectResultsCount(results, 1)
}
