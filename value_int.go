//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "Port=ListenPort"

package appcfg

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Port) Value(err *error) (val int) {
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}
