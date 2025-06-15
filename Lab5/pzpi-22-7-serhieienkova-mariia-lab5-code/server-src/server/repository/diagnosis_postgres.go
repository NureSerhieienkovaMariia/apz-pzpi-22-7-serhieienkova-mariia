package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DiagnosisPostgres struct {
	db *sqlx.DB
}

func (r *DiagnosisPostgres) GetAllByPatientId(id int) ([]structures.Diagnosis, error) {
	var diagnoses []structures.Diagnosis
	query := fmt.Sprintf("SELECT d.id, d.name, d.description, d.recomendations FROM %s d JOIN %s pd ON d.id = pd.diagnosis_id WHERE pd.patient_id = $1", diagnosesTable, patientsDiagnosesTable)
	err := r.db.Select(&diagnoses, query, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred during 'get all diagnoses by patient id' query: %w", err)
	}
	return diagnoses, nil
}

func (r *DiagnosisPostgres) AttachDiagnosisToPatient(diagnosis structures.PatientDiagnosis) error {
	query := fmt.Sprintf("INSERT INTO %s (patient_id, diagnosis_id) VALUES ($1, $2)", patientsDiagnosesTable)
	_, err := r.db.Exec(query, diagnosis.PatientId, diagnosis.DiagnosisId)
	if err != nil {
		return fmt.Errorf("error occurred during 'attach diagnosis to patient' query: %w", err)
	}
	return nil
}

func NewDiagnosisPostgres(db *sqlx.DB) *DiagnosisPostgres {
	return &DiagnosisPostgres{db: db}
}

func (r *DiagnosisPostgres) Create(diagnosis structures.Diagnosis) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description, recomendations) values ($1, $2, $3) RETURNING id", diagnosesTable)
	row := r.db.QueryRow(query, diagnosis.Name, diagnosis.Description, diagnosis.Recommendation)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DiagnosisPostgres) GetAll() ([]structures.Diagnosis, error) {
	var diagnoses []structures.Diagnosis
	query := fmt.Sprintf("SELECT * FROM %s", diagnosesTable)
	err := r.db.Select(&diagnoses, query)
	if err != nil {
		return nil, err
	}
	return diagnoses, nil
}

func (r *DiagnosisPostgres) Get(id int) (structures.Diagnosis, error) {
	var diagnosis structures.Diagnosis
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", diagnosesTable)
	err := r.db.Get(&diagnosis, query, id)
	if err != nil {
		return diagnosis, err
	}
	return diagnosis, nil
}

func (r *DiagnosisPostgres) Update(id int, input structures.Diagnosis) error {
	query := "UPDATE diagnoses SET"
	var args []interface{}
	argId := 1

	if input.Name != "" {
		query += fmt.Sprintf(" name=$%d,", argId)
		args = append(args, input.Name)
		argId++
	}
	if input.Description != "" {
		query += fmt.Sprintf(" description=$%d,", argId)
		args = append(args, input.Description)
		argId++
	}
	if input.Recommendation != "" {
		query += fmt.Sprintf(" recomendations=$%d,", argId)
		args = append(args, input.Recommendation)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *DiagnosisPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", diagnosesTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
