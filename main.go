package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/proto"

	pb "serialization-bench/pb"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type Person struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Addresses []Address `json:"addresses"`
}

// EasyJSON support
//
//go:generate easyjson -all main.go
type PersonEasyJSON Person

func newPerson() Person {
	return Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Addresses: []Address{
			{"123 Main St", "Springfield", "12345"},
			{"456 Market St", "Metropolis", "67890"},
		},
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// ========== Marshal Handlers ==========

func jsonEncodingHandler(w http.ResponseWriter, _ *http.Request) {
	data, _ := json.Marshal(newPerson())
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func jsoniterHandler(w http.ResponseWriter, _ *http.Request) {
	jsoni := jsoniter.ConfigCompatibleWithStandardLibrary
	data, _ := jsoni.Marshal(newPerson())
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func gojsonHandler(w http.ResponseWriter, _ *http.Request) {
	data, _ := gojson.Marshal(newPerson())
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func msgpackHandler(w http.ResponseWriter, _ *http.Request) {
	data, _ := msgpack.Marshal(newPerson())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}

func protobufHandler(w http.ResponseWriter, _ *http.Request) {
	personPB := &pb.Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Addresses: []*pb.Address{
			{Street: "123 Main St", City: "Springfield", ZipCode: "12345"},
			{Street: "456 Market St", City: "Metropolis", ZipCode: "67890"},
		},
	}
	data, _ := proto.Marshal(personPB)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}

// ========== Unmarshal Handlers ==========

func unmarshalJSON(w http.ResponseWriter, r *http.Request) {
	var p Person
	json.NewDecoder(r.Body).Decode(&p)
	writeJSON(w, p)
}

func unmarshalJSONIter(w http.ResponseWriter, r *http.Request) {
	var p Person
	jsoni := jsoniter.ConfigCompatibleWithStandardLibrary
	jsoni.NewDecoder(r.Body).Decode(&p)
	writeJSON(w, p)
}

func unmarshalGoJSON(w http.ResponseWriter, r *http.Request) {
	var p Person
	gojson.NewDecoder(r.Body).Decode(&p)
	writeJSON(w, p)
}

func unmarshalMsgpack(w http.ResponseWriter, r *http.Request) {
	var p Person
	data, _ := io.ReadAll(r.Body)
	msgpack.Unmarshal(data, &p)
	writeJSON(w, p)
}

func unmarshalProtobuf(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var personPB pb.Person
	proto.Unmarshal(data, &personPB)
	writeJSON(w, personPB)
}

// ========== Main ==========

func main() {
	http.HandleFunc("/json", jsonEncodingHandler)
	http.HandleFunc("/jsoniter", jsoniterHandler)
	http.HandleFunc("/gojson", gojsonHandler)
	http.HandleFunc("/msgpack", msgpackHandler)
	http.HandleFunc("/protobuf", protobufHandler)

	http.HandleFunc("/unmarshal/json", unmarshalJSON)
	http.HandleFunc("/unmarshal/jsoniter", unmarshalJSONIter)
	http.HandleFunc("/unmarshal/gojson", unmarshalGoJSON)
	http.HandleFunc("/unmarshal/msgpack", unmarshalMsgpack)
	http.HandleFunc("/unmarshal/protobuf", unmarshalProtobuf)
	port := ":8888"

	fmt.Println("Server listening on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
