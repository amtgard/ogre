package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Player is an Amtgard player
type Player struct {
	ID                  int            `db:"mundane_id"`
	GivenName           string         `db:"given_name"`
	Surname             string         `db:"surname"`
	OtherName           string         `db:"other_name"`
	Username            string         `db:"username"`
	Persona             string         `db:"persona"`
	Email               string         `db:"email"`
	ParkID              int            `db:"park_id"`
	KingdomID           int            `db:"kingdom_id"`
	Token               string         `db:"token"`
	Modified            string         `db:"modified"`
	Restricted          bool           `db:"restricted"`
	Waivered            bool           `db:"waivered"`
	WaiverExt           string         `db:"waiver_ext"`
	HasHeraldry         bool           `db:"has_heraldry"`
	HasImage            bool           `db:"has_image"`
	CompanyID           int            `db:"company_id"`
	TokenExpires        string         `db:"token_expires"`
	PasswordExpires     string         `db:"password_expires"`
	PasswordSalt        string         `db:"password_salt"`
	XToken              string         `db:"xtoken"`
	PenaltyBox          int            `db:"penalty_box"`
	Active              bool           `db:"active"`
	Suspended           bool           `db:"suspended"`
	SuspendedByID       sql.NullInt64  `db:"suspended_by_id"`
	SuspendedAt         sql.NullString `db:"suspended_at"`
	SuspendedUntil      sql.NullString `db:"suspended_until"`
	Suspension          sql.NullString `db:"suspension"`
	ReeveQualified      bool           `db:"reeve_qualified"`
	ReeveQualifiedUntil string         `db:"reeve_qualified_until"`
}

func playerList(w http.ResponseWriter, r *http.Request) {
	players := []Player{}

	db := dbInit()

	err := db.Select(&players, "SELECT * FROM ork_mundane ORDER BY persona ASC")
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func playerShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	player := Player{}

	db := dbInit()

	err := db.Get(&player, "SELECT * FROM ork_mundane WHERE mundane_id=?", id)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(player)
}
