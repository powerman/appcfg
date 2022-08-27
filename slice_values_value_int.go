//go:generate -command GENNY sh -c "$(git rev-parse --show-toplevel)/.buildcache/bin/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" genny
//go:generate GENNY -in=$GOFILE -out=gen.$GOFILE gen "PortSlice=IntSlice,IntBetweenSlice,ListenPortSlice"
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
