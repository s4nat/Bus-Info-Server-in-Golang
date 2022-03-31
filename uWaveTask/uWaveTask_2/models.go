package main

type Busstop struct {
	BusstopId        int        `json:"id,omitempty"`
	BusstopName      string     `json:"name,omitempty"`
	Geometry         []Geometry `json:"geometry,omitempty"`
	IncomingVehicles []Vehicle  `json:"forecast,omitempty"`
}

type Vehicle struct {
	VehicleId        int       `json:"vehicle_id,omitempty"`
	RouteVariantId   int       `json:"rv_id,omitempty"`
	VehicleReg       string    `json:"vehicle,omitempty"`
	RegistrationCode string    `json:"registration_code,omitempty"`
	Position         *Position `json:"position,omitempty"`
	ForecastSeconds  float64   `json:"forecast_seconds,omitempty"`
}

type Busline struct {
	BuslineId int       `json:"id,omitempty"`
	RouteName string    `json:"routename,omitempty"`
	Vehicles  []Vehicle `json:"vehicles,omitempty"`
}

type Position struct {
	Bearing  int    `json:"bearing,omitempty"`
	DeviceTs int    `json:"device_ts,omitempty"`
	Lat      string `json:"lat,omitempty"`
	Lon      string `json:"lon,omitempty"`
	Speed    int    `json:"speed,omitempty"`
}

type Geometry struct {
	Lat string `json:"lat,omitempty"`
	Lon string `json:"lon,omitempty"`
}
