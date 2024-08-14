package services

import (
	"context"

	"github.com/Kvothe838/test-diagnosis/internal/models"
	"github.com/Kvothe838/test-diagnosis/internal/pkg/clock"
	"github.com/Kvothe838/test-diagnosis/internal/pkg/uuid"
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

func NewInteractor(repo repository, UUID uuid.UUID, clock clock.Clock) *interactor {
	return &interactor{
		repo:  repo,
		UUID:  UUID,
		clock: clock,
	}
}

type interactor struct {
	repo  repository
	UUID  uuid.UUID
	clock clock.Clock
}
