package structs

type Organisation struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	EIN         string                `json:"ein"`
	DLN         *string               `json:"dln"`
	XMLBatchID  *string               `json:"xml_batch_id"`
	Website     *string               `json:"website"`
	Description *string               `json:"description"`
	Location    *Location             `json:"location"`
	People      *[]Person             `json:"people,omitempty"`
	Metadata    *OrganisationMetadata `json:"metadata"`
}

type OrganisationMetadata struct {
	ID              *int     `json:"id"`
	OrganisationID  *int     `json:"organisation_id,omitempty"`
	GrossReceipts   *float64 `json:"gross_receipts"`
	TotalRevenue    *float64 `json:"total_revenue"`
	TotalExpenses   *float64 `json:"total_expenses"`
	ExcessOrDeficit *float64 `json:"excess_or_deficit"`
}

type OrganisationLocation struct {
	OrganisationID int `json:"organisation_id"`
	Location
}
