package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Response defines the structure for the API JSON response.
type Response struct {
	Message string `json:"message"`
	Random  int    `json:"random"`
}

// homeHandler handles the root route.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

// apiHandler handles the /api route and returns a random number.
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
	response := Response{
		Message: "Here is your random number",
		Random:  rand.Intn(100), // returns a random integer between 0 and 99
	}

	// Set the header to JSON and encode the response.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Create a new router.
	r := mux.NewRouter()

	// Define the routes.
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/api", apiHandler).Methods("GET")

	// Start the server.
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

