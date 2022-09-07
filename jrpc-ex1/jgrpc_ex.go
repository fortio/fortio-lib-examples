package main

import (
	"flag"
	"fmt"

	"fortio.org/fortio/jrpc"
	"fortio.org/fortio/version"
)

func main() {
	flag.Parse()
	// Assume some fortio echo debug server is running.
	code, bytes, err := jrpc.FetchURL("http://localhost:8080/debug")
	_, _, fullV := version.FromBuildInfo()
	fmt.Print(fullV)
	fmt.Printf("err %v, code %d, bytes:\n%s", err, code, string(bytes))
}
