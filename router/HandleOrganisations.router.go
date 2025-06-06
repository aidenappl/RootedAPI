package router

import (
	"net/http"

	"github.com/aidenappl/rootedapi/db"
	"github.com/aidenappl/rootedapi/service"
	"github.com/aidenappl/rootedapi/structs"
)

func HandleOrganisations(w http.ResponseWriter, r *http.Request) {
	orgs, err := service.GetOrganisations(db.DB, service.GetOrganisationsRequest{
		BaseListRequest: structs.BaseListRequest{
			Limit:       50,     // Default limit
			Offset:      0,      // Default offset
			SortBy:      "name", // Default sort by name
			SortOrder:   "asc",  // Default sort order
			SearchQuery: "",     // No search query by default
		},
	})
	if err != nil {
		http.Error(w, "Failed to retrieve organisations", http.StatusInternalServerError)
		return
	}
	
	
}

func HandleOrganisation(w http.ResponseWriter, r *http.Request) {
	// This function will handle the retrieval of all organisations
	// For now, we will just return a placeholder response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List of all organisations"))
}
