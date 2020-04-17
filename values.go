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

type Duration struct{ value *time.Duration }

func (v *Duration) set(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	v.value = &d
	return nil
}

// Value is like Get except it returns zero value and set *err to
// RequiredError if unset.
func (v *Duration) Value(err *error) (val time.Duration) {
	if v.value == nil {
		*err = &RequiredError{v}
		return val
	}
	return *v.value
}

type String struct{ value *string }

func (v *String) set(s string) error {
	v.value = &s
	return nil
}

type NotEmptyString struct{ value *string }

func (v *NotEmptyString) set(s string) error {
	if strings.TrimSpace(s) == "" {
		return errors.New("empty or contain only whitespaces")
	}
	v.value = &s
	return nil
}

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

// NewOneOfString returns OneOfString without value set.
func NewOneOfString(oneOf []string) OneOfString {
	return OneOfString{oneOf: oneOf}
}

// MustOneOfString returns OneOfString initialized with given value or panics.
func MustOneOfString(s string, oneOf []string) OneOfString {
	var v = OneOfString{oneOf: oneOf}
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}

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
