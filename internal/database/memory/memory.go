package memory

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"strings"
	"time"
)

type diagnosesRepository struct {
	diagnoses []models.Diagnose
}

func NewDiagnosesRepository() *diagnosesRepository {
	patients := []models.Patient{
		{
			ID:      uuid.New().String(),
			Name:    "Roberto",
			Surname: "Carlos",
		},
	}

	now := time.Now()

	return &diagnosesRepository{
		diagnoses: []models.Diagnose{
			{
				ID:      uuid.New().String(),
				Patient: patients[0],
				Date:    now,
			},
		},
	}
}

func (r diagnosesRepository) SearchDiagnoses(ctx context.Context, filters models.SearchDiagnosesFilters) ([]models.Diagnose, error) {
	filterByPatientName := len(filters.PatientName) != 0
	filterByDate := !filters.Date.IsZero()

	filteredDiagnoses := lo.Filter(r.diagnoses, func(diagnose models.Diagnose, _ int) bool {
		passPatientNameFilter := !filterByPatientName || doesPatientNameFilterMatch(diagnose.Patient.GetFullName(), filters.PatientName)
		passDateFilter := !filterByDate || doesDateFilterMatch(diagnose.Date, filters.Date)

		return passPatientNameFilter && passDateFilter
	})

	return filteredDiagnoses, nil
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
