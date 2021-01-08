package data

type Error string

const (
	NotFoundError            = Error("Not found")
	AlreadyExistsError       = Error("Already exists")
	ForeingKeyViolationError = Error("Foreign key violation")
)

func (e Error) Error() string {
	return string(e)
}
