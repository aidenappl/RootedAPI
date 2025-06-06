package service

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/aidenappl/rootedapi/db"
	"github.com/aidenappl/rootedapi/structs"
	"github.com/aidenappl/rootedapi/tools"
)

type GetOrganisationsRequest struct {
	structs.BaseListRequest
}

func GetOrganisations(db db.Queryable, req GetOrganisationsRequest) (*[]structs.Organisation, error) {
	q := sq.Select(
		"o.id",
		"o.name",
		"o.ein",
		"o.dln",
		"o.xml_batch_id",
		"o.website",
		"o.description",

		// Location fields
		"l.id as location_id",
		"l.address_line_1",
		"l.city",
		"l.state",
		"l.zip_code",

		// Metadata fields
		"m.id as metadata_id",
		"m.gross_reciepts_amt",
		"m.total_revenue_amt",
		"m.total_expenses_amt",
		"m.excess_or_deficit_for_year_amt",
	).
		From("website.organisations o").
		LeftJoin("website.organisation_locations l ON o.id = l.organisation_id").
		LeftJoin("website.organisation_metadata m ON o.id = m.organisation_id")

	if req.SortBy != nil {
		q = q.OrderBy(*req.SortBy + " " + *req.SortOrder)
	} else {
		q = q.OrderBy("id desc")
	}

	q = ApplyPagination(q, req.Limit, req.Offset)

	if req.SearchQuery != nil {
		q = q.Where(sq.Like{"name": "%" + *req.SearchQuery + "%"})
	}

	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var organisations []structs.Organisation
	for rows.Next() {
		var org structs.Organisation
		var loc structs.OrganisationLocation
		var metadata structs.OrganisationMetadata
		if err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.EIN,
			&org.DLN,
			&org.XMLBatchID,
			&org.Website,
			&org.Description,

			&loc.Location.ID,
			&loc.Location.AddressLine1,
			&loc.Location.City,
			&loc.Location.State,
			&loc.Location.ZipCode,

			&metadata.ID,
			&metadata.GrossReceipts,
			&metadata.TotalRevenue,
			&metadata.TotalExpenses,
			&metadata.ExcessOrDeficit,
		); err != nil {
			return nil, err
		}
		org.Location = &loc.Location
		if metadata.ID != nil {
			org.Metadata = &metadata
		}
		organisations = append(organisations, org)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &organisations, nil
}

type GetOrganisationRequest struct {
	ID int `json:"id"`
}

func GetOrganisation(db db.Queryable, req GetOrganisationRequest) (*structs.Organisation, error) {
	orgs, err := GetOrganisations(db, GetOrganisationsRequest{
		BaseListRequest: structs.BaseListRequest{
			Limit:       tools.IntP(1),
			Offset:      tools.IntP(0),
			SortBy:      tools.StringP("id"),
			SortOrder:   tools.StringP("asc"),
			SearchQuery: tools.StringP("id:" + string(req.ID)),
		},
	})
	if err != nil {
		return nil, err
	}
	return &(*orgs)[0], nil
}
