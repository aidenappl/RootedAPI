package structs

type Location struct {
	ID             int     `json:"id"`
	OrganisationID int     `json:"organisation_id"`
	AddressLine1   *string `json:"address_line1"`
	City           *string `json:"city"`
	State          *string `json:"state"`
	ZipCode        *string `json:"zip_code"`
}
