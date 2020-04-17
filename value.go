package appcfg

import "reflect"

// Value provides a way to set value of any type from (one or several)
// string with validation. It's suitable for receiving configuration from
// hardcoded defaults, environment variables, command line flags, etc.
//
// Implements common interfaces for flags (flag.Value, flag.Getter,
// github.com/spf13/pflag.Value and github.com/urfave/cli.Generic), but
// require more strict semantics to distinguish between zero and unset
// values (Set resets value to unset on error, Get return nil on unset).
//
// Types containing multiple values should also implement SliceValue and
// use it Replace() method to avoid joining values from different sources
// (like environment variable and flag).
//
// Types containing boolean value should also provide IsBoolFlag() method
// returning true.
type Value interface {
	String() string   // May be called on zero-valued receiver.
	Set(string) error // Unset value on error.
	Get() interface{} // Return nil when not set (no default and no Set() or last Set() failed).
	Type() string     // Value type for help/usage.
}

// SliceValue is same as github.com/spf13/pflag.SliceValue.
type SliceValue interface {
	// Append adds the specified value to the end of the flag value list.
	Append(string) error
	// Replace will fully overwrite any data currently in the flag value list.
	Replace([]string) error
	// GetSlice returns the flag value list as an array of strings.
	GetSlice() []string
}

//nolint:gochecknoglobals // By design.
var typValue = reflect.TypeOf(new(Value)).Elem()

func implementsValue(typ reflect.Type) bool {
	return typ.Implements(typValue) || reflect.PtrTo(typ).Implements(typValue)
}
