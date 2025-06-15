package structures

type Prescription struct {
	ID              int    `json:"id" db:"id"`
	TreatmentPlanID int    `json:"treatment_plan_id" db:"treatment_plan_id"`
	MedicineID      int    `json:"medicine_id" db:"medicine_id"`
	Dosage          string `json:"dosage" db:"dosage"`
	Frequency       string `json:"frequency" db:"frequency"`
}
