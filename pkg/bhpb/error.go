package bhpb

func (m *Error) Error() string {
	return m.String() + " " + m.Message
}

func GetErrorCode(err error) ErrorCode {
	if e, ok := err.(*Error); ok {
		return e.Code
	}
	return ErrorCode_UNKNOWN_ERROR
}

func ToError(err error) *Error {
	if e, ok := err.(*Error); ok {
		return e
	}
	return nil
}

func IsAlreadyExist(err error) bool {
	return GetErrorCode(err) == ErrorCode_ALREADY_EXISTS
}

func IsNotFound(err error) bool {
	return GetErrorCode(err) == ErrorCode_NOT_FOUND
}
