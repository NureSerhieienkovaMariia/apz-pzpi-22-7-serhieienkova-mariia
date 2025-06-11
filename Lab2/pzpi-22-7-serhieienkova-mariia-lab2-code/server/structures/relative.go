package structures

const RelativeType string = "relative"

type Relative struct {
	User
}

func (r Relative) Validate() error {
	return nil
}
