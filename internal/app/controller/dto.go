package controller

import (
	"TopDoctorsBackendChallenge/internal/models"
	"github.com/samber/lo"
	"time"
)

type patientDTO struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type diagnoseDTO struct {
	Patient      patientDTO `json:"patient"`
	Date         string     `json:"date,omitempty"`
	Description  string     `json:"description"`
	Prescription *string    `json:"prescription,omitempty""`
}

func mapToDiagnosesDTO(diagnoses []models.Diagnose) []diagnoseDTO {
	return lo.Map(diagnoses, func(diagnose models.Diagnose, _ int) diagnoseDTO {
		return diagnoseDTO{
			Patient:      mapToPatientDTO(diagnose.Patient),
			Date:         mapToDateDTO(diagnose.Date),
			Description:  diagnose.Description,
			Prescription: diagnose.Prescription,
		}
	})
}

func mapToPatientDTO(patient models.Patient) patientDTO {
	return patientDTO{
		Name:    patient.Name,
		Surname: patient.Surname,
	}
}

func mapToDateDTO(date time.Time) string {
	return date.Format(time.DateTime)
}
