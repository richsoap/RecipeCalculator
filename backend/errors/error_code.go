package errors

//go:generate stringer -type errorCode
type errorCode uint32

const (
	NOT_INITIALIZED errorCode = iota
)

func (e errorCode) Error() string {
	return e.String()
}
