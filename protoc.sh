#!/bin/bash
set -x
set -e

mkdir -p pb
protoc \
  --go_out=pb \
  --go_opt=paths=source_relative \
  person.proto

