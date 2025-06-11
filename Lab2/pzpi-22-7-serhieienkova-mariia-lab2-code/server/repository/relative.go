package repository

import (
	"fmt"

	"clinic/server/structures"
	"github.com/jmoiron/sqlx"
)

type RelativePostgres struct {
	db *sqlx.DB
}

func (r *RelativePostgres) GetAllByPatientId(id int) ([]structures.Relative, error) {
	var relatives []structures.Relative
	query := fmt.Sprintf("SELECT r.id, r.email, r.name, r.surname FROM %s r JOIN %s pr ON r.id = pr.relative_id WHERE pr.patient_id = $1", relativesTable, patientsRelativesTable)
	err := r.db.Select(&relatives, query, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred during 'get all relatives by patient id' query: %w", err)
	}
	return relatives, nil
}

func NewRelativePostgres(db *sqlx.DB) *RelativePostgres {
	return &RelativePostgres{db: db}
}

func (r *RelativePostgres) Create(relative structures.Relative) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, surname, password_hash) values ($1, $2, $3, $4) RETURNING id", relativesTable)

	row := r.db.QueryRow(query, relative.Email, relative.Name, relative.Surname, relative.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RelativePostgres) GetByCreds(email string, passwordHash string) (structures.Relative, error) {
	var relative structures.Relative
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password_hash=$2", relativesTable)
	err := r.db.Get(&relative, query, email, passwordHash)
	return relative, err
}

func (r *RelativePostgres) GetById(relativeId int) (structures.Relative, error) {
	var relative structures.Relative

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, relativesTable)
	err := r.db.Get(&relative, query, relativeId)
	if err != nil {
		return relative, fmt.Errorf("error occurred during 'get relative by id' query: %w", err)
	}
	return relative, err
}
