package errors

type Error int

func (e Error) Error() string {
	return errorMap[e]
}
