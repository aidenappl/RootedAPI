package service

import (
	"github.com/aidenappl/rootedapi/db"
	"github.com/aidenappl/rootedapi/structs"
)

type GetOrganisationsRequest struct {
	structs.BaseListRequest
}

func GetOrganisations(db db.Queryable, req GetOrganisationsRequest) (*[]structs.Organisation, error) {
	// This function will handle the retrieval of all organisations
	// For now, we will just return a placeholder response
	// In a real implementation, you would query the database here

	// Placeholder response
	organisations := []structs.Organisation{
		{
			ID:           1,
			Name:         "Rooted Nonprofit",
			EIN:          "12-3456789",
			DIN:          nil,
			XML_Batch_ID: nil,
			Website:      nil,
			Description:  nil,
		},
	}
	return &organisations, nil
}
