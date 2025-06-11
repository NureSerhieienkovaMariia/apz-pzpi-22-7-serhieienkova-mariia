package structures

type Diagnosis struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Description    string `json:"description" db:"description"`
	Recommendation string `json:"recommendations" db:"recomendations"`
}
