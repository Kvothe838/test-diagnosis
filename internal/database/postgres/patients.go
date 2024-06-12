package postgres

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
)

func (r repository) getPatientByID(ctx context.Context, ID string) (models.Patient, error) {
	return models.Patient{
		ID: ID,
	}, nil
}
