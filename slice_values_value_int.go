//go:generate mise exec -- genny -in=$GOFILE -out=gen.$GOFILE gen "PortSlice=IntSlice,IntBetweenSlice,ListenPortSlice"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

package appcfg

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *PortSlice) Value(err *error) (val []int) { //nolint:gocritic // ptrToRefParam.
	if v.Get() == nil {
		*err = &RequiredError{v}
		return val
	}
	return v.values
}
