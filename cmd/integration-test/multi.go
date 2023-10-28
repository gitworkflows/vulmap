package main

import (
	"github.com/khulnasoft-lab/vulmap/v3/pkg/testutils"
)

var multiProtoTestcases = []TestCaseInfo{
	{Path: "protocols/multi/dynamic-values.yaml", TestCase: &multiProtoDynamicExtractor{}},
	{Path: "protocols/multi/evaluate-variables.yaml", TestCase: &multiProtoDynamicExtractor{}},
	{Path: "protocols/multi/exported-response-vars.yaml", TestCase: &multiProtoDynamicExtractor{}},
}

type multiProtoDynamicExtractor struct{}

// Execute executes a test case and returns an error if occurred
func (h *multiProtoDynamicExtractor) Execute(templatePath string) error {
	results, err := testutils.RunVulmapTemplateAndGetResults(templatePath, "blog.khulnasoft-lab.io", debug)
	if err != nil {
		return err
	}
	return expectResultsCount(results, 1)
}
