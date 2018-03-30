package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/ryutah/gke-endpoints-sample/models"
)

func main() {
	flag.Parse()

	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1")
	v1.Path("/echo").Methods("POST").HandlerFunc(echo)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	msg := new(models.EchoMessage)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read body : %#v", err), http.StatusInternalServerError)
		return
	}
	if err := msg.UnmarshalBinary(body); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse body : %#v", err), http.StatusBadRequest)
		return
	}

	glog.Infof("This is Info; Receive Message: %#v", msg)
	glog.Warningf("This is Warn; Receive Message: %#v", msg)
	glog.Errorf("This is Error; Receive Message: %#v", msg)

	resp, err := msg.MarshalBinary()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal body : %#v", err), http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
