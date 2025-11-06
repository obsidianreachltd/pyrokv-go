package errors

const (
	// Server Returned Error Codes
	ErrKeyNotFound      = 0x01
	ErrKeyValueExpired  = 0x02
	ErrKeyValueTooLarge = 0x03
	ErrBadRequest       = 0x04
	ErrInternalServer   = 0x05
	// Client Side Error Codes
	ErrClientTimeout = 0xF1
	ErrClientClosed  = 0xF2
)

type PyroKVError struct {
	Code byte
	Msg  string
}

func NewPyroKVError(code byte, msg string) *PyroKVError {
	return &PyroKVError{
		Code: code,
		Msg:  msg,
	}
}

func (e *PyroKVError) Error() string {
	return e.Msg
}

func ErrorFromResponsePayload(payload byte) error {
	switch payload {
	case ErrKeyNotFound:
		return NewPyroKVError(ErrKeyNotFound, "key not found")
	case ErrKeyValueExpired:
		return NewPyroKVError(ErrKeyValueExpired, "key value expired")
	case ErrKeyValueTooLarge:
		return NewPyroKVError(ErrKeyValueTooLarge, "key value too large")
	case ErrBadRequest:
		return NewPyroKVError(ErrBadRequest, "bad request")
	case ErrInternalServer:
		return NewPyroKVError(ErrInternalServer, "internal server error")
	case ErrClientTimeout:
		return NewPyroKVError(ErrClientTimeout, "client timeout")
	case ErrClientClosed:
		return NewPyroKVError(ErrClientClosed, "client closed")
	default:
		return NewPyroKVError(0xFF, "unknown error")
	}
}

func isPyroKVError(err error, code byte) bool {
	if pkvErr, ok := err.(*PyroKVError); ok {
		return pkvErr.Code == code
	}
	return false
}

func IsKeyNotFoundError(err error) bool {
	return isPyroKVError(err, ErrKeyNotFound)
}

func IsKeyValueExpiredError(err error) bool {
	return isPyroKVError(err, ErrKeyValueExpired)
}

func IsKeyValueTooLargeError(err error) bool {
	return isPyroKVError(err, ErrKeyValueTooLarge)
}

func IsBadRequestError(err error) bool {
	return isPyroKVError(err, ErrBadRequest)
}

func IsInternalServerError(err error) bool {
	return isPyroKVError(err, ErrInternalServer)
}

func IsClientTimeoutError(err error) bool {
	return isPyroKVError(err, ErrClientTimeout)
}

func IsClientClosedError(err error) bool {
	return isPyroKVError(err, ErrClientClosed)
}
