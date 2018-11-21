package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()

	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/kingdoms", kingdomList).Methods("GET")
	router.HandleFunc("/kingdom/{id:[0-9]+}", kingdomShow).Methods("GET")
	router.HandleFunc("/players", playerList).Methods("GET")
	router.HandleFunc("/player/{id:[0-9]+}", playerShow).Methods("GET")

	fmt.Println("OGRE is online!")
	log.Fatal(http.ListenAndServe(":3736", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("We're alive!")
}
