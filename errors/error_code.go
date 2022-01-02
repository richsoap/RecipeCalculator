package errors

//go:generate stringer -type errorCode
type errorCode uint32

const (
	NOT_INITIALIZED errorCode = iota
	RECIPE_NOT_PROVIDED
	RECIPE_NOT_FOUND
	ITEM_NOT_FOUND
	BROKEN_DATA
	CIRCLE_DEPENDENCY
	CONFLICT_OPTIONS
)

func (e errorCode) Error() string {
	return e.String()
}
