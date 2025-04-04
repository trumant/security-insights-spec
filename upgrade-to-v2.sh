#!/bin/bash
VI_INSIGHTS=$1
PROJECT_NAME=$2

# Check if insights-v1.yml exists locally first
if [ -f "insights-v1.yml" ]; then
    echo "Found local insights-v1.yml, skipping download"
else
    echo "Downloading v1 insights"
    curl $VI_INSIGHTS -o insights-v1.yml --silent
fi

# double-check that it actually conforms to the v1 schema
echo "Validating v1 insights against the v1 schema"
cue vet insights-v1.yml v1-schema.cue -d '#v1'
# convert the v1 data to v2 data
if [ $? -ne 0 ]; then
    echo "Error: v1 insights file failed schema validation. Please correct the errors in v1 data before attempting another conversion"
    exit 1
fi

echo "Converting v1 data to v2 insights and saving to security-insights.yml"
cue export .:security_insights_spec -l input: insights-v1.yml -e output --out yaml -t project=${PROJECT_NAME} > security-insights.yml

echo "Validating v2 insights against the v2 schema"
cue vet security-insights.yml schema.cue -d '#v2'

