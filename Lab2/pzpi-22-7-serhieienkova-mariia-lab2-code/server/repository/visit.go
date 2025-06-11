package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type VisitPostgres struct {
	db *sqlx.DB
}

func (r *VisitPostgres) GetAllByTreatmentPlanId(id int) ([]structures.Visit, error) {
	var visits []structures.Visit
	query := fmt.Sprintf("SELECT * FROM %s WHERE treatment_plan_id = $1", visitsTable)
	err := r.db.Select(&visits, query, id)
	if err != nil {
		return nil, err
	}
	return visits, nil
}

func (r *VisitPostgres) GetVisitsForNextDays(i int) ([]structures.Visit, error) {
	var visits []structures.Visit
	query := fmt.Sprintf("SELECT * FROM %s WHERE date BETWEEN NOW() AND NOW() + INTERVAL '%d days'", visitsTable, i)
	err := r.db.Select(&visits, query)
	if err != nil {
		return nil, err
	}
	return visits, nil
}

func NewVisitPostgres(db *sqlx.DB) *VisitPostgres {
	return &VisitPostgres{db: db}
}

func (r *VisitPostgres) Create(visit structures.Visit) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (treatment_plan_id, reason, date, note) VALUES ($1, $2, $3, $4) RETURNING id", visitsTable)
	row := r.db.QueryRow(query, visit.TreatmentPlanID, visit.Reason, visit.Date, visit.Notes)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *VisitPostgres) GetAll() ([]structures.Visit, error) {
	var visits []structures.Visit
	query := fmt.Sprintf("SELECT * FROM %s", visitsTable)
	err := r.db.Select(&visits, query)
	if err != nil {
		return nil, err
	}
	return visits, nil
}

func (r *VisitPostgres) Get(id int) (structures.Visit, error) {
	var visit structures.Visit
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", visitsTable)
	err := r.db.Get(&visit, query, id)
	if err != nil {
		return visit, err
	}
	return visit, nil
}

func (r *VisitPostgres) Update(id int, input structures.Visit) error {
	query := "UPDATE visits SET"
	var args []interface{}
	argId := 1

	if input.TreatmentPlanID != 0 {
		query += fmt.Sprintf(" treatment_plan_id=$%d,", argId)
		args = append(args, input.TreatmentPlanID)
		argId++
	}
	if input.Reason != "" {
		query += fmt.Sprintf(" reason=$%d,", argId)
		args = append(args, input.Reason)
		argId++
	}
	if !input.Date.IsZero() {
		query += fmt.Sprintf(" date=$%d,", argId)
		args = append(args, input.Date)
		argId++
	}
	if input.Notes != "" {
		query += fmt.Sprintf(" notes=$%d,", argId)
		args = append(args, input.Notes)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *VisitPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", visitsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
