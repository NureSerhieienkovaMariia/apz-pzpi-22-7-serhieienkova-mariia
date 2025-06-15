package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type HealthNotePostgres struct {
	db *sqlx.DB
}

func (r *HealthNotePostgres) GetAllByPatientId(id int) ([]structures.HealthNote, error) {
	var healthNotes []structures.HealthNote
	query := fmt.Sprintf("SELECT * FROM %s WHERE patient_id = $1", patientNotesTable)
	err := r.db.Select(&healthNotes, query, id)
	if err != nil {
		return nil, err
	}
	return healthNotes, nil
}

func NewHealthNotePostgres(db *sqlx.DB) *HealthNotePostgres {
	return &HealthNotePostgres{db: db}
}

func (r *HealthNotePostgres) Create(healthNote structures.HealthNote) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (patient_id, timestamp, note) VALUES ($1, $2, $3) RETURNING id", patientNotesTable)
	row := r.db.QueryRow(query, healthNote.PatientID, healthNote.Timestamp, healthNote.Note)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *HealthNotePostgres) GetAll() ([]structures.HealthNote, error) {
	var healthNotes []structures.HealthNote
	query := fmt.Sprintf("SELECT * FROM %s", patientNotesTable)
	err := r.db.Select(&healthNotes, query)
	if err != nil {
		return nil, err
	}
	return healthNotes, nil
}

func (r *HealthNotePostgres) Get(id int) (structures.HealthNote, error) {
	var healthNote structures.HealthNote
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", patientNotesTable)
	err := r.db.Get(&healthNote, query, id)
	if err != nil {
		return healthNote, err
	}
	return healthNote, nil
}

func (r *HealthNotePostgres) Update(id int, input structures.HealthNote) error {
	query := "UPDATE health_notes SET"
	var args []interface{}
	argId := 1

	if input.PatientID != 0 {
		query += fmt.Sprintf(" patient_id=$%d,", argId)
		args = append(args, input.PatientID)
		argId++
	}
	if !input.Timestamp.IsZero() {
		query += fmt.Sprintf(" timestamp=$%d,", argId)
		args = append(args, input.Timestamp)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *HealthNotePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", patientNotesTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
