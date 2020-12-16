//go:generate gobin -m -run github.com/cheekybits/genny -in=$GOFILE -out=gen.$GOFILE gen "Port=Int,IntBetween,ListenPort"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

package appcfg

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Port) Value(err *error) (val int) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}
