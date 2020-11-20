/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package model

import elastic "github.com/olivere/elastic/v7"

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
