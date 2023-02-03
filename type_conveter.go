package gotypconv

// The `String` function returns a pointer to the provided string value.
func String(v string) *string {
	return &v
}

// The `StringValue“ function returns the underlying value of the string pointer passed to it.
// `""` or if the pointer is `nil`
func StringValue(v *string) string {
	return *v
}

// `Bool` functions returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// `BoolValue` returns the value of the bool pointer passed in or
// false if the pointer is nil.
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}
