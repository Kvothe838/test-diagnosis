package services

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
)

type diagnosesRepository interface {
	CreateDiagnosis(context.Context, models.Diagnosis) (models.Diagnosis, error)
	SearchDiagnoses(context.Context, models.SearchDiagnosesFilters) ([]models.Diagnosis, error)
}

type patientsRepository interface {
	DoesPatientExist(ctx context.Context, patientID string) (bool, error)
}

type repository interface {
	diagnosesRepository
	patientsRepository
}

func NewInteractor(repo repository) *interactor {
	return &interactor{
		repo: repo,
	}
}

type interactor struct {
	repo repository
}
