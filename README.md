# go-serialization-bench
This repository demonstrates and benchmarks multiple serialization and deserialization strategies in Go, exposed via HTTP endpoints. It includes standard and third-party libraries for JSON, MessagePack, and Protocol Buffers.

## 🧩 Features

- Serializes and deserializes a complex Go struct using:
  - `encoding/json` (standard library)
  - `jsoniter`
  - `go-json`
  - `easyjson`
  - `msgpack`
  - `protobuf` (v2 API)
- 12 HTTP endpoints:
  - 6 for `marshal` (`GET`)
  - 6 for `unmarshal` (`POST`)
- Includes a shell script to benchmark each endpoint using [`hey`](https://github.com/rakyll/hey)

## 📦 Structure


