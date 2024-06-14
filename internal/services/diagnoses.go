package services

import (
	"TopDoctorsBackendChallenge/internal/models"
	topDoctorsErrors "TopDoctorsBackendChallenge/internal/pkg/errors"
	"context"
	"github.com/pkg/errors"
)

func (in *interactor) CreateDiagnosis(ctx context.Context, patientID, description string, prescription *string) (models.Diagnosis, error) {
	patientExists, err := in.repo.DoesPatientExist(ctx, patientID)
	if err != nil {
		return models.Diagnosis{}, errors.Wrap(err, "could not check if patient exists")
	}

	if !patientExists {
		return models.Diagnosis{}, topDoctorsErrors.PatientNotFoundErr
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
		return models.Diagnosis{}, errors.Wrap(err, "could not create diagnosis")
	}

	return diagnosis, nil
}

func (in *interactor) SearchDiagnoses(ctx context.Context, filters models.SearchDiagnosesFilters) ([]models.Diagnosis, error) {
	diagnoses, err := in.repo.SearchDiagnoses(ctx, filters)
	if err != nil {
		return nil, errors.Wrap(err, "could not search for diagnoses")
	}

	return diagnoses, nil
}
