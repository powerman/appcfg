package appcfg

import (
	"fmt"
	"math"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// DurationSlice can be set to comma-separated strings valid for time.ParseDuration().
type DurationSlice struct {
	values    []time.Duration
	completed bool
}

func (v *DurationSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []time.Duration{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		d, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		v.values = append(v.values, d)
	}
	return nil
}

// BoolSlice can be set to comma-separated strings
// 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
type BoolSlice struct {
	values    []bool
	completed bool
}

func (v *BoolSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []bool{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		v.values = append(v.values, b)
	}
	return nil
}

// StringArray can be set to any strings, even empty.
type StringArray struct {
	values    []string
	completed bool
}

func (v *StringArray) set(s string) error {
	v.values = append(v.values, s)
	return nil
}

// StringSlice can be set to any comma-separated strings, even empty.
type StringSlice struct {
	values    []string
	completed bool
}

func (v *StringSlice) set(ss string) error {
	v.values = append(v.values, strings.Split(ss, ",")...)
	return nil
}

// NotEmptyStringArray can be set to any strings which contains at least one
// non-whitespace symbol.
type NotEmptyStringArray struct {
	values    []string
	completed bool
}

func (v *NotEmptyStringArray) set(s string) error {
	if v.values == nil && s == "" {
		v.values = []string{}
		return nil
	}
	if strings.TrimSpace(s) == "" {
		return errEmptyOrWhite
	}
	v.values = append(v.values, s)
	return nil
}

// NotEmptyStringSlice can be set to any comma-separated strings which contains at least one
// non-whitespace symbol.
type NotEmptyStringSlice struct {
	values    []string
	completed bool
}

func (v *NotEmptyStringSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []string{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		if strings.TrimSpace(s) == "" {
			return errEmptyOrWhite
		}
		v.values = append(v.values, s)
	}
	return nil
}

// OneOfStringSlice can be set to predefined (using NewOneOfString or MustOneOfString)
// comma-separated strings.
type OneOfStringSlice struct {
	values    []string
	completed bool
	oneOf     []string
}

func (v *OneOfStringSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []string{}
		return nil
	}
SET:
	for _, s := range strings.Split(ss, ",") {
		for _, item := range v.oneOf {
			if s == item {
				v.values = append(v.values, s)
				continue SET
			}
		}
		return fmt.Errorf("%w %q", errNotOneOf, v.oneOf)
	}
	return nil
}

// EndpointSlice can be set to valid urls with hostname. Also it'll trim
// all / symbols at end, to make it easier to append paths to endpoint.
type EndpointSlice struct {
	values    []string
	completed bool
}

func (v *EndpointSlice) set(s string) error {
	if v.values == nil && s == "" {
		v.values = []string{}
		return nil
	}
	s = strings.TrimRight(s, "/")
	p, err := url.Parse(s)
	if err != nil {
		return err
	} else if p.Host == "" {
		return errNoHost
	}
	v.values = append(v.values, s)
	return nil
}

// IntSlice can be set to comma-separated integer values.
// It's allowed to use 0b, 0o and 0x prefixes, and also underscores.
type IntSlice struct {
	values    []int
	completed bool
}

func (v *IntSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []int{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		i64, err := strconv.ParseInt(s, 0, strconv.IntSize)
		if err != nil {
			return err
		}
		i := int(i64)
		if int64(i) != i64 {
			return fmt.Errorf("%w int: %s", errOverflows, s)
		}
		v.values = append(v.values, i)
	}
	return nil
}

// Int64Slice can be set to comma-separated 64-bit integer values.
// It's allowed to use 0b, 0o and 0x prefixes, and also underscores.
type Int64Slice struct {
	values    []int64
	completed bool
}

func (v *Int64Slice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []int64{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		i, err := strconv.ParseInt(s, 0, parseBits)
		if err != nil {
			return err
		}
		v.values = append(v.values, i)
	}
	return nil
}

