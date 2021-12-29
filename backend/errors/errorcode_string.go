// Code generated by "stringer -type errorCode"; DO NOT EDIT.

package errors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NOT_INITIALIZED-0]
	_ = x[RECIPE_NOT_PROVIDED-1]
	_ = x[RECIPE_NOT_FOUND-2]
	_ = x[ITEM_NOT_FOUND-3]
	_ = x[BROKEN_DATA-4]
	_ = x[CIRCLE_DEPENDENCY-5]
}

const _errorCode_name = "NOT_INITIALIZEDRECIPE_NOT_PROVIDEDRECIPE_NOT_FOUNDITEM_NOT_FOUNDBROKEN_DATACIRCLE_DEPENDENCY"

var _errorCode_index = [...]uint8{0, 15, 34, 50, 64, 75, 92}

func (i errorCode) String() string {
	if i >= errorCode(len(_errorCode_index)-1) {
		return "errorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _errorCode_name[_errorCode_index[i]:_errorCode_index[i+1]]
}
