package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Kingdom is an Amtgard kingdom
type Kingdom struct {
	ID           int    `db:"kingdom_id"`
	Name         string `db:"name"`
	Abbreviation string `db:"abbreviation"`
	HasHeraldry  bool   `db:"has_heraldry"`
	ParentID     int    `db:"parent_kingdom_id"`
	Modified     string `db:"modified"`
	Active       string `db:"active"`
}

func kingdomList(w http.ResponseWriter, r *http.Request) {
	kingdoms := []Kingdom{}

	db := dbInit()

	err := db.Select(&kingdoms, "SELECT * FROM ork_kingdom ORDER BY name ASC")
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(kingdoms)
}

func kingdomShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	kingdom := Kingdom{}

	db := dbInit()

	err := db.Get(&kingdom, "SELECT * FROM ork_kingdom WHERE kingdom_id=?", id)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(kingdom)
}
