package models

import "time"

type Diagnosis struct {
	ID           string
	Patient      Patient
	Date         time.Time
	Description  string
	Prescription *string
}

type SearchDiagnosesFilters struct {
	PatientName string
	Date        time.Time
}
