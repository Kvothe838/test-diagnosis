package services_tests

import (
	"TopDoctorsBackendChallenge/internal/models"
	"TopDoctorsBackendChallenge/internal/pkg/clock"
	topDoctorsErrors "TopDoctorsBackendChallenge/internal/pkg/errors"
	uuid2 "TopDoctorsBackendChallenge/internal/pkg/uuid"
	"TopDoctorsBackendChallenge/internal/services"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateDiagnosis(t *testing.T) {
	patientID := "patient-id"
	patientName := "Pedro"
	patientSurname := "Picapiedras"
	patientDocumentType := "DNI"
	patientDocumentInfo := "0000000A"
	description := "Test description"
	prescription := "Test prescription"
	now := time.Now()
	diagnosisID := uuid.New().String()

	tests := []struct {
		name                string
		prepareRepositories func(*MockRepository)
		patientID           string
		description         string
		prescription        *string
		assertOnResult      func(diagnosis models.Diagnosis, err error)
	}{
		{
			name: "success",
			prepareRepositories: func(mock *MockRepository) {
				mock.EXPECT().DoesPatientExist(
					gomock.Any(),
					patientID,
				).Return(
					true,
					nil,
				)

				mock.EXPECT().CreateDiagnosis(
					gomock.Any(),
					models.Diagnosis{
						ID:           diagnosisID,
						Patient:      models.Patient{ID: patientID},
						Date:         now,
						Description:  description,
						Prescription: &prescription,
					},
				).Return(
					models.Diagnosis{
						ID: diagnosisID,
						Patient: models.Patient{
							ID:      patientID,
							Name:    patientName,
							Surname: patientSurname,
							Document: models.Document{
								ID:   1,
								Info: patientDocumentInfo,
								Type: models.DocumentType{
									ID:   1,
									Name: patientDocumentType,
								},
							}},
						Date:         now,
						Description:  description,
						Prescription: &prescription,
					},
					nil,
				)
			},
			patientID:    patientID,
			description:  description,
			prescription: &prescription,
			assertOnResult: func(diagnosis models.Diagnosis, err error) {
				assert.NoError(t, err)
				assert.Equal(t, diagnosisID, diagnosis.ID)
				assert.Equal(t, description, diagnosis.Description)
				assert.NotNil(t, diagnosis.Prescription)
				assert.Equal(t, prescription, *diagnosis.Prescription)
				patient := diagnosis.Patient
				assert.Equal(t, patientID, patient.ID)
				assert.Equal(t, patientName, patient.Name)
				assert.Equal(t, patientSurname, patient.Surname)
				assert.Equal(t, 1, patient.Document.ID)
				assert.Equal(t, patientDocumentInfo, patient.Document.Info)
				assert.Equal(t, 1, patient.Document.Type.ID)
				assert.Equal(t, patientDocumentType, patient.Document.Type.Name)
			},
		},
		{
			name: "patient does not exist",
			prepareRepositories: func(mock *MockRepository) {
				mock.EXPECT().DoesPatientExist(
					gomock.Any(),
					patientID,
				).Return(
					false,
					nil,
				)

				mock.EXPECT().CreateDiagnosis(
					gomock.Any(),
					models.Diagnosis{
						ID:           diagnosisID,
						Patient:      models.Patient{ID: patientID},
						Date:         now,
						Description:  description,
						Prescription: &prescription,
					},
				).Times(0)
			},
			patientID:    patientID,
			description:  description,
			prescription: &prescription,
			assertOnResult: func(diagnosis models.Diagnosis, err error) {
				assert.Error(t, err)
				assert.ErrorIs(t, errors.Cause(err), topDoctorsErrors.PatientNotFoundErr)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := NewMockRepository(gomock.NewController(t))

			if test.prepareRepositories != nil {
				test.prepareRepositories(repo)
			}

			in := services.NewInteractor(repo, uuid2.NewFake(diagnosisID), clock.NewFake(now))
			ctx := context.Background()
			diagnosis, err := in.CreateDiagnosis(ctx, test.patientID, test.description, test.prescription)
			if test.assertOnResult != nil {
				test.assertOnResult(diagnosis, err)
			}
		})
	}

}
