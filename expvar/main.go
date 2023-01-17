package main

import (
	_ "expvar"
	"net/http"
)

func main() {
	_ = http.ListenAndServe(":6060", http.DefaultServeMux)
}
