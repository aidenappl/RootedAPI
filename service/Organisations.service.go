package service

import (
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/aidenappl/rootedapi/db"
	"github.com/aidenappl/rootedapi/structs"
	"github.com/aidenappl/rootedapi/tools"
)

type GetOrganisationsRequest struct {
	WhereID  *int      `json:"where_id"`
	Requires *[]string `json:"requires"`
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
		"l.lat",
		"l.lng",

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
		q = q.OrderBy("o.id DESC")
	}

	if req.Requires != nil {
		for _, require := range *req.Requires {
			// only show organisations with the following requirements
			switch require {
			case "coordinates":
				q = q.Where("l.lat IS NOT NULL AND l.lng IS NOT NULL")
			}
		}
	}

	q = ApplyPagination(q, req.Limit, req.Offset)

	if req.SearchQuery != nil {
		q = q.Where(sq.Like{"name": "%" + *req.SearchQuery + "%"})
	}

	if req.WhereID != nil {
		q = q.Where("o.id = $1", *req.WhereID)
	}

	query, args, err := q.ToSql()
	if err != nil {
		return nil, errors.New("error building SQL query: " + err.Error())
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, errors.New("error executing query: " + err.Error())
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
			&loc.Location.Lat,
			&loc.Location.Lng,

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
		return nil, errors.New("error reading rows: " + err.Error())
	}
	return &organisations, nil
}

func GetOrganisation(db db.Queryable, ID int) (*structs.Organisation, error) {
	orgs, err := GetOrganisations(db, GetOrganisationsRequest{
		BaseListRequest: structs.BaseListRequest{
			Limit: tools.IntP(1),
		},
		WhereID: &ID,
	})
	if err != nil {
		return nil, errors.New("error retrieving organisation: " + err.Error())
	}

	if len(*orgs) == 0 {
		return nil, errors.New("organisation not found")
	}

	return &(*orgs)[0], nil
}

func GetOrganisationPeople(db db.Queryable, orgID int) (*[]structs.Person, error) {
	q := sq.Select(
		"p.id",
		"p.name",
		"p.bookkeeper",
		"p.title",
		"p.phone_number",
		"p.average_hours",
		"p.compensation",

		// People Locations
		"pl.id as person_location_id",
		"pl.address_line_1",
		"pl.city",
		"pl.state",
		"pl.zip_code",
	).
		From("website.people p").
		Where("p.organisation_id = $1", orgID).
		OrderBy("p.id DESC").
		LeftJoin("website.people_locations pl ON p.id = pl.person_id")

	query, args, err := q.ToSql()
	if err != nil {
		return nil, errors.New("error building SQL query: " + err.Error())
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, errors.New("error executing query: " + err.Error())
	}

	defer rows.Close()
	var people []structs.Person
	for rows.Next() {
		var person structs.Person
		var personLocation structs.PeopleLocation
		if err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Bookkeeper,
			&person.Title,
			&person.PhoneNumber,
			&person.AverageHoursPerWeek,
			&person.Compensation,

			&personLocation.Location.ID,
			&personLocation.Location.AddressLine1,
			&personLocation.Location.City,
			&personLocation.Location.State,
			&personLocation.Location.ZipCode,
		); err != nil {
			return nil, err
		}
		if personLocation.Location.ID != nil {
			person.Location = &personLocation
		} else {
			person.Location = nil
		}
		people = append(people, person)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("error reading rows: " + err.Error())
	}
	return &people, nil
}
