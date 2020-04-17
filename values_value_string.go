//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "NotEmptyString=String,OneOfString,Endpoint"

package appcfg

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *NotEmptyString) Value(err *error) (val string) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}
