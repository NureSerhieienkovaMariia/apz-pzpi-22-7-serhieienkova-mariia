package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PatientRelativePostgres struct {
	db *sqlx.DB
}

func NewPatientRelativePostgres(db *sqlx.DB) *PatientRelativePostgres {
	return &PatientRelativePostgres{db: db}
}

func (r *PatientRelativePostgres) Create(patientRelative structures.PatientRelative) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (patient_id, relative_id, access_to_records) VALUES ($1, $2, $3) RETURNING id", patientRelativeTable)
	row := r.db.QueryRow(query, patientRelative.PatientID, patientRelative.RelativeID, patientRelative.AccessToRecords)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PatientRelativePostgres) GetAll() ([]structures.PatientRelative, error) {
	var patientRelatives []structures.PatientRelative
	query := fmt.Sprintf("SELECT * FROM %s", patientRelativeTable)
	err := r.db.Select(&patientRelatives, query)
	if err != nil {
		return nil, err
	}
	return patientRelatives, nil
}

func (r *PatientRelativePostgres) Get(id int) (structures.PatientRelative, error) {
	var patientRelative structures.PatientRelative
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", patientRelativeTable)
	err := r.db.Get(&patientRelative, query, id)
	if err != nil {
		return patientRelative, err
	}
	return patientRelative, nil
}

func (r *PatientRelativePostgres) GetAllByPatientId(patientId int) ([]structures.PatientRelative, error) {
	var patientRelatives []structures.PatientRelative
	query := fmt.Sprintf("SELECT * FROM %s WHERE patient_id = $1", patientRelativeTable)
	err := r.db.Select(&patientRelatives, query, patientId)
	if err != nil {
		return nil, err
	}
	return patientRelatives, nil
}

func (r *PatientRelativePostgres) GetAllByRelativeId(relativeId int) ([]structures.PatientRelative, error) {
	var patientRelatives []structures.PatientRelative
	query := fmt.Sprintf("SELECT * FROM %s WHERE relative_id = $1", patientRelativeTable)
	err := r.db.Select(&patientRelatives, query, relativeId)
	if err != nil {
		return nil, err
	}
	return patientRelatives, nil
}

func (r *PatientRelativePostgres) Update(id int, input structures.PatientRelative) error {
	query := "UPDATE patient_relatives SET"
	var args []interface{}
	argId := 1

	if input.PatientID != 0 {
		query += fmt.Sprintf(" patient_id=$%d,", argId)
		args = append(args, input.PatientID)
		argId++
	}
	if input.RelativeID != 0 {
		query += fmt.Sprintf(" relative_id=$%d,", argId)
		args = append(args, input.RelativeID)
		argId++
	}
	query += fmt.Sprintf(" access_to_records=$%d,", argId)
	args = append(args, input.AccessToRecords)
	argId++

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PatientRelativePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", patientRelativeTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
