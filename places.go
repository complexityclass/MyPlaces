package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CurrencyType string

const (
	CurrencyRUB CurrencyType = "RUB"
	CurrencyEUR CurrencyType = "EUR"
	CurrencyUSD CurrencyType = "USD"
)

type Location struct {
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Place struct {
	Name      string        `json:"place"`
	Date      time.Time     `json:"date"`
	Duration  time.Duration `json:"duration"`
	Cost      float32       `json:"cost"`
	Currency  CurrencyType  `json:"currency"`
	Transport string        `json:"transport"`
	Comment   string        `json:"comment"`
	Location
}

func GetPlaces(w http.ResponseWriter, _ *http.Request) {
	var data = generateData()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Print(err)
	}
}

func generateData() (places []Place) {
	var t = time.Now()
	places = make([]Place, 0)
	for i := 0; i < 50; i++ {
		name := "SHOP Korablik #" + strconv.Itoa(i)
		currTime := time.Now()
		duration := time.Since(t)
		cost := float32(i) * 12.3
		lat := 55.940879 + float64(i)*0.3
		lot := 37.490907 + float64(i)*0.3
		var curr CurrencyType
		if i%3 == 0 {
			curr = CurrencyRUB
		} else if i%3 == 1 {
			curr = CurrencyEUR
		} else {
			curr = CurrencyUSD
		}

		var place = Place{
			Name:      name,
			Date:      currTime,
			Duration:  duration,
			Cost:      cost,
			Currency:  curr,
			Transport: "Bus",
			Comment:   "Good kids shop. Bought gifts there",
			Location: Location{
				Address:   "Likhachevskoye Shosse, 11а, Dolgoprudny, Moskovskaya oblast', 141707",
				Latitude:  lat,
				Longitude: lot,
			},
		}

		places = append(places, place)
	}

	return
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/places", GetPlaces).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