// UintSlice can be set to comma-separated unsigned integer values.
// It's allowed to use 0b, 0o and 0x prefixes, and also underscores.
type UintSlice struct {
	values    []uint
	completed bool
}

func (v *UintSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []uint{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		i64, err := strconv.ParseUint(s, 0, strconv.IntSize)
		if err != nil {
			return err
		}
		i := uint(i64)
		if uint64(i) != i64 {
			return fmt.Errorf("%w unsigned int: %s", errOverflows, s)
		}
		v.values = append(v.values, i)
	}
	return nil
}

// Uint64Slice can be set to comma-separated unsigned 64-bit integer values.
// It's allowed to use 0b, 0o and 0x prefixes, and also underscores.
type Uint64Slice struct {
	values    []uint64
	completed bool
}

func (v *Uint64Slice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []uint64{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		i, err := strconv.ParseUint(s, 0, parseBits)
		if err != nil {
			return err
		}
		v.values = append(v.values, i)
	}
	return nil
}

// Float64Slice can be set to comma-separated 64-bit floating-point numbers.
type Float64Slice struct {
	values    []float64
	completed bool
}

func (v *Float64Slice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []float64{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		i, err := strconv.ParseFloat(s, parseBits)
		if err != nil {
			return err
		}
		v.values = append(v.values, i)
	}
	return nil
}

// IntBetweenSlice can be set to comma-separated integer values between given (using
// NewIntBetween or MustIntBetween) min/max values (inclusive).
type IntBetweenSlice struct {
	values    []int
	completed bool
	min, max  int
}

func (v *IntBetweenSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []int{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		} else if v.min > i || i > v.max {
			return fmt.Errorf("%w %d and %d", errNotBetween, v.min, v.max)
		}
		v.values = append(v.values, i)
	}
	return nil
}

// PortSlice can be set to comma-separated integer values between 1 and 65535.
type PortSlice struct {
	values    []int
	completed bool
}

func (v *PortSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []int{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		} else if 0 >= i || i > math.MaxUint16 {
			return fmt.Errorf("%w 1 and %d", errNotBetween, math.MaxUint16)
		}
		v.values = append(v.values, i)
	}
	return nil
}

// ListenPortSlice can be set to comma-separated integer values between 0 and 65535.
type ListenPortSlice struct {
	values    []int
	completed bool
}

func (v *ListenPortSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []int{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		} else if 0 > i || i > math.MaxUint16 {
			return fmt.Errorf("%w 0 and %d", errNotBetween, math.MaxUint16)
		}
		v.values = append(v.values, i)
	}
	return nil
}

// IPNetSlice can be set to comma-separated CIDR address.
type IPNetSlice struct {
	values    []*net.IPNet
	completed bool
}

func (v *IPNetSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []*net.IPNet{}
		return nil
	}
	for _, s := range strings.Split(ss, ",") {
		_, ipNet, err := net.ParseCIDR(s)
		if err != nil {
			return err
		}
		v.values = append(v.values, ipNet)
	}
	return nil
}

// HostPortSlice can be set to comma-separated CIDR addresses.
type HostPortSlice struct {
	values    []string
	tuples    []HostPortTuple
	completed bool
}

func (v *HostPortSlice) set(ss string) error {
	if v.values == nil && ss == "" {
		v.values = []string{}
		v.tuples = []HostPortTuple{}
		return nil
	}
	if v.values == nil {
		v.tuples = nil
	}
	for _, s := range strings.Split(ss, ",") {
		var tuple HostPortTuple
		host, port, err := net.SplitHostPort(s)
		if err != nil {
			return err
		}
		if host == "" {
			return errNoHost
		}
		tuple.Host = host
		tuple.Port, err = strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("port: %w", err)
		}
		v.values = append(v.values, s)
		v.tuples = append(v.tuples, tuple)
	}
	return nil
}
