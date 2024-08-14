package memory

import (
	"context"
	"strings"
	"time"

	"github.com/Kvothe838/test-diagnosis/internal/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type repository struct {
	diagnoses []models.Diagnosis
	patients  []models.Patient
}

func NewRepository() *repository {
	patients := []models.Patient{
		{
			ID:      uuid.New().String(),
			Name:    "Roberto",
			Surname: "Carlos",
		},
	}

	now := time.Now()

	return &repository{
		diagnoses: []models.Diagnosis{
			{
				ID:      uuid.New().String(),
				Patient: patients[0],
				Date:    now,
			},
		},
		patients: patients,
	}
}

func (r repository) SearchDiagnoses(ctx context.Context, filters models.SearchDiagnosesFilters) ([]models.Diagnosis, error) {
	filterByPatientName := len(filters.PatientName) != 0
	filterByDate := !filters.Date.IsZero()

	filteredDiagnoses := lo.Filter(r.diagnoses, func(diagnose models.Diagnosis, _ int) bool {
		passPatientNameFilter := !filterByPatientName || doesPatientNameFilterMatch(diagnose.Patient.GetFullName(), filters.PatientName)
		passDateFilter := !filterByDate || doesDateFilterMatch(diagnose.Date, filters.Date)

		return passPatientNameFilter && passDateFilter
	})

	return filteredDiagnoses, nil
}

func (r *repository) CreateDiagnosis(ctx context.Context, diagnosis models.Diagnosis) (models.Diagnosis, error) {
	r.diagnoses = append(r.diagnoses, diagnosis)
	for _, patient := range r.patients {
		if patient.ID == diagnosis.Patient.ID {
			diagnosis.Patient = patient
		}
	}

	return diagnosis, nil
}

func doesDateFilterMatch(diagnoseDate, filtersDate time.Time) bool {
	return getStartOfDay(diagnoseDate).Equal(getStartOfDay(filtersDate))
}

func getStartOfDay(date time.Time) time.Time {
	year, month, day := date.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, date.Location())

	return startOfDay
}

func doesPatientNameFilterMatch(patientFullName, filterPatientName string) bool {
	return strings.Contains(strings.ToLower(patientFullName), strings.ToLower(filterPatientName))
}
