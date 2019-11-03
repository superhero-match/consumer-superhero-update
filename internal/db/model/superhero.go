package model

type Superhero struct {
	ID                    string  `db:"id"`
	LookingForGender      int     `db:"looking_for_gender"`
	Age                   int     `db:"age"`
	LookingForAgeMin      int     `db:"looking_for_age_min"`
	LookingForAgeMax      int     `db:"looking_for_age_max"`
	LookingForDistanceMax int     `db:"looking_for_distance_max"`
	DistanceUnit          string  `db:"distance_unit"`
	Lat                   float64 `db:"lat"`
	Lon                   float64 `db:"lon"`
	Country               string  `db:"country"`
	City                  string  `db:"city"`
	SuperPower            string  `db:"superpower"`
	AccountType           string  `db:"account_type"`
	UpdatedAt             string  `db:"updated_at"`
}
