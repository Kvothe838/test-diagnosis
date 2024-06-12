package controller

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
)

type diagnosesInteractor interface {
	SearchDiagnoses(context.Context, models.SearchDiagnosesFilters) ([]models.Diagnose, error)
}

type Interactor interface {
	diagnosesInteractor
}
