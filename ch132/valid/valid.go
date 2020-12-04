package valid

type Validator interface {
	Validate(username, password string) bool
}

type validatorImpl struct{}

func New() Validator {
	return &validatorImpl{}
}

func (v *validatorImpl) Validate(username, password string) bool {
	return username == "alice" && password == "bob"
}
