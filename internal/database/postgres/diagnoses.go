package postgres

import (
	"TopDoctorsBackendChallenge/internal/models"
	"context"
	"github.com/pkg/errors"
)

func (r repository) SearchDiagnoses(ctx context.Context, filters models.SearchDiagnosesFilters) ([]models.Diagnose, error) {
	baseQuery := `
    SELECT d.id, d.patient_id, d.date, d.description, d.prescription 
    FROM diagnoses d
    JOIN patients p ON d.patient_id = p.id
	WHERE 1=1
	`
	query, args := buildQuery(baseQuery, filters)
	rows, err := r.conn.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "could not execute search diagnoses query")
	}

	defer rows.Close()

	diagnoses := make([]models.Diagnose, 0)

	for rows.Next() {
		var diagnose models.Diagnose

		err = rows.Scan(&diagnose.ID, &diagnose.Patient.ID, &diagnose.Date, &diagnose.Description, &diagnose.Prescription)

		if err != nil {
			return nil, errors.Wrap(err, "could not scan row")
		}

		diagnose.Patient, err = r.getPatientByID(ctx, diagnose.Patient.ID)
		if err != nil {
			return nil, errors.Wrapf(err, "could not get patient by ID %s for diagnose ID %s",
				diagnose.Patient.ID, diagnose.ID)
		}

		diagnoses = append(diagnoses, diagnose)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error encountered during iteration of rows")
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
		query += " AND CONCAT(p.name, ' ', p.surname) LIKE :full_name"
		args["full_name"] = "%" + filters.PatientName + "%"
	}

	return query, args
}
