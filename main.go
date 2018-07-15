package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var drinkers []Drinker = []Drinker{
	Drinker{Name: "Dicker Nagus", Drinks: []Drink{Drink{Name: "Mate", Amount: 9001}}},
	Drinker{Name: "Gast", Drinks: []Drink{Drink{Name: "Flens", Amount: 42}}},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/drinkers", GetDrinkers).Methods("GET")
	router.HandleFunc("/drinker/{name}", GetDrinker).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetDrinkers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(drinkers)
}
func GetDrinker(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range drinkers {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return

		}
	}
}

type Drinker struct {
	Name string `json:"name"`
	Drinks []Drink `json:"drinks"`
}

type Drink struct {
	Name   string `json:"name"`
	Amount int    `json:"int"`
}
