package appcfg

import (
	"errors"
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/pflag"
)

// ProvideStruct updates cfg using values from given providers. Given cfg
// must be a ref to struct with all exported fields having Value type and
// struct tag with tags for given providers. Current values in cfg, if
// any, will be used as defaults.
//
// Providers will be called for each exported field in cfg, in order, with
// next provider will be called only if previous providers won't provide a
// value for a current field.
//
// It is recommended to add cfg fields to FlagSet after all other
// providers will be applied - this way usage message on -h flag will be
// able to show values set by other providers as flag defaults.
//
// Returns error if any provider will try to set invalid value.
func ProvideStruct(cfg interface{}, providers ...Provider) error {
	var lastErr error
	forStruct(cfg, func(value Value, name string, tags Tags) {
		for _, provider := range providers {
			ok, err := provider.Provide(value, name, tags)
			if err != nil {
				lastErr = fmt.Errorf("%s: %w", field(name, tags), err)
				break
			}
			if ok {
				break
			}
		}
	})
	return lastErr
}

// RequiredError is returned from Value(&err) methods if value wasn't set.
type RequiredError struct{ Value }

// Error implements error interface.
func (*RequiredError) Error() string { return "value required" }

// WrapErr adds more details about err.Value (if err is a RequiredError)
// by looking for related flag name and field name/tags in given fs and
// cfgs, otherwise returns err as is.
func WrapErr(err error, fs *flag.FlagSet, cfgs ...interface{}) error {
	if reqErr := new(RequiredError); errors.As(err, &reqErr) {
		var flagName string
		if fs != nil {
			fs.VisitAll(func(f *flag.Flag) {
				if f.Value == reqErr.Value {
					flagName = "-" + f.Name
				}
			})
		}
		return wrapErr(reqErr, flagName, cfgs...)
	}
	return err
}

// WrapPErr adds more details about err.Value (if err is a RequiredError)
// by looking for related flag name and field name/tags in given fs and
// cfgs, otherwise returns err as is.
func WrapPErr(err error, fs *pflag.FlagSet, cfgs ...interface{}) error {
	if reqErr := new(RequiredError); errors.As(err, &reqErr) {
		var flagName string
		if fs != nil {
			fs.VisitAll(func(f *pflag.Flag) {
				if f.Value == reqErr.Value {
					flagName = "--" + f.Name
				}
			})
		}
		return wrapErr(reqErr, flagName, cfgs...)
	}
	return err
}

func wrapErr(reqErr *RequiredError, flagName string, cfgs ...interface{}) error {
	var lastErr error
	for _, cfg := range cfgs {
		forStruct(cfg, func(value Value, name string, tags Tags) {
			if value == reqErr.Value {
				lastErr = fmt.Errorf("%s: %w", field(name, flagName, tags), reqErr)
			}
		})
	}
	if lastErr == nil {
		panic("required value not found in cfgs")
	}
	return lastErr
}

func forStruct(cfg interface{}, handle func(Value, string, Tags)) {
	val := reflect.ValueOf(cfg)
	typ := val.Type()
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		panic("cfg: must be a ptr to struct")
	}
	typ = typ.Elem()

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.PkgPath != "" {
			continue
		}
		if !implementsValue(f.Type) {
			panic(fmt.Sprintf("cfg.%s: must implements Value", f.Name))
		}
		f.Tag = reflect.StructTag(strings.ReplaceAll(string(f.Tag), "\n", " "))
		f.Tag = reflect.StructTag(strings.ReplaceAll(string(f.Tag), "\t", " "))
		value := val.Elem().FieldByName(f.Name).Addr().Interface().(Value)
		handle(value, f.Name, f.Tag)
	}
}

func field(name string, sources ...interface{}) string {
	s := strings.TrimSpace(strings.Join(strings.Fields(fmt.Sprintln(sources...)), " "))
	if s != "" {
		return fmt.Sprintf("%s (%s)", name, s)
	}
	return name
}

// AddFlag defines a flag with the specified name and usage string.
// Calling it again with same fs, value and name will have no effect.
func AddFlag(fs *flag.FlagSet, value flag.Value, name string, usage string) {
	if f := fs.Lookup(name); f != nil && f.Value == value {
		return
	}
	fs.Var(value, name, usage)
}

// AddPFlag defines a flag with the specified name and usage string.
// Calling it again with same fs, value and name will have no effect.
func AddPFlag(fs *pflag.FlagSet, value pflag.Value, name string, usage string) {
	if f := fs.Lookup(name); f != nil && f.Value == value {
		return
	}
	fs.Var(value, name, usage)
}
