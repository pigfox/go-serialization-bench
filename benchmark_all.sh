#!/bin/bash
set -x
set -e
clear

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

echo "Benchmarking 10 endpoints with 1M requests each (1000 concurrent)..."
echo

for endpoint in "${ENDPOINTS[@]}"; do
  echo ">>> Benchmarking $endpoint"
  hey -n 1000000 -c 1000 "http://localhost:8888$endpoint"
  echo
done
