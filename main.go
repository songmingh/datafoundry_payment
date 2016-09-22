package main

import (
	"net/http"

	"github.com/zonesan/clog"
)

func main() {

	router := createRouter()

	clog.Fatal(http.ListenAndServe(":8080", router))
}
