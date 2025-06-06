package router

import "net/http"

func HandleOrganisations(w http.ResponseWriter, r *http.Request) {
	// This function will handle the retrieval of all organisations
	// For now, we will just return a placeholder response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List of all organisations"))
}

func HandleOrganisation(w http.ResponseWriter, r *http.Request) {
	// This function will handle the retrieval of all organisations
	// For now, we will just return a placeholder response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List of all organisations"))
}
