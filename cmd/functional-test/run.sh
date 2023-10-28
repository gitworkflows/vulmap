#!/bin/bash

# reading os type from arguments
CURRENT_OS=$1

if [ "${CURRENT_OS}" == "windows-latest" ];then
    extension=.exe
fi

echo "::group::Building functional-test binary"
go build -o functional-test$extension
echo "::endgroup::"

echo "::group::Building Vulmap binary from current branch"
go build -o vulmap_dev$extension ../vulmap
echo "::endgroup::"

echo "::group::Installing vulmap templates"
./vulmap_dev$extension -update-templates
echo "::endgroup::"

echo "::group::Building latest release of vulmap"
go build -o vulmap$extension -v github.com/khulnasoft-lab/vulmap/v3/cmd/vulmap
echo "::endgroup::"

echo 'Starting Vulmap functional test'
./functional-test$extension -main ./vulmap$extension -dev ./vulmap_dev$extension -testcases testcases.txt
