package structures

type PatientDiagnosis struct {
	Id          int `json:"id" db:"id"`
	PatientId   int `json:"patient_id" db:"patient_id"`
	DiagnosisId int `json:"diagnosis_id" db:"diagnosis_id"`
}
