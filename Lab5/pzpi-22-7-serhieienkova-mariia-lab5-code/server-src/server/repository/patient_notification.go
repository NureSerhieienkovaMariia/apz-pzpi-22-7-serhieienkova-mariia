package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PatientNotificationPostgres struct {
	db *sqlx.DB
}

func NewPatientNotificationPostgres(db *sqlx.DB) *PatientNotificationPostgres {
	return &PatientNotificationPostgres{db: db}
}

func (r *PatientNotificationPostgres) Create(notification structures.PatientNotification) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (timestamp, patient_id, topic, message) VALUES ($1, $2, $3, $4) RETURNING id", patientNotificationTable)
	row := r.db.QueryRow(query, notification.Timestamp, notification.PatientID, notification.Topic, notification.Message)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PatientNotificationPostgres) GetAll() ([]structures.PatientNotification, error) {
	var notifications []structures.PatientNotification
	query := fmt.Sprintf("SELECT * FROM %s", patientNotificationTable)
	err := r.db.Select(&notifications, query)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *PatientNotificationPostgres) Get(id int) (structures.PatientNotification, error) {
	var notification structures.PatientNotification
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", patientNotificationTable)
	err := r.db.Get(&notification, query, id)
	if err != nil {
		return notification, err
	}
	return notification, nil
}

func (r *PatientNotificationPostgres) Update(id int, input structures.PatientNotification) error {
	query := "UPDATE patient_notifications SET"
	var args []interface{}
	argId := 1

	if !input.Timestamp.IsZero() {
		query += fmt.Sprintf(" timestamp=$%d,", argId)
		args = append(args, input.Timestamp)
		argId++
	}
	if input.PatientID != 0 {
		query += fmt.Sprintf(" patient_id=$%d,", argId)
		args = append(args, input.PatientID)
		argId++
	}
	if input.Topic != "" {
		query += fmt.Sprintf(" topic=$%d,", argId)
		args = append(args, input.Topic)
		argId++
	}
	if input.Message != "" {
		query += fmt.Sprintf(" message=$%d,", argId)
		args = append(args, input.Message)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PatientNotificationPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", patientNotificationTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
