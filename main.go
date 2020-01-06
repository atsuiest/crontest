package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Run a Job
func runFrequent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := `{"mensaje":"Se ha ejecutado la tarea correctamente"}`
	json.NewEncoder(w).Encode(response)
}

func main() {
	// current_time := time.Now()
	// fmt.Println("Current time in String: ", current_time.Hour(), ":", current_time.Minute(), ":", current_time.Second())
	// fmt.Println(current_time.Clock())

	// Init router
	r := mux.NewRouter()

	// Handle endpoints
	r.HandleFunc("/run/frequent", runFrequent).Methods("POST")

	// Run server
	log.Fatal(http.ListenAndServe(":8000", r))
}
