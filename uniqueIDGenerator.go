package main

import (
	"flag"
	"fmt"
	"learnGo/golangBeginnerExercises/UniqueIDGenerator/metrics"
	"net/http"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var uniqueID int

func main() {
	flag.Parse()
	http.HandleFunc("/getNewId", getUniqueID)
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		err := http.ListenAndServeTLS(":8043", "certs/https-server.crt", "certs/https-server.key", nil)
		if err != nil {
			glog.Fatal("Unable to start sever: ", err.Error())
		}
	}()

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			glog.Fatal("Unable to start sever: ", err.Error())
		}
	}()
	glog.Info("Both listeners started succesfully")
	glog.Flush()
	select {}
}

func getUniqueID(w http.ResponseWriter, req *http.Request) {
	uniqueID++
	fmt.Fprintf(w, "{\"id\":%d}", uniqueID)
	metrics.IncUniqueIDsGenerated()
}
