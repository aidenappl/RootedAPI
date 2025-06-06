package structs

type Person struct {
	ID                  int      `json:"id"`
	OrganisationID      int      `json:"organisation_id"`
	Bookkeeper          bool     `json:"bookkeeper"`
	Name                *string  `json:"name"`
	Title               *string  `json:"title"`
	PhoneNumber         *string  `json:"phone_number"`
	AverageHoursPerWeek *int     `json:"average_hours_per_week"`
	Compensation        *float64 `json:"compensation"`
}

type PeopleLocation struct {
	PersonID int `json:"person_id"`
	Location
}
