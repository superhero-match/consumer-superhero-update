/*
  Copyright (C) 2019 - 2022 MWSOFT
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
package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	elastic "github.com/olivere/elastic/v7"
	"go.uber.org/zap"

	"github.com/superhero-match/consumer-superhero-update/cmd/consumer/model"
	dbm "github.com/superhero-match/consumer-superhero-update/internal/db/model"
	esm "github.com/superhero-match/consumer-superhero-update/internal/es/model"
)

// Read consumes the Kafka topic and updates superhero in DB and Elasticsearch.
func (r *reader) Read() error {
	ctx := context.Background()

	for {
		fmt.Print("before FetchMessage")
		m, err := r.Consumer.FetchMessage(ctx)
		fmt.Print("after FetchMessage")
		if err != nil {
			r.Logger.Error(
				"failed to fetch message",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		fmt.Printf(
			"message at topic/partition/offset \n%v/\n%v/\n%v: \n%s = \n%s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)

		var s model.Superhero
		if err := json.Unmarshal(m.Value, &s); err != nil {
			_ = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to unmarshal JSON to superhero",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				err = r.Consumer.Close()
				if err != nil {
					r.Logger.Error(
						"failed to close consumer",
						zap.String("err", err.Error()),
						zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
					)

					return err
				}

				return err
			}
		}

		err = r.DB.UpdateSuperhero(dbm.Superhero{
			ID:                    s.ID,
			LookingForGender:      s.LookingForGender,
			Age:                   s.Age,
			LookingForAgeMin:      s.LookingForAgeMin,
			LookingForAgeMax:      s.LookingForAgeMax,
			LookingForDistanceMax: s.LookingForDistanceMax,
			DistanceUnit:          s.DistanceUnit,
			Lat:                   s.Lat,
			Lon:                   s.Lon,
			Country:               s.Country,
			City:                  s.City,
			SuperPower:            s.SuperPower,
			AccountType:           s.AccountType,
			UpdatedAt:             s.UpdatedAt,
		})
		if err != nil {
			r.Logger.Error(
				"failed to update superhero in database",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		err = r.ES.UpdateSuperhero(&esm.Superhero{
			ID:                    s.ID,
			LookingForGender:      s.LookingForGender,
			Age:                   s.Age,
			LookingForAgeMin:      s.LookingForAgeMin,
			LookingForAgeMax:      s.LookingForAgeMax,
			LookingForDistanceMax: s.LookingForDistanceMax,
			DistanceUnit:          s.DistanceUnit,
			Location: elastic.GeoPoint{
				Lat: s.Lat,
				Lon: s.Lon,
			},
			Country:     s.Country,
			City:        s.City,
			SuperPower:  s.SuperPower,
			AccountType: s.AccountType,
			UpdatedAt:   s.UpdatedAt,
		})
		if err != nil {
			r.Logger.Error(
				"failed to update superhero in Elasticsearch",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}

		err = r.Consumer.CommitMessages(ctx, m)
		if err != nil {
			r.Logger.Error(
				"failed to commit message",
				zap.String("err", err.Error()),
				zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
			)

			err = r.Consumer.Close()
			if err != nil {
				r.Logger.Error(
					"failed to close consumer",
					zap.String("err", err.Error()),
					zap.String("time", time.Now().UTC().Format(r.TimeFormat)),
				)

				return err
			}

			return err
		}
	}
}
