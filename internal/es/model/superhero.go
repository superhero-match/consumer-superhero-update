package model

import "gopkg.in/olivere/elastic.v7"

type Superhero struct {
	ID                    string           `json:"superhero_id"`
	LookingForGender      int              `json:"looking_for_gender"`
	Age                   int              `json:"age"`
	LookingForAgeMin      int              `json:"looking_for_age_min"`
	LookingForAgeMax      int              `json:"looking_for_age_max"`
	LookingForDistanceMax int              `json:"looking_for_distance_max"`
	DistanceUnit          string           `json:"distance_unit"`
	Location              elastic.GeoPoint `json:"location"`
	Country               string           `json:"country"`
	City                  string           `json:"city"`
	SuperPower            string           `json:"superpower"`
	AccountType           string           `json:"account_type"`
	UpdatedAt             string           `json:"updated_at"`
}
