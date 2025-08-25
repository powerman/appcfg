package appcfg

import (
	"errors"
	"reflect"
)

var (
	errEmptyOrWhite = errors.New("empty or contain only whitespaces")
	errNoHost       = errors.New("no host")
	errOverflows    = errors.New("value overflows")
	errNotOneOf     = errors.New("not one of")
	errNotBetween   = errors.New("not between")
)

const parseBits = 64

// Value provides a way to set value of any type from (one or several)
// string with validation. It's suitable for receiving configuration from
// hardcoded defaults, environment variables, command line flags, etc.
//
// Implements common interfaces for flags (flag.Value, flag.Getter,
// github.com/spf13/pflag.Value and github.com/urfave/cli.Generic), but
// require more strict semantics to distinguish between zero and unset
// values (Set resets value to unset on error, Get return nil on unset).
//
// Type containing multiple values should either replace or append value (or values if it can
// parse multiple values from a string) on Set depending on call sequence: replace if Set was
// called after Get, otherwise append.
//
// Types containing boolean value should also provide IsBoolFlag() method
// returning true.
type Value interface {
	String() string     // May be called on zero-valued receiver.
	Set(s string) error // Unset value on error.
	Get() any           // Return nil when not set (no default and no Set() or last Set() failed).
	Type() string       // Value type for help/usage.
}

//nolint:gochecknoglobals // By design.
var typValue = reflect.TypeOf(new(Value)).Elem()

func implementsValue(typ reflect.Type) bool {
	return typ.Implements(typValue) || reflect.PointerTo(typ).Implements(typValue)
}
