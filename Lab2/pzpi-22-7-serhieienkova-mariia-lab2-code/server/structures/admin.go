package structures

const AdminType string = "admin"

type Admin struct {
	User
}

func (a Admin) Validate() error {
	return nil
}
