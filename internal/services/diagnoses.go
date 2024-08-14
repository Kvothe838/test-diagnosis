package services

import (
	"context"
	"fmt"

	"github.com/Kvothe838/test-diagnosis/internal/models"
	"github.com/Kvothe838/test-diagnosis/internal/pkg/errors"
)

func (in *interactor) CreateDiagnosis(ctx context.Context, patientID, description string, prescription *string) (models.Diagnosis, error) {
	patientExists, err := in.repo.DoesPatientExist(ctx, patientID)
	if err != nil {
		return models.Diagnosis{}, fmt.Errorf("could not check if patient exists: %w", err)
	}

	if !patientExists {
		return models.Diagnosis{}, errors.PatientNotFoundErr
	}

	now := in.clock.Now()
	diagnosisID := in.UUID.GetNew()

	diagnosis := models.Diagnosis{
		ID: diagnosisID,
		Patient: models.Patient{
			ID: patientID,
		},
		Date:         now,
		Description:  description,
		Prescription: prescription,
	}

	diagnosis, err = in.repo.CreateDiagnosis(ctx, diagnosis)
	if err != nil {
		return models.Diagnosis{}, fmt.Errorf("could not create diagnosis: %w", err)
	}

	return diagnosis, nil
}

func (in *interactor) SearchDiagnoses(ctx context.Context, filters models.SearchDiagnosesFilters) ([]models.Diagnosis, error) {
	diagnoses, err := in.repo.SearchDiagnoses(ctx, filters)
	if err != nil {
		return nil, fmt.Errorf("could not search for diagnoses: %w", err)
	}

	return diagnoses, nil
}
