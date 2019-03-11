package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/healthcheck", healthCheck)

	router.Get("/kingdoms", kingdomList)
	router.Get("/kingdom/{kingdomID:[0-9]+}", kingdomShow)
	router.Get("/kingdom/{kingdomID:[0-9]+}/events", kingdomEventsShow)
	router.Get("/kingdom/{kingdomID:[0-9]+}/officers", kingdomOfficersShow)

	router.Get("/players", playerList)
	router.Get("/player/{playerID:[0-9]+}", playerShow)
	router.Get("/player/{playerID:[0-9]+}/classes", playerClassesShow)

	router.Get("/parks", parkList)
	router.Get("/park/{parkID:[0-9]+}", parkShow)

	fmt.Println("OGRE is online!")
	log.Fatal(http.ListenAndServe(":3736", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("We're alive!")
}
