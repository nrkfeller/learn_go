package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Device struct {
	ID       string    `json:"id,omitempty"`
	Dtype    string    `json:"dtype,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type Location struct {
	Longitude int `json:"longitude,omitempty"`
	Latitude  int `json:"latitude,omitempty"`
}

var devices []Device

func main() {
	router := mux.NewRouter()

	devices = append(devices, Device{ID: "1", Dtype: "Sensor", Location: &Location{Longitude: 52345, Latitude: 75467}})
	devices = append(devices, Device{ID: "2", Dtype: "Rasberry"})

	router.HandleFunc("/device", GetDevicesEndpoint).Methods("GET")
	router.HandleFunc("/device/{id}", GetDeviceEndpoint).Methods("GET")
	router.HandleFunc("/device/{id}", CreateDeviceEndpoint).Methods("POST")
	router.HandleFunc("/device/{id}", DeleteDeviceEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}

func CreateDeviceEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var device Device
	_ = json.NewDecoder(r.Body).Decode(&device)
	device.ID = params["id"]
	devices = append(devices, device)
	json.NewEncoder(w).Encode(devices)
}

func GetDeviceEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range devices {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func GetDevicesEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(devices)
}

func DeleteDeviceEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for k, v := range devices {
		if v.ID == params["id"] {
			devices = append(devices[:k], devices[k+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(devices)
}
