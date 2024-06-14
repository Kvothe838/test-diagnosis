package controller

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
)

type diagnosesInteractor interface {
	CreateDiagnosis(ctx context.Context, patientID, description string, prescription *string) (models.Diagnosis, error)
	SearchDiagnoses(context.Context, models.SearchDiagnosesFilters) ([]models.Diagnosis, error)
}

type Interactor interface {
	diagnosesInteractor
}
