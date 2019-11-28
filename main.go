package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main(){
	// Exporting in port 9338 because it is one of the 
	// free exporter ports. For more info visit:
	// https://github.com/prometheus/prometheus/wiki/Default-port-allocations
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Begining to serve on port 9338")
	log.Fatal(http.ListenAndServe(":9338",nil))
}