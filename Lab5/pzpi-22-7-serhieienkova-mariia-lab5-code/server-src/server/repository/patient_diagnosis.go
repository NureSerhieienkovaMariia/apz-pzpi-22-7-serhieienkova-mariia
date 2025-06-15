package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PatientDiagnosisPostgres struct {
	db *sqlx.DB
}

func NewPatientDiagnosisPostgres(db *sqlx.DB) *PatientDiagnosisPostgres {
	return &PatientDiagnosisPostgres{db: db}
}

func (r *PatientDiagnosisPostgres) Create(patientDiagnosis structures.PatientDiagnosis) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (patient_id, diagnosis_id) VALUES ($1, $2) RETURNING id", patientsDiagnosesTable)
	row := r.db.QueryRow(query, patientDiagnosis.PatientId, patientDiagnosis.DiagnosisId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PatientDiagnosisPostgres) GetAll() ([]structures.PatientDiagnosis, error) {
	var patientDiagnoses []structures.PatientDiagnosis
	query := fmt.Sprintf("SELECT * FROM %s", patientsDiagnosesTable)
	err := r.db.Select(&patientDiagnoses, query)
	if err != nil {
		return nil, err
	}
	return patientDiagnoses, nil
}

func (r *PatientDiagnosisPostgres) Get(id int) (structures.PatientDiagnosis, error) {
	var patientDiagnosis structures.PatientDiagnosis
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", patientsDiagnosesTable)
	err := r.db.Get(&patientDiagnosis, query, id)
	if err != nil {
		return patientDiagnosis, err
	}
	return patientDiagnosis, nil
}

func (r *PatientDiagnosisPostgres) Update(id int, input structures.PatientDiagnosis) error {
	query := "UPDATE patient_diagnoses SET"
	var args []interface{}
	argId := 1

	if input.PatientId != 0 {
		query += fmt.Sprintf(" patient_id=$%d,", argId)
		args = append(args, input.PatientId)
		argId++
	}
	if input.DiagnosisId != 0 {
		query += fmt.Sprintf(" diagnosis_id=$%d,", argId)
		args = append(args, input.DiagnosisId)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PatientDiagnosisPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", patientsDiagnosesTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
