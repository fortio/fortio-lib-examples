package main

import (
	"flag"
	"log"
	"os"

	"fortio.org/fortio/jrpc"
	"fortio.org/fortio/version"
)

func main() {
	// Assumes some fortio echo debug server is running. pass -url localhost:8080/?status=503 for error case for instance.
	url := flag.String("url", "http://localhost:8080/debug", "url to fetch")
	flag.Parse()
	code, bytes, err := jrpc.FetchURL(*url)
	_, _, fullV := version.FromBuildInfo()
	log.Printf("Running: %s", fullV)
	log.Printf("err %v, code %d, bytes (len %d):", err, code, len(bytes))
	os.Stdout.Write(bytes)
}
