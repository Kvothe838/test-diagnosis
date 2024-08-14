package controller

import (
	"time"

	"github.com/Kvothe838/test-diagnosis/internal/models"
	"github.com/samber/lo"
)

type diagnosisDTO struct {
	Patient      patientDTO `json:"patient"`
	Date         string     `json:"date,omitempty"`
	Description  string     `json:"description"`
	Prescription *string    `json:"prescription,omitempty"`
}

type patientDTO struct {
	Name     string       `json:"name"`
	Surname  string       `json:"surname"`
	Document documentDTO  `json:"document"`
	Contacts []contactDTO `json:"contacts"`
}

type documentDTO struct {
	Info string          `json:"info"`
	Type documentTypeDTO `json:"type"`
}

type documentTypeDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type contactDTO struct {
	Type contactTypeDTO `json:"type"`
	Info string         `json:"info"`
}

type contactTypeDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func mapToDiagnosesDTO(diagnoses []models.Diagnosis) []diagnosisDTO {
	return lo.Map(diagnoses, func(diagnosis models.Diagnosis, _ int) diagnosisDTO {
		return mapToDiagnosisDTO(diagnosis)
	})
}

func mapToDiagnosisDTO(diagnosis models.Diagnosis) diagnosisDTO {
	return diagnosisDTO{
		Patient:      mapToPatientDTO(diagnosis.Patient),
		Date:         mapToDateDTO(diagnosis.Date),
		Description:  diagnosis.Description,
		Prescription: diagnosis.Prescription,
	}
}

func mapToDateDTO(date time.Time) string {
	return date.Format(time.Layout)
}

func mapToPatientDTO(patient models.Patient) patientDTO {
	return patientDTO{
		Name:     patient.Name,
		Surname:  patient.Surname,
		Document: mapToDocumentDTO(patient.Document),
		Contacts: mapToContactsDTO(patient.Contacts),
	}
}

func mapToDocumentDTO(document models.Document) documentDTO {
	return documentDTO{
		Info: document.Info,
		Type: mapToDocumentTypeDTO(document.Type),
	}
}

func mapToDocumentTypeDTO(documentType models.DocumentType) documentTypeDTO {
	return documentTypeDTO{
		ID:   documentType.ID,
		Name: documentType.Name,
	}
}

func mapToContactsDTO(contacts []models.Contact) []contactDTO {
	return lo.Map(contacts, func(contact models.Contact, _ int) contactDTO {
		return mapToContactDTO(contact)
	})
}

func mapToContactDTO(contact models.Contact) contactDTO {
	return contactDTO{
		Type: mapToContactTypeDTO(contact.Type),
		Info: contact.Info,
	}
}

func mapToContactTypeDTO(contactType models.ContactType) contactTypeDTO {
	return contactTypeDTO{
		ID:   contactType.ID,
		Name: contactType.Name,
	}
}
