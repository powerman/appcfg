// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny


package appcfg

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *String) Value(err *error) (val string) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}


// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *OneOfString) Value(err *error) (val string) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}


// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Endpoint) Value(err *error) (val string) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}
