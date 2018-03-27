package utils

//DuplicaPodIPError error for duplicate pod
type DuplicaPodIPError struct {
	message string
}

//NewDuplicaPodIPError creates a new duplicateIPError
func NewDuplicaPodIPError(message string) *DuplicaPodIPError {
	return &DuplicaPodIPError{
		message: message,
	}
}

func (e *DuplicaPodIPError) Error() string {
	return e.message
}

//DuplicaServiceIPError error for duplicate service Error
type DuplicaServiceIPError struct {
	message string
}

//NewDuplicaServiceIPError creates a new duplicateIPError
func NewDuplicaServiceIPError(message string) *DuplicaServiceIPError {
	return &DuplicaServiceIPError{
		message: message,
	}
}

func (e *DuplicaServiceIPError) Error() string {
	return e.message
}

