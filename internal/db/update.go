package db

import "github.com/consumer-superhero-update/internal/db/model"

// UpdateSuperhero saves newly registered Superhero.
func(db *DB) UpdateSuperhero(s model.Superhero) error {
	_, err := db.stmtUpdateSuperhero.Exec(
		s.ID,
		s.MainProfilePicURL,
		s.LookingForGender,
		s.Age,
		s.LookingForAgeMin,
		s.LookingForAgeMax,
		s.LookingForDistanceMax,
		s.DistanceUnit,
		s.Lat,
		s.Lon,
		s.Country,
		s.City,
		s.SuperPower,
		s.AccountType,
		s.IsDeleted,
		s.DeletedAt,
		s.IsBlocked,
		s.BlockedAt,
		s.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
