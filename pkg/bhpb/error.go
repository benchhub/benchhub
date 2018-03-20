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

func IsAgentAlreadyReigstered(err error) bool {
	return GetErrorCode(err) == ErrorCode_AGENT_ALREADY_REGISTED
}
