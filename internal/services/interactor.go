package services

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
)

type diagnosesRepository interface {
	SearchDiagnoses(context.Context, models.SearchDiagnosesFilters) ([]models.Diagnose, error)
}

func NewInteractor(diagnosesRepo diagnosesRepository) *interactor {
	return &interactor{
		diagnosesRepo: diagnosesRepo,
	}
}

type interactor struct {
	diagnosesRepo diagnosesRepository
}
