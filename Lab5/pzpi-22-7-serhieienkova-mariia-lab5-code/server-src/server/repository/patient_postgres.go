package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"time"
)

type PatientPostgres struct {
	db *sqlx.DB
}

func (r *PatientPostgres) AttachRelative(relative structures.PatientRelative) error {
	query := fmt.Sprintf("INSERT INTO %s (patient_id, relative_id, access_to_records) VALUES ($1, $2, $3)", patientsRelativesTable)
	_, err := r.db.Exec(query, relative.PatientID, relative.RelativeID, relative.AccessToRecords)
	if err != nil {
		return fmt.Errorf("error occurred during 'attach relative to patient' query: %w", err)
	}
	return nil
}

func (r *PatientPostgres) GetAllByRelativeId(id int) ([]structures.Patient, error) {
	var patients []structures.Patient
	query := fmt.Sprintf("SELECT p.id, p.name, p.surname, p.birthday, p.email FROM %s p JOIN %s pr ON p.id = pr.patient_id WHERE pr.relative_id = $1", patientsTable, patientsRelativesTable)
	err := r.db.Select(&patients, query, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred during 'get all patients by relative id' query: %w", err)
	}
	return patients, nil
}

func (r *PatientPostgres) GetAllByDoctorId(id int) ([]structures.Patient, error) {
	var patients []structures.Patient
	query := fmt.Sprintf("SELECT p.id, p.name, p.surname, p.birthday FROM %s p JOIN %s tp ON p.id = tp.patient_id WHERE tp.doctor_id = $1", patientsTable, treatmentPlansTable)
	err := r.db.Select(&patients, query, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred during 'get all patients by doctor id' query: %w", err)
	}
	return patients, nil
}

func NewPatientPostgres(db *sqlx.DB) *PatientPostgres {
	return &PatientPostgres{db: db}
}

func (r *PatientPostgres) Create(patient structures.Patient) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, surname, password_hash, birthday, sex) values ($1, $2, $3, $4, $5, $6) RETURNING id", patientsTable)
	row := r.db.QueryRow(query, patient.Email, patient.Name, patient.Surname, patient.PasswordHash, patient.Birthday, patient.Sex)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PatientPostgres) GetAll() ([]structures.Patient, error) {
	var patients []structures.Patient
	query := fmt.Sprintf("SELECT * FROM %s", patientsTable)
	err := r.db.Select(&patients, query)
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *PatientPostgres) GetById(id int) (structures.Patient, error) {
	var patient structures.Patient
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", patientsTable)
	err := r.db.Get(&patient, query, id)
	if err != nil {
		return patient, err
	}
	return patient, nil
}

func (r *PatientPostgres) Update(id int, input structures.Patient) error {
	query := fmt.Sprintf("UPDATE %s SET", patientsTable)
	var args []interface{}
	argId := 1

	if input.Name != "" {
		query += fmt.Sprintf(" name=$%d,", argId)
		args = append(args, input.Name)
		argId++
	}
	if input.Surname != "" {
		query += fmt.Sprintf(" surname=$%d,", argId)
		args = append(args, input.Surname)
		argId++
	}
	if !reflect.DeepEqual(input.Birthday, time.Time{}) {
		query += fmt.Sprintf(" birthday=$%d,", argId)
		args = append(args, input.Birthday)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PatientPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", patientsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PatientPostgres) GetDiagnosesByPatientID(patientID int) ([]structures.Diagnosis, error) {
	var diagnoses []structures.Diagnosis
	query := `SELECT d.* FROM ` + diagnosesTable + ` d
              JOIN ` + patientsTable + ` p ON p.diagnosis_id = d.id
              WHERE p.id = $1`
	err := r.db.Select(&diagnoses, query, patientID)
	return diagnoses, err
}

func (r *PatientPostgres) GetMedicinesByPatientID(patientID int) ([]structures.Medicine, error) {
	var medicines []structures.Medicine
	query := `SELECT m.* FROM ` + medicinesTable + ` m
              JOIN ` + patientMedicineTable + ` pm ON pm.medicine_id = m.id
              WHERE pm.patient_id = $1`
	err := r.db.Select(&medicines, query, patientID)
	return medicines, err
}

func (r *PatientPostgres) GetDevicesByPatientID(patientID int) ([]structures.Device, error) {
	var devices []structures.Device
	query := `SELECT * FROM ` + devicesTable + ` WHERE patient_id = $1`
	err := r.db.Select(&devices, query, patientID)
	return devices, err
}

func (r *PatientPostgres) GetIndicatorsByPatientID(patientID int) ([]structures.IndicatorsStamp, error) {
	var indicators []structures.IndicatorsStamp
	query := `SELECT * FROM ` + indicatorsStampsTable + ` WHERE device_id IN 
              (SELECT id FROM ` + devicesTable + ` WHERE patient_id = $1)`
	err := r.db.Select(&indicators, query, patientID)
	return indicators, err
}

func (r *PatientPostgres) GetIndicatorsNotificationsByPatientID(patientID int) ([]structures.IndicatorsNotification, error) {
	var indicatorsNotifications []structures.IndicatorsNotification
	query := `SELECT n.* FROM ` + indicatorsNotificationsTable + ` n
              JOIN ` + indicatorsStampsTable + ` ind ON ind.id = n.indicator_stamp_id
              JOIN ` + devicesTable + ` d ON d.id = ind.device_id
              WHERE d.patient_id = $1`
	err := r.db.Select(&indicatorsNotifications, query, patientID)
	return indicatorsNotifications, err
}

func (r *PatientPostgres) GetByCreds(email string, hash string) (structures.Patient, error) {
	var patient structures.Patient
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1 AND password_hash = $2", patientsTable)
	err := r.db.Get(&patient, query, email, hash)
	if err != nil {
		return structures.Patient{}, err
	}
	return patient, nil
}
