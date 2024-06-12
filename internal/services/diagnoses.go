package services

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
	"github.com/pkg/errors"
)

func (in *interactor) SearchDiagnoses(ctx context.Context, filters models.SearchDiagnosesFilters) ([]models.Diagnose, error) {
	diagnoses, err := in.diagnosesRepo.SearchDiagnoses(ctx, filters)
	if err != nil {
		return nil, errors.Wrap(err, "could not search for diagnoses")
	}

	return diagnoses, nil
}
