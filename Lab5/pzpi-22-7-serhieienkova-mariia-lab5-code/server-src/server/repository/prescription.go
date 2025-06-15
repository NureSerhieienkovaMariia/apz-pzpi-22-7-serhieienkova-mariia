package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PrescriptionPostgres struct {
	db *sqlx.DB
}

func (r *PrescriptionPostgres) GetAllByTreatmentPlanId(id int) ([]structures.Prescription, error) {
	var prescriptions []structures.Prescription
	query := fmt.Sprintf("SELECT p.id, p.treatment_plan_id, p.medicine_id, p.dosage, p.frequency FROM %s p JOIN %s tp ON p.treatment_plan_id = tp.id WHERE tp.id = $1", prescriptionsTable, treatmentPlansTable)
	err := r.db.Select(&prescriptions, query, id)
	if err != nil {
		return nil, err
	}
	return prescriptions, nil
}

func NewPrescriptionPostgres(db *sqlx.DB) *PrescriptionPostgres {
	return &PrescriptionPostgres{db: db}
}

func (r *PrescriptionPostgres) Create(prescription structures.Prescription) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (treatment_plan_id, medicine_id, dosage, frequency) VALUES ($1, $2, $3, $4) RETURNING id", prescriptionsTable)
	row := r.db.QueryRow(query, prescription.TreatmentPlanID, prescription.MedicineID, prescription.Dosage, prescription.Frequency)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PrescriptionPostgres) GetAll() ([]structures.Prescription, error) {
	var prescriptions []structures.Prescription
	query := fmt.Sprintf("SELECT * FROM %s", prescriptionsTable)
	err := r.db.Select(&prescriptions, query)
	if err != nil {
		return nil, err
	}
	return prescriptions, nil
}

func (r *PrescriptionPostgres) Get(id int) (structures.Prescription, error) {
	var prescription structures.Prescription
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", prescriptionsTable)
	err := r.db.Get(&prescription, query, id)
	if err != nil {
		return prescription, err
	}
	return prescription, nil
}

func (r *PrescriptionPostgres) Update(id int, input structures.Prescription) error {
	query := "UPDATE prescription SET"
	var args []interface{}
	argId := 1

	if input.TreatmentPlanID != 0 {
		query += fmt.Sprintf(" treatment_plan_id=$%d,", argId)
		args = append(args, input.TreatmentPlanID)
		argId++
	}
	if input.MedicineID != 0 {
		query += fmt.Sprintf(" medicine_id=$%d,", argId)
		args = append(args, input.MedicineID)
		argId++
	}
	if input.Dosage != "" {
		query += fmt.Sprintf(" dosage=$%d,", argId)
		args = append(args, input.Dosage)
		argId++
	}
	if input.Frequency != "" {
		query += fmt.Sprintf(" frequency=$%d,", argId)
		args = append(args, input.Frequency)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PrescriptionPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", prescriptionsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
