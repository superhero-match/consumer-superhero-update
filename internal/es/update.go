package es

import (
	"context"
	"fmt"
	"github.com/consumer-superhero-update/internal/es/model"
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
