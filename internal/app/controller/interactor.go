package controller

import (
	"context"

	"github.com/Kvothe838/test-diagnosis/internal/models"
)

type diagnosesInteractor interface {
	CreateDiagnosis(ctx context.Context, patientID, description string, prescription *string) (models.Diagnosis, error)
	SearchDiagnoses(context.Context, models.SearchDiagnosesFilters) ([]models.Diagnosis, error)
}

type Interactor interface {
	diagnosesInteractor
}
