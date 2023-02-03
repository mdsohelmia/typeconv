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
