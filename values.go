package appcfg

import (
	"errors"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"
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
		return errors.New("empty or contain only whitespaces")
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
	return fmt.Errorf("not one of %q", v.oneOf)
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
		return errors.New("no host")
	}
	v.value = &s
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
		return fmt.Errorf("not between %d and %d", v.min, v.max)
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
		return fmt.Errorf("not between 1 and %d", math.MaxUint16)
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
		return fmt.Errorf("not between 0 and %d", math.MaxUint16)
	}
	v.value = &i
	return nil
}
