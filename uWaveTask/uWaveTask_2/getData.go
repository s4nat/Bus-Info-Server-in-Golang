package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetTimings(busStopId int) Busstop {
	busStopIdStr := strconv.FormatInt(int64(busStopId), 10)
	url := "https://dummy.uwave.sg/busstop/" + busStopIdStr

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//We Read the response body on the line below.
	body, _ := ioutil.ReadAll(resp.Body)

	var busstop Busstop

	json.Unmarshal(body, &busstop)
	return busstop
}

func GetLocations(busLineId int) Busline {
	busLineIdStr := strconv.FormatInt(int64(busLineId), 10)
	url := "https://dummy.uwave.sg/busline/" + busLineIdStr

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//We Read the response body on the line below.
	body, _ := ioutil.ReadAll(resp.Body)

	var busline Busline

	json.Unmarshal(body, &busline)
	return busline
}
