package structures

type Device struct {
	Id        int    `json:"id" db:"id"`
	Password  string `json:"password" db:"password_hash"`
	PatientId int    `json:"patient_id" db:"patient_id"`
}
