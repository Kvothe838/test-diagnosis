package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Kvothe838/test-diagnosis/internal/models"
)

func (r *repository) DoesPatientExist(ctx context.Context, patientID string) (bool, error) {
	query := `
		SELECT 1
		FROM patients
		WHERE id = $1
	`

	var patientExists int
	err := r.conn.QueryRowContext(ctx, query, patientID).Scan(&patientExists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, fmt.Errorf("could not check if patient exists: %w", err)
	}

	return true, nil
}

func (r *repository) getPatientByID(ctx context.Context, ID string) (models.Patient, error) {
	query := `
		SELECT id, name, surname, document_id
		FROM patients
		WHERE id = $1
	`

	row := r.conn.QueryRowContext(ctx, query, ID)
	var patient models.Patient
	err := row.Scan(&patient.ID, &patient.Name, &patient.Surname, &patient.Document.ID)
	if err != nil {
		return models.Patient{}, fmt.Errorf("could not scan patient: %w", err)
	}

	patient.Document, err = r.getPatientDocumentByID(ctx, patient.Document.ID)
	if err != nil {
		return models.Patient{}, fmt.Errorf("could not get patient document: %w", err)
	}

	patient.Contacts, err = r.getPatientContacts(ctx, patient.ID)
	if err != nil {
		return models.Patient{}, fmt.Errorf("could not get patient contacts: %w", err)
	}

	return patient, nil
}

func (r *repository) getPatientDocumentByID(ctx context.Context, ID int) (models.Document, error) {
	query := `
		SELECT d.info, dt.id, dt.name
		FROM documents d
		JOIN document_types dt ON d.type_id = dt.id
		WHERE d.id = $1
	`

	row := r.conn.QueryRowContext(ctx, query, ID)
	var document models.Document
	err := row.Scan(&document.Info, &document.Type.ID, &document.Type.Name)
	if err != nil {
		return models.Document{}, fmt.Errorf("could not scan patient: %w", err)
	}

	return document, nil
}

func (r *repository) getPatientContacts(ctx context.Context, patientID string) ([]models.Contact, error) {
	query := `
		SELECT c.info, ct.id, ct.name
		FROM contacts c
		JOIN contact_types ct ON c.type_id = ct.id
		WHERE c.patient_id = $1
	`

	rows, err := r.conn.QueryContext(ctx, query, patientID)
	if err != nil {
		return nil, fmt.Errorf("could not execute query to get patient contacts for patient ID %s, err: %w", patientID, err)
	}

	contacts := make([]models.Contact, 0)

	for rows.Next() {
		var contact models.Contact

		err = rows.Scan(&contact.Info, &contact.Type.ID, &contact.Type.Name)
		if err != nil {
			return nil, fmt.Errorf("could not scan contact: %w", err)
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}
