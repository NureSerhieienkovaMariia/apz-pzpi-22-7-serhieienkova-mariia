package repository

import (
	"fmt"

	"clinic/server/structures"
	"github.com/jmoiron/sqlx"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (r *AdminPostgres) Create(admin structures.Admin) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, surname, password_hash) values ($1, $2, $3, $4) RETURNING id", adminsTable)

	row := r.db.QueryRow(query, admin.Email, admin.Name, admin.Surname, admin.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AdminPostgres) GetByCreds(email string, passwordHash string) (structures.Admin, error) {
	var admin structures.Admin
	fmt.Println(fmt.Sprintf("query admin by email: %v, passwordHash: %v", email, passwordHash))
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password_hash=$2", adminsTable)
	fmt.Println(fmt.Sprintf("query: %v", query))
	err := r.db.Get(&admin, query, email, passwordHash)
	return admin, err
}

func (r *AdminPostgres) GetById(adminId int) (structures.Admin, error) {
	var admin structures.Admin

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`,
		adminsTable)
	err := r.db.Get(&admin, query, adminId)
	if err != nil {
		return admin, fmt.Errorf("error occured during 'get admin by id' query: %w", err)
	}
	return admin, err
}
