package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aidenappl/rootedapi/env"
	"github.com/aidenappl/rootedapi/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// Public Endpoints
	// [Organisations]
	// - This endpoint retrieves a list of all organisations
	r.HandleFunc("/organisations", router.HandleOrganisations).Methods("GET")
	// - This endpoint retrieves a specific organisation by ID
	r.HandleFunc("/organisations/{id}", router.HandleOrganisation).Methods("GET")
	// - This endpoint retrieves the people associated with a specific organisation by ID
	r.HandleFunc("/organisations/{id}/people", router.HandleOrganisationPeople).Methods("GET")

	// Health Check Endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// CORS Middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	fmt.Printf("✅ Rooted API running on port %s\n", env.Port)
	log.Fatal(http.ListenAndServe(":"+env.Port, corsMiddleware.Handler(r)))
}
