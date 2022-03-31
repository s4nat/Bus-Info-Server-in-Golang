package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var busstopList = []string{
	"378204", "383050", "378202", "383049", "382998", "378237", "378233", "378230",
	"378229", "378228", "378227", "382995", "378224", "378226", "383010", "383009",
	"383006", "383004", "378234", "383003", "378222", "383048", "378203", "382999",
	"378225", "383014", "383013", "383011", "377906", "383018", "383015", "378207",
}

var buslineList = []string{"44478", "44479", "44480", "44481"}

// serve home - displays welcome message
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to bus timing and location API</h1>"))
}

// returns list of busses arriving at busstop
func serveBusstop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	returnMap := make(map[string][]string)
	returnMap["List of bus stops"] = busstopList
	json.NewEncoder(w).Encode(returnMap)
}

// serves full busstop json
func serveBusstopId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// get id from request
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)
	json.NewEncoder(w).Encode(GetTimings(idInt))

}

// returns name of busstop via ID
func getBusstopName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// get busstop id from request
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)
	busstopStruc := GetTimings(idInt)
	returnMap := make(map[string]string)
	returnMap["Name of bus stop"] = busstopStruc.BusstopName
	json.NewEncoder(w).Encode(returnMap)
}

// returns location (geometry) of busstop via ID
func getBusstopLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// get busstop id from request
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)
	busstopStruc := GetTimings(idInt)
	returnMap := make(map[string][]Geometry)
	returnMap["Location of bus stop"] = busstopStruc.Geometry
	json.NewEncoder(w).Encode(returnMap)
}

// returns number of incoming busses via busstop ID
func getIncomingLength(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// get busstop id from request
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)
	busstopStruc := GetTimings(idInt)
	returnMap := make(map[string]int)
	returnMap["Number of incoming vehicles"] = len(busstopStruc.IncomingVehicles)
	json.NewEncoder(w).Encode(returnMap)
}

// returns list of Busline IDs
func serveBusline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	returnMap := make(map[string][]string)
	returnMap["List of bus stops"] = buslineList
	json.NewEncoder(w).Encode(returnMap)
}

// returns full bus line json
func serveBuslineId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// get id from request
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)
	json.NewEncoder(w).Encode(GetLocations(idInt))
}

// returns name of bus line via ID
func serveBuslineName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// get id from request
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)
	buslineStruc := GetLocations(idInt)
	returnMap := make(map[string]string)
	returnMap["Name of bus line"] = buslineStruc.RouteName
	json.NewEncoder(w).Encode(returnMap)
}

// returns list of busses on bus line via ID
func getBuslineBusList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// get id from request
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)
	buslineStruc := GetLocations(idInt)
	returnMap := make(map[string]any)
	returnMap["Name of bus line"] = buslineStruc.RouteName
	returnMap["List of buses in bus line"] = buslineStruc.Vehicles
	json.NewEncoder(w).Encode(returnMap)
}

// returns current location of bus via bus stop ID and bus ID
func getBusLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// get id from request
	params := mux.Vars(r)
	id := params["id"]
	bus_id := params["bus_id"]
	idInt, _ := strconv.Atoi(id)
	bus_idInt, _ := strconv.Atoi(bus_id)
	buslineStruc := GetLocations(idInt)
	returnMap := make(map[string]any)
	returnMap["ID of bus"] = bus_idInt
	for _, vehicle := range buslineStruc.Vehicles {
		if vehicle.VehicleId == bus_idInt {
			returnMap["Location of bus"] = vehicle.Position
		}
	}
	json.NewEncoder(w).Encode(returnMap)
}
