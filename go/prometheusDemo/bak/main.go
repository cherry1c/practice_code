package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func basicInfo() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func main() {

}
