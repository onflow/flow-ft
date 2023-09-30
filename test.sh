#!/bin/bash

set -e

flow test --cover --covercode="contracts" --coverprofile="coverage.lcov" test/*_tests.cdc