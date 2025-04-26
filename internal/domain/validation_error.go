package domain

type ValidationError struct {
	Message string
	Errors  map[string]string
}

func (ve *ValidationError) Error() string {
	return ve.Message
}
