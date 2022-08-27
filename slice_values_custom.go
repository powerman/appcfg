package appcfg

import (
	"net"
	"time"
)

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *DurationSlice) Value(err *error) (val []time.Duration) { //nolint:gocritic // ptrToRefParam.
	if v.Get() == nil {
		*err = &RequiredError{v}
		return val
	}
	return v.values
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *BoolSlice) Value(err *error) (val []bool) { //nolint:gocritic // ptrToRefParam.
	if v.Get() == nil {
		*err = &RequiredError{v}
		return val
	}
	return v.values
}

// IsBoolFlag implements extended flag.Value interface.
func (*BoolSlice) IsBoolFlag() bool {
	return true
}

// NewOneOfStringSlice returns OneOfStringSlice without value set.
func NewOneOfStringSlice(oneOf []string) OneOfStringSlice {
	return OneOfStringSlice{oneOf: oneOf}
}

// MustOneOfStringSlice returns OneOfStringSlice initialized with given value or panics.
func MustOneOfStringSlice(oneOf []string, ss ...string) OneOfStringSlice {
	if len(ss) == 0 {
		panic("require at least 1 arg")
	}
	v := NewOneOfStringSlice(oneOf)
	for _, s := range ss {
		err := v.Set(s)
		if err != nil {
			panic(err)
		}
	}
	v.completed = true
	return v
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Int64Slice) Value(err *error) (val []int64) { //nolint:gocritic // ptrToRefParam.
	if v.Get() == nil {
		*err = &RequiredError{v}
		return val
	}
	return v.values
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *UintSlice) Value(err *error) (val []uint) { //nolint:gocritic // ptrToRefParam.
	if v.Get() == nil {
		*err = &RequiredError{v}
		return val
	}
	return v.values
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Uint64Slice) Value(err *error) (val []uint64) { //nolint:gocritic // ptrToRefParam.
	if v.Get() == nil {
		*err = &RequiredError{v}
		return val
	}
	return v.values
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Float64Slice) Value(err *error) (val []float64) { //nolint:gocritic // ptrToRefParam.
	if v.Get() == nil {
		*err = &RequiredError{v}
		return val
	}
	return v.values
}

// NewIntBetweenSlice returns IntBetweenSlice without value set.
func NewIntBetweenSlice(min, max int) IntBetweenSlice {
	return IntBetweenSlice{min: min, max: max}
}

// MustIntBetweenSlice returns IntBetweenSlice initialized with given value or panics.
func MustIntBetweenSlice(min, max int, ss ...string) IntBetweenSlice {
	if len(ss) == 0 {
		panic("require at least 1 arg")
	}
	v := NewIntBetweenSlice(min, max)
	for _, s := range ss {
		err := v.Set(s)
		if err != nil {
			panic(err)
		}
	}
	v.completed = true
	return v
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *IPNetSlice) Value(err *error) (val []*net.IPNet) { //nolint:gocritic // ptrToRefParam.
	if v.Get() == nil {
		*err = &RequiredError{v}
		return val
	}
	return v.values
}

type HostPortTuple struct {
	Host string
	Port int
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *HostPortSlice) Value(err *error) []HostPortTuple { //nolint:gocritic // ptrToRefParam.
	if v.Get() == nil {
		*err = &RequiredError{v}
		return nil
	}
	return v.tuples
}
