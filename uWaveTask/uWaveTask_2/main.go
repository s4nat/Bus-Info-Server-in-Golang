package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")

	r.HandleFunc("/busstop", serveBusstop).Methods("GET")
	r.HandleFunc("/busstop/{id}", serveBusstopId).Methods("GET")
	r.HandleFunc("/busstop/{id}/name", getBusstopName).Methods("GET")
	r.HandleFunc("/busstop/{id}/location", getBusstopLocation).Methods("GET")
	r.HandleFunc("/busstop/{id}/incomingNo", getIncomingLength).Methods("GET")

	r.HandleFunc("/busline", serveBusline).Methods("GET")
	r.HandleFunc("/busline/{id}", serveBuslineId).Methods("GET")
	r.HandleFunc("/busline/{id}/name", serveBuslineName).Methods("GET")
	r.HandleFunc("/busline/{id}/vehicles", getBuslineBusList).Methods("GET")
	r.HandleFunc("/busline/{id}/{bus_id}/location", getBusLocation).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}
