package services_tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Kvothe838/test-diagnosis/internal/models"
	"github.com/Kvothe838/test-diagnosis/internal/pkg/clock"
	internalErrors "github.com/Kvothe838/test-diagnosis/internal/pkg/errors"
	uuid2 "github.com/Kvothe838/test-diagnosis/internal/pkg/uuid"
	"github.com/Kvothe838/test-diagnosis/internal/services"

	"github.com/golang/mock/gomock"
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
				assert.True(t, errors.Is(err, internalErrors.PatientNotFoundErr))
			},
		},
		{
			name: "error on checking if patient exists",
			prepareRepositories: func(mock *MockRepository) {
				mock.EXPECT().DoesPatientExist(
					gomock.Any(),
					patientID,
				).Return(
					false,
					errors.New("error"),
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
				assert.ErrorContains(t, err, "could not check if patient exists")
			},
		},
		{
			name: "error on creating diagnosis",
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
					models.Diagnosis{},
					errors.New("error"),
				)
			},
			patientID:    patientID,
			description:  description,
			prescription: &prescription,
			assertOnResult: func(diagnosis models.Diagnosis, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not create diagnosis")
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

func TestSearchDiagnoses(t *testing.T) {
	patientName := "pedro"
	date := time.Now()
	now := time.Now()
	diagnosisID := uuid.New().String()
	description := "Test description"
	prescription := "Test prescription"

	tests := []struct {
		name                string
		prepareRepositories func(*MockRepository)
		filters             models.SearchDiagnosesFilters
		assertOnResult      func(diagnoses []models.Diagnosis, err error)
	}{
		{
			name: "success with empty response",
			prepareRepositories: func(mock *MockRepository) {
				mock.EXPECT().SearchDiagnoses(
					gomock.Any(),
					models.SearchDiagnosesFilters{
						PatientName: patientName,
						Date:        date,
					},
				).Return(
					nil,
					nil,
				)
			},
			filters: models.SearchDiagnosesFilters{
				PatientName: patientName,
				Date:        date,
			},
			assertOnResult: func(diagnoses []models.Diagnosis, err error) {
				assert.NoError(t, err)
				assert.Empty(t, diagnoses)
			},
		},
		{
			name: "success with responses",
			prepareRepositories: func(mock *MockRepository) {
				mock.EXPECT().SearchDiagnoses(
					gomock.Any(),
					models.SearchDiagnosesFilters{
						PatientName: patientName,
						Date:        date,
					},
				).Return(
					[]models.Diagnosis{
						{
							ID:           diagnosisID,
							Patient:      models.Patient{},
							Date:         date,
							Description:  description,
							Prescription: &prescription,
						},
					},
					nil,
				)
			},
			filters: models.SearchDiagnosesFilters{
				PatientName: patientName,
				Date:        date,
			},
			assertOnResult: func(diagnoses []models.Diagnosis, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, diagnoses)
				firstDiagnosis := diagnoses[0]
				assert.Equal(t, diagnosisID, firstDiagnosis.ID)
				assert.Equal(t, date, firstDiagnosis.Date)
				assert.Equal(t, description, firstDiagnosis.Description)
				assert.NotNil(t, firstDiagnosis.Prescription)
				assert.Equal(t, prescription, *firstDiagnosis.Prescription)
			},
		},
		{
			name: "error when search diagnoses",
			prepareRepositories: func(mock *MockRepository) {
				mock.EXPECT().SearchDiagnoses(
					gomock.Any(),
					models.SearchDiagnosesFilters{
						PatientName: patientName,
						Date:        date,
					},
				).Return(
					nil,
					errors.New("error"),
				)
			},
			filters: models.SearchDiagnosesFilters{
				PatientName: patientName,
				Date:        date,
			},
			assertOnResult: func(diagnoses []models.Diagnosis, err error) {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "could not search for diagnoses")
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
			diagnoses, err := in.SearchDiagnoses(ctx, test.filters)
			if test.assertOnResult != nil {
				test.assertOnResult(diagnoses, err)
			}
		})
	}

}
