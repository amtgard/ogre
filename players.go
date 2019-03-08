package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// PlayerClass is Amtgard class information for a player
type PlayerClass struct {
	ClassID               int             `db:"class_id"`
	Active                bool            `db:"active"`
	ClassName             string          `db:"class_name"`
	Weeks                 int             `db:"weeks"`
	Attendances           sql.NullInt64   `db:"attendances"`
	Credits               sql.NullFloat64 `db:"credits"`
	ClassReconciliationID sql.NullInt64   `db:"class_reconciliation_id"`
	Reconciled            sql.NullBool    `db:"reconciled"`
}

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

func playerClassesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	classes := []PlayerClass{}

	sql := `
	SELECT
	c.class_id,
	c.active,
	c.name as class_name,
	COUNT(a.week) as weeks,
	SUM(a.attendances) as attendances,
	SUM(a.credits) as credits,
	cr.class_reconciliation_id,
	cr.reconciled
	FROM ork_class c
	LEFT JOIN
	(
	SELECT
	ssa.class_id,
	COUNT(ssa.attendance_id) as attendances,
	SUM(ssa.credits) as credits,
	week(ssa.date, 6) as week
		FROM
		(
		SELECT
		min(killdupe.attendance_id) as attendance_id
		FROM ork_attendance killdupe
		WHERE killdupe.mundane_id = ? GROUP BY killdupe.date
	) kd
		LEFT JOIN ork_attendance ssa ON ssa.attendance_id = kd.attendance_id
		WHERE
		ssa.mundane_id = ?
		GROUP BY ssa.class_id, ssa.date
	) a ON a.class_id = c.class_id
	LEFT JOIN ork_class_reconciliation cr ON cr.class_id = c.class_id AND cr.mundane_id = ?
	WHERE c.active = 1
	GROUP BY c.class_id, c.active, c.name, cr.class_reconciliation_id, cr.reconciled
`

	db := dbInit()

	err := db.Select(&classes, sql, id, id, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(classes)
}
