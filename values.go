package appcfg

import (
	"errors"
	"fmt"
	"math"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	errEmptyOrWhite = errors.New("empty or contain only whitespaces")
	errNoHost       = errors.New("no host")
	errOverflows    = errors.New("value overflows")
	errNotOneOf     = errors.New("not one of")
	errNotBetween   = errors.New("not between")
)

// Duration can be Set only to string valid for time.ParseDuration().
type Duration struct{ value *time.Duration }

func (v *Duration) set(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	v.value = &d
	return nil
}

// Bool can be set to 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
type Bool struct{ value *bool }

func (v *Bool) set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	v.value = &b
	return nil
}

// String can be set to any string, even empty.
type String struct{ value *string }

func (v *String) set(s string) error {
	v.value = &s
	return nil
}

// NotEmptyString can be set to any string which contains at least one
// non-whitespace symbol.
type NotEmptyString struct{ value *string }

func (v *NotEmptyString) set(s string) error {
	if strings.TrimSpace(s) == "" {
		return errEmptyOrWhite
	}
	v.value = &s
	return nil
}

// OneOfString can be set to any of predefined (using NewOneOfString or
// MustOneOfString) values.
type OneOfString struct {
	value *string
	oneOf []string
}

func (v *OneOfString) set(s string) error {
	for _, item := range v.oneOf {
		if s == item {
			v.value = &s
			return nil
		}
	}
	return fmt.Errorf("%w %q", errNotOneOf, v.oneOf)
}

// Endpoint can be set only to valid url with hostname. Also it'll trim
// all / symbols at end, to make it easier to append paths to endpoint.
type Endpoint struct{ value *string }

func (v *Endpoint) set(s string) error {
	s = strings.TrimRight(s, "/")
	p, err := url.Parse(s)
	if err != nil {
		return err
	} else if p.Host == "" {
		return errNoHost
	}
	v.value = &s
	return nil
}

// Int can be set to integer value.
// It's allowed to use 0b, 0o and 0x prefixes, and also underscores.
type Int struct{ value *int }

func (v *Int) set(s string) error {
	i64, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	i := int(i64)
	if int64(i) != i64 {
		return fmt.Errorf("%w int: %s", errOverflows, s)
	}
	v.value = &i
	return nil
}

// Int64 can be set to 64-bit integer value.
// It's allowed to use 0b, 0o and 0x prefixes, and also underscores.
type Int64 struct{ value *int64 }

func (v *Int64) set(s string) error {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}
	v.value = &i
	return nil
}

// Uint can be set to unsigned integer value.
// It's allowed to use 0b, 0o and 0x prefixes, and also underscores.
type Uint struct{ value *uint }

func (v *Uint) set(s string) error {
	i64, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	i := uint(i64)
	if uint64(i) != i64 {
		return fmt.Errorf("%w unsigned int: %s", errOverflows, s)
	}
	v.value = &i
	return nil
}

// Uint64 can be set to unsigned 64-bit integer value.
// It's allowed to use 0b, 0o and 0x prefixes, and also underscores.
type Uint64 struct{ value *uint64 }

func (v *Uint64) set(s string) error {
	i, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return err
	}
	v.value = &i
	return nil
}

// Float64 can be set to 64-bit floating-point number.
type Float64 struct{ value *float64 }

func (v *Float64) set(s string) error {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	v.value = &i
	return nil
}

// IntBetween can be set to integer value between given (using
// NewIntBetween or MustIntBetween) min/max values (inclusive).
type IntBetween struct {
	value    *int
	min, max int
}

func (v *IntBetween) set(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	} else if !(v.min <= i && i <= v.max) {
		return fmt.Errorf("%w %d and %d", errNotBetween, v.min, v.max)
	}
	v.value = &i
	return nil
}

// Port can be set to integer value between 1 and 65535.
type Port struct{ value *int }

func (v *Port) set(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	} else if !(0 < i && i <= math.MaxUint16) {
		return fmt.Errorf("%w 1 and %d", errNotBetween, math.MaxUint16)
	}
	v.value = &i
	return nil
}

// ListenPort can be set to integer value between 0 and 65535.
type ListenPort struct{ value *int }

func (v *ListenPort) set(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	} else if !(0 <= i && i <= math.MaxUint16) {
		return fmt.Errorf("%w 0 and %d", errNotBetween, math.MaxUint16)
	}
	v.value = &i
	return nil
}

// IPNet can be set to CIDR address.
type IPNet struct{ value **net.IPNet }

func (v *IPNet) set(s string) error {
	_, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return err
	}
	v.value = &ipNet
	return nil
}
