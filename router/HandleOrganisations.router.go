package router

import (
	"net/http"

	"github.com/aidenappl/rootedapi/db"
	"github.com/aidenappl/rootedapi/responder"
	"github.com/aidenappl/rootedapi/service"
	"github.com/aidenappl/rootedapi/structs"
)

type HandleOrganisationsRequest struct {
	structs.BaseListRequest
}

func HandleOrganisations(w http.ResponseWriter, r *http.Request) {

	var req HandleOrganisationsRequest
	if err := ParseURLParams(r, &req); err != nil {
		responder.SendError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	orgs, err := service.GetOrganisations(db.DB, service.GetOrganisationsRequest{
		BaseListRequest: structs.BaseListRequest{
			Limit:     req.Limit,
			Offset:    req.Offset,
			SortOrder: req.SortOrder,
		},
	})
	if err != nil {
		responder.SendError(w, http.StatusConflict, "Failed to retrieve organisations", err)
		return
	}

	responder.New(w, orgs)
}

func HandleOrganisation(w http.ResponseWriter, r *http.Request) {
	// This function will handle the retrieval of all organisations
	// For now, we will just return a placeholder response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List of all organisations"))
}
