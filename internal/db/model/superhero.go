/*
  Copyright (C) 2019 - 2020 MWSOFT
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
