package appcfg

import (
	"net"
	"time"
)

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Duration) Value(err *error) (val time.Duration) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Bool) Value(err *error) (val bool) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}

// IsBoolFlag implements extended flag.Value interface.
func (*Bool) IsBoolFlag() bool {
	return true
}

// NewOneOfString returns OneOfString without value set.
func NewOneOfString(oneOf []string) OneOfString {
	return OneOfString{oneOf: oneOf}
}

// MustOneOfString returns OneOfString initialized with given value or panics.
func MustOneOfString(s string, oneOf []string) OneOfString {
	v := NewOneOfString(oneOf)
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Int64) Value(err *error) (val int64) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Uint) Value(err *error) (val uint) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Uint64) Value(err *error) (val uint64) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Float64) Value(err *error) (val float64) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}

// NewIntBetween returns IntBetween without value set.
func NewIntBetween(minVal, maxVal int) IntBetween {
	return IntBetween{min: minVal, max: maxVal}
}

// MustIntBetween returns IntBetween initialized with given value or panics.
func MustIntBetween(s string, minVal, maxVal int) IntBetween {
	v := NewIntBetween(minVal, maxVal)
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *IPNet) Value(err *error) (val *net.IPNet) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *HostPort) Value(err *error) (host string, port int) { //nolint:gocritic // ptrToRefParam.
	if v.value == nil {
		*err = &RequiredError{v}
		return host, port
	}
	return v.host, v.port
}
