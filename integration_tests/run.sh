#!/bin/bash

echo "::group::Build vulmap"
rm integration-test vulmap 2>/dev/null
cd ../cmd/vulmap
go build -race .
mv vulmap ../../integration_tests/vulmap 
echo "::endgroup::"

echo "::group::Build vulmap integration-test"
cd ../integration-test
go build
mv integration-test ../../integration_tests/integration-test 
cd ../../integration_tests
echo "::endgroup::"

echo "::group::Installing vulmap templates"
./vulmap -update-templates
echo "::endgroup::"

./integration-test
if [ $? -eq 0 ]
then
  exit 0
else
  exit 1
fi
