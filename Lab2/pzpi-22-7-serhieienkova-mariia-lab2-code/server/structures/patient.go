package structures

import "time"

const PatientType string = "patient"

type Patient struct {
	User
	Birthday time.Time `json:"birthday" db:"birthday"`
	Sex      bool      `json:"sex" db:"sex"`
}

func (p Patient) Validate() error {
	return nil
}

type PrescriptionInfo struct {
	Medicine  Medicine `json:"medicine"`
	Dosage    string   `json:"dosage"`
	Frequency string   `json:"frequency"`
}

type TreatmentPlanFullInfo struct {
	Doctor        Doctor             `json:"doctor"`
	StartDate     time.Time          `json:"start_date"`
	EndDate       time.Time          `json:"end_date"`
	Visits        []Visit            `json:"visits"`
	Prescriptions []PrescriptionInfo `json:"prescriptions"`
}

type PatientFullInfo struct {
	Patient
	Diagnoses      []Diagnosis             `json:"diagnoses"`
	TreatmentPlans []TreatmentPlanFullInfo `json:"treatment_plans"`
}
