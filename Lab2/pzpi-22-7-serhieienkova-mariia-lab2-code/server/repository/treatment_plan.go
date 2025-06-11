package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TreatmentPlanPostgres struct {
	db *sqlx.DB
}

func (r *TreatmentPlanPostgres) GetByPatientId(id int) ([]structures.TreatmentPlan, error) {
	var treatmentPlans []structures.TreatmentPlan
	query := fmt.Sprintf("SELECT * FROM %s WHERE patient_id = $1", treatmentPlansTable)
	err := r.db.Select(&treatmentPlans, query, id)
	if err != nil {
		return nil, err
	}
	return treatmentPlans, nil
}

func (r *TreatmentPlanPostgres) GetAllByDoctorId(id int) ([]structures.TreatmentPlan, error) {
	var treatmentPlans []structures.TreatmentPlan
	query := fmt.Sprintf("SELECT * FROM %s WHERE doctor_id = $1", treatmentPlansTable)
	err := r.db.Select(&treatmentPlans, query, id)
	if err != nil {
		return nil, err
	}
	return treatmentPlans, nil
}

func NewTreatmentPlanPostgres(db *sqlx.DB) *TreatmentPlanPostgres {
	return &TreatmentPlanPostgres{db: db}
}

func (r *TreatmentPlanPostgres) Create(treatmentPlan structures.TreatmentPlan) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (patient_id, doctor_id, start_date, end_date) VALUES ($1, $2, $3, $4) RETURNING id", treatmentPlansTable)
	row := r.db.QueryRow(query, treatmentPlan.PatientID, treatmentPlan.DoctorID, treatmentPlan.StartDate, treatmentPlan.EndDate)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TreatmentPlanPostgres) GetAll() ([]structures.TreatmentPlan, error) {
	var treatmentPlans []structures.TreatmentPlan
	query := fmt.Sprintf("SELECT * FROM %s", treatmentPlansTable)
	err := r.db.Select(&treatmentPlans, query)
	if err != nil {
		return nil, err
	}
	return treatmentPlans, nil
}

func (r *TreatmentPlanPostgres) Get(id int) (structures.TreatmentPlan, error) {
	var treatmentPlan structures.TreatmentPlan
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", treatmentPlansTable)
	err := r.db.Get(&treatmentPlan, query, id)
	if err != nil {
		return treatmentPlan, err
	}
	return treatmentPlan, nil
}

func (r *TreatmentPlanPostgres) Update(id int, input structures.TreatmentPlan) error {
	query := "UPDATE treatment_plans SET"
	var args []interface{}
	argId := 1

	if input.PatientID != 0 {
		query += fmt.Sprintf(" patient_id=$%d,", argId)
		args = append(args, input.PatientID)
		argId++
	}
	if input.DoctorID != 0 {
		query += fmt.Sprintf(" doctor_id=$%d,", argId)
		args = append(args, input.DoctorID)
		argId++
	}
	if !input.StartDate.IsZero() {
		query += fmt.Sprintf(" start_date=$%d,", argId)
		args = append(args, input.StartDate)
		argId++
	}
	if !input.EndDate.IsZero() {
		query += fmt.Sprintf(" end_date=$%d,", argId)
		args = append(args, input.EndDate)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TreatmentPlanPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", treatmentPlansTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
