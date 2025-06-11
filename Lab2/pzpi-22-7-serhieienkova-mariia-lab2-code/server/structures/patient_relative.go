package structures

type PatientRelative struct {
	ID              int  `json:"id" db:"id"`
	PatientID       int  `json:"patient_id" db:"patient_id"`
	RelativeID      int  `json:"relative_id" db:"relative_id"`
	AccessToRecords bool `json:"access_to_records" db:"access_to_records"`
}
