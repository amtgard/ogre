package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// Park is an Amtgard park
type Park struct {
	ID            int     `db:"park_id"`
	KingdomID     int     `db:"kingdom_id"`
	Name          string  `db:"name"`
	Abbreviation  string  `db:"abbreviation"`
	HasHeraldry   bool    `db:"has_heraldry"`
	URL           string  `db:"url"`
	ParkTitleID   int     `db:"parktitle_id"`
	Active        string  `db:"active"`
	Address       string  `db:"address"`
	City          string  `db:"city"`
	Province      string  `db:"province"`
	PostalCode    string  `db:"postal_code"`
	GoogleGeocode string  `db:"google_geocode"`
	Latitude      float64 `db:"latitude"`
	Longitude     float64 `db:"longitude"`
	Location      string  `db:"location"`
	MapURL        string  `db:"map_url"`
	Description   string  `db:"description"`
	Directions    string  `db:"directions"`
	Modified      string  `db:"modified"`
}

func parkList(w http.ResponseWriter, r *http.Request) {
	parks := []Park{}

	db := dbInit()

	err := db.Select(&parks, "SELECT * FROM ork_park ORDER BY name ASC")
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(parks)
}

func parkShow(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "parkID")
	park := Park{}

	db := dbInit()

	err := db.Get(&park, "SELECT * FROM ork_park WHERE park_id=?", id)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(park)
}
