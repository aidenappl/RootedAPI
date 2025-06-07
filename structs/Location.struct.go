package structs

type Location struct {
	ID             *int     `json:"location_id"`
	OrganisationID *int     `json:"organisation_id,omitempty"`
	AddressLine1   *string  `json:"address_line1"`
	City           *string  `json:"city"`
	State          *string  `json:"state"`
	ZipCode        *string  `json:"zip_code"`
	Lat            *float64 `json:"lat,omitempty"`
	Lng            *float64 `json:"lng,omitempty"`
}
