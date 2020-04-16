package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

var uniqueID int

func main() {
	flag.Parse()
	http.HandleFunc("/getNewId", getUniqueID)

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
	glog.Fatal("Both listeners started succesfully")
	glog.Flush()
	select {}
}

func getUniqueID(w http.ResponseWriter, req *http.Request) {
	uniqueID++
	fmt.Fprintf(w, "{\"id\":%d}", uniqueID)
}
