package structures

const DoctorType string = "doctor"

type Doctor struct {
	User
}

func (d Doctor) Validate() error {
	return nil
}
