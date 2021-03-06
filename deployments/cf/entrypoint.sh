#!/bin/bash

set -e -u

if [ -z "${PORT}" ]; then
  echo "Error: No PORT found" >&2
  exit 1
fi
if [ -z "${VCAP_SERVICES}" ]; then
  echo "Error: No VCAP_SERVICES found" >&2
  exit 1
fi

# S3
s3_credentials="$(echo "${VCAP_SERVICES}" | jq -r '.["dynstrg"][0].credentials // ""')"
if [ -z "${s3_credentials}" ]; then
  echo "Error: Please bind an S3 service" >&2
  exit 1
fi
s3_endpoint="$(echo "${s3_credentials}" | jq -r '.accessHost // ""')"
s3_endpoint="${s3_endpoint#'https://'}"
s3_access_key_id="$(echo "${s3_credentials}" | jq -r '.accessKey // ""')"
s3_secret_access_key="$(echo "${s3_credentials}" | jq -r '.sharedSecret // ""')"

# Run binary
ENDPOINT="${s3_endpoint}" ACCESS_KEY_ID="${s3_access_key_id}" SECRET_ACCESS_KEY="${s3_secret_access_key}" ./s3manager
