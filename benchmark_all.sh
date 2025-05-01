#!/bin/bash
set -x
set -e

ENDPOINTS=(
  "/json"
  "/jsoniter"
  "/gojson"
  "/msgpack"
  "/protobuf"
  "/unmarshal/json"
  "/unmarshal/jsoniter"
  "/unmarshal/gojson"
  "/unmarshal/msgpack"
  "/unmarshal/protobuf"
)

echo "Benchmarking 12 endpoints with 100k requests each (100 concurrent)..."
echo

for endpoint in "${ENDPOINTS[@]}"; do
  echo ">>> Benchmarking $endpoint"
  hey -n 100000 -c 100 "http://localhost:8888$endpoint"
  echo
done
