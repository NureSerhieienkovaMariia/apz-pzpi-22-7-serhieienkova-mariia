package repository

import (
	"fmt"

	"clinic/server/structures"
	"github.com/jmoiron/sqlx"
)

type DoctorPostgres struct {
	db *sqlx.DB
}

func NewDoctorPostgres(db *sqlx.DB) *DoctorPostgres {
	return &DoctorPostgres{db: db}
}

func (r *DoctorPostgres) Create(doctor structures.Doctor) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, surname, password_hash) values ($1, $2, $3, $4) RETURNING id", doctorsTable)

	row := r.db.QueryRow(query, doctor.Email, doctor.Name, doctor.Surname, doctor.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *DoctorPostgres) GetByCreds(email string, passwordHash string) (structures.Doctor, error) {
	var doctor structures.Doctor
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password_hash=$2", doctorsTable)
	err := r.db.Get(&doctor, query, email, passwordHash)
	return doctor, err
}

func (r *DoctorPostgres) GetById(doctorId int) (structures.Doctor, error) {
	var doctor structures.Doctor

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, doctorsTable)
	err := r.db.Get(&doctor, query, doctorId)
	if err != nil {
		return doctor, fmt.Errorf("error occurred during 'get doctor by id' query: %w", err)
	}
	return doctor, err
}
