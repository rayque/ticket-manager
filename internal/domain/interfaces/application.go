package interfaces

type UUID interface {
	NewUUID() string
	Generate() string
}

type DuplicateEmailError struct {
	Email string
}

func (e *DuplicateEmailError) Error() string {
	return "user with email " + e.Email + " already exists"
}
