package router

import (
	"net/http"
	"strconv"

	"github.com/aidenappl/rootedapi/db"
	"github.com/aidenappl/rootedapi/responder"
	"github.com/aidenappl/rootedapi/service"
	"github.com/aidenappl/rootedapi/structs"
	"github.com/gorilla/mux"
)

type HandleOrganisationsRequest struct {
	Requires *[]string `json:"requires"`
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
		Requires: req.Requires,
	})
	if err != nil {
		responder.SendError(w, http.StatusConflict, "Failed to retrieve organisations", err)
		return
	}

	responder.New(w, orgs)
}

type HandleGetOrganisationRequest struct {
	OrgID int `json:"org_id"`
}

func HandleOrganisation(w http.ResponseWriter, r *http.Request) {
	var req HandleGetOrganisationRequest
	params := mux.Vars(r)
	if orgID, ok := params["id"]; ok {
		if orgID == "" {
			responder.SendError(w, http.StatusBadRequest, "Organisation ID is required", nil)
			return
		}
		intID, err := strconv.Atoi(orgID)
		if err != nil {
			responder.SendError(w, http.StatusBadRequest, "Invalid Organisation ID", err)
			return
		}
		req.OrgID = intID
	} else {
		responder.SendError(w, http.StatusBadRequest, "Organisation ID is required", nil)
		return
	}

	orgs, err := service.GetOrganisation(db.DB, req.OrgID)
	if err != nil {
		responder.SendError(w, http.StatusConflict, "Failed to retrieve organisations", err)
		return
	}

	responder.New(w, orgs)
}

// HandleOrganisationPeople handles the request to get people associated with an organisation
type HandleOrganisationPeopleRequest struct {
	OrgID int `json:"org_id"`
}

func HandleOrganisationPeople(w http.ResponseWriter, r *http.Request) {
	var req HandleOrganisationPeopleRequest
	params := mux.Vars(r)
	if orgID, ok := params["id"]; ok {
		if orgID == "" {
			responder.SendError(w, http.StatusBadRequest, "Organisation ID is required", nil)
			return
		}
		intID, err := strconv.Atoi(orgID)
		if err != nil {
			responder.SendError(w, http.StatusBadRequest, "Invalid Organisation ID", err)
			return
		}
		req.OrgID = intID
	} else {
		responder.SendError(w, http.StatusBadRequest, "Organisation ID is required", nil)
		return
	}

	people, err := service.GetOrganisationPeople(db.DB, req.OrgID)
	if err != nil {
		responder.SendError(w, http.StatusConflict, "Failed to retrieve people for organisation", err)
		return
	}

	responder.New(w, people)
}
