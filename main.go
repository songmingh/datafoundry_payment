package main

import (
	"net/http"

	"github.com/zonesan/clog"
)

func main() {

	router := createRouter()

	//clog.SetLogLevel(clog.LOG_LEVEL_DEBUG)
	clog.Fatal(http.ListenAndServe(":8080", router))
}
