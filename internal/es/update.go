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
package es

import (
	"context"
	"fmt"
	"github.com/superhero-match/consumer-superhero-update/internal/es/model"
)

// UpdateSuperhero updates existing Superhero in Elasticsearch.
func (es *ES) UpdateSuperhero(s *model.Superhero) error {
	sourceID, err := es.GetDocumentID(s.ID)
	if err != nil {
		return err
	}

	update, err := es.Client.Update().
		Index(es.Index).
		Id(sourceID).
		Doc(map[string]interface{}{
			"looking_for_gender":       s.LookingForGender,
			"age":                      s.Age,
			"looking_for_age_min":      s.LookingForAgeMin,
			"looking_for_age_max":      s.LookingForAgeMax,
			"looking_for_distance_max": s.LookingForDistanceMax,
			"distance_unit":            s.DistanceUnit,
			"location":                 s.Location,
			"country":                  s.Country,
			"city":                     s.City,
			"superpower":               s.SuperPower,
			"account_type":             s.AccountType,
			"updated_at":               s.UpdatedAt,
		}).
		Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("updated id: ", update.Id)

	return nil
}
