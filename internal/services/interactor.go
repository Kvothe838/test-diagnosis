package services

import (
	"TopDoctorsBackendChallenge/internal/models"
	"TopDoctorsBackendChallenge/internal/pkg/clock"
	uuid2 "TopDoctorsBackendChallenge/internal/pkg/uuid"
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

func NewInteractor(repo repository, UUID uuid2.UUID, clock clock.Clock) *interactor {
	return &interactor{
		repo:  repo,
		UUID:  UUID,
		clock: clock,
	}
}

type interactor struct {
	repo  repository
	UUID  uuid2.UUID
	clock clock.Clock
}
