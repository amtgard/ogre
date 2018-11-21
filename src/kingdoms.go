package main

import (
	"database/sql"
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

// Event is an Amtgard event
type Event struct {
	EventID               int            `db:"event_id"`
	KingdomID             int            `db:"kingdom_id"`
	ParkID                int            `db:"park_id"`
	MundaneID             int            `db:"mundane_id"`
	UnitID                int            `db:"unit_id"`
	Name                  string         `db:"name"`
	HasHeraldry           bool           `db:"has_heraldry"`
	Modified              string         `db:"modified"`
	KingdomName           string         `db:"kingdom_name"`
	ParkName              sql.NullString `db:"park_name"`
	Persona               sql.NullString `db:"persona"`
	EventStart            string         `db:"event_start"`
	EventCalendarDetailID int            `db:"event_calendardetail_id"`
	UnitName              sql.NullString `db:"unit_name"`
	ShortDescription      string         `db:"short_description"`
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

func kingdomEventsShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	events := []Event{}

	sql := `
	SELECT
	DISTINCT
	e.*,
	k.name AS kingdom_name,
	p.name AS park_name,
	m.persona,
	cd.event_start,
	cd.event_calendardetail_id,
	u.name AS unit_name,
	SUBSTRING(cd.description, 1, 100) AS short_description
	FROM ork_event e
	LEFT JOIN ork_kingdom k ON k.kingdom_id = e.kingdom_id
	LEFT JOIN ork_park p ON p.park_id = e.park_id
	LEFT JOIN ork_mundane m ON m.mundane_id = e.mundane_id
	LEFT JOIN ork_event_calendardetail cd ON e.event_id = cd.event_id
	LEFT JOIN ork_unit u ON e.unit_id = u.unit_id
	WHERE
	e.kingdom_id = ?
	AND e.park_id = 0
	AND cd.event_start IS NOT NULL
	AND cd.event_start > DATE_ADD(NOW(), INTERVAL - 7 DAY)
	AND cd.current = 1
	ORDER BY
	cd.event_start,
	kingdom_name,
	park_name,
	e.name`

	db := dbInit()

	err := db.Select(&events, sql, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(events)
}
