package postgres

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
	"github.com/pkg/errors"
)

func (r repository) getPatientByID(ctx context.Context, ID string) (models.Patient, error) {
	query := `
		SELECT id, name, surname, document_id
		FROM patients
		WHERE id = ?
	`

	row := r.conn.QueryRowContext(ctx, query, ID)
	var patient models.Patient
	err := row.Scan(&patient.ID, &patient.Name, &patient.Surname, &patient.Document.ID)
	if err != nil {
		return models.Patient{}, errors.Wrap(err, "could not scan patient")
	}

	patient.Document, err = r.getPatientDocumentByID(ctx, patient.Document.ID)
	if err != nil {
		return models.Patient{}, errors.Wrap(err, "could not get patient document")
	}

	patient.Contacts, err = r.getPatientContacts(ctx, patient.ID)
	if err != nil {
		return models.Patient{}, errors.Wrap(err, "could not get patient contacts")
	}

	return patient, nil
}

func (r repository) getPatientDocumentByID(ctx context.Context, ID int) (models.Document, error) {
	query := `
		SELECT d.info, dt.id, dt.name
		FROM documents d
		JOIN document_types dt ON d.type_id = dt.id
		WHERE d.id = ?
	`

	row := r.conn.QueryRowContext(ctx, query, ID)
	var document models.Document
	err := row.Scan(&document.Info, &document.Type.ID, &document.Type.Name)
	if err != nil {
		return models.Document{}, errors.Wrap(err, "could not scan patient")
	}

	return document, nil
}

func (r *repository) getPatientContacts(ctx context.Context, patientID string) ([]models.Contact, error) {
	query := `
		SELECT c.info, ct.id, ct.name
		FROM contacts c
		JOIN contact_types ct ON c.type_id = ct.id
		WHERE c.patient_id = ?
	`

	rows, err := r.conn.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, errors.Wrapf(err, "could not execute query to get patient contacts for patient ID %s", patientID)
	}

	contacts := make([]models.Contact, 0)

	for rows.Next() {
		var contact models.Contact

		err = rows.Scan(&contact.Info, contact.Type.ID, contact.Type.Name)
		if err != nil {
			return nil, errors.Wrap(err, "could not scan contact")
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}
