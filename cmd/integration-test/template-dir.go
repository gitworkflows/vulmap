package main

import (
	"os"

	"github.com/khulnasoft-lab/vulmap/v3/pkg/testutils"
	errorutil "github.com/khulnasoft-lab/utils/errors"
)

var templatesDirTestCases = []TestCaseInfo{
	{Path: "protocols/dns/cname-fingerprint.yaml", TestCase: &templateDirWithTargetTest{}},
}

type templateDirWithTargetTest struct{}

// Execute executes a test case and returns an error if occurred
func (h *templateDirWithTargetTest) Execute(filePath string) error {
	tempdir, err := os.MkdirTemp("", "vulmap-update-dir-*")
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("failed to create temp dir")
	}
	defer os.RemoveAll(tempdir)

	results, err := testutils.RunVulmapTemplateAndGetResults(filePath, "8x8exch02.8x8.com", debug, "-ud", tempdir)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}