package postgres

import (
	"context"
	"fmt"

	"github.com/Kvothe838/test-diagnosis/internal/models"
)

func (r *repository) CreateDiagnosis(ctx context.Context, diagnosis models.Diagnosis) (models.Diagnosis, error) {
	query := `
		INSERT INTO diagnoses (id, patient_id, description, prescription, date)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.conn.ExecContext(ctx, query, diagnosis.ID, diagnosis.Patient.ID, diagnosis.Description, diagnosis.Prescription, diagnosis.Date)
	if err != nil {
		return models.Diagnosis{}, fmt.Errorf("could not execute create diagnosis query: %w", err)
	}

	diagnosis.Patient, err = r.getPatientByID(ctx, diagnosis.Patient.ID)
	if err != nil {
		return models.Diagnosis{}, fmt.Errorf("could not get patient after creating diagnosis: %w", err)
	}

	return diagnosis, nil
}
func (r *repository) SearchDiagnoses(ctx context.Context, filters models.SearchDiagnosesFilters) ([]models.Diagnosis, error) {
	baseQuery := `
    SELECT d.id, d.patient_id, d.date, d.description, d.prescription 
    FROM diagnoses d
    JOIN patients p ON d.patient_id = p.id
	WHERE 1=1
	`
	query, args := buildQuery(baseQuery, filters)
	rows, err := r.conn.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("could not execute search diagnoses query: %w", err)
	}

	defer rows.Close()

	diagnoses := make([]models.Diagnosis, 0)

	for rows.Next() {
		var diagnose models.Diagnosis

		err = rows.Scan(&diagnose.ID, &diagnose.Patient.ID, &diagnose.Date, &diagnose.Description, &diagnose.Prescription)

		if err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}

		diagnose.Patient, err = r.getPatientByID(ctx, diagnose.Patient.ID)
		if err != nil {
			return nil, fmt.Errorf("could not get patient by ID %s for diagnose ID %s, err: %w",
				diagnose.Patient.ID, diagnose.ID, err)
		}

		diagnoses = append(diagnoses, diagnose)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error encountered during iteration of rows: %w", err)
	}

	return diagnoses, nil
}

func buildQuery(baseQuery string, filters models.SearchDiagnosesFilters) (string, map[string]interface{}) {
	query := baseQuery
	args := make(map[string]interface{})

	if !filters.Date.IsZero() {
		query += " AND d.date = :date"
		args["date"] = filters.Date
	}

	if len(filters.PatientName) != 0 {
		query += " AND LOWER(CONCAT(p.name, ' ', p.surname)) LIKE LOWER(:full_name)"
		args["full_name"] = "%" + filters.PatientName + "%"
	}

	return query, args
}
