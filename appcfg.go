// Package appcfg provides helpers to get valid application configuration
// values from different sources (flags, env, config files, services like
// consul, etc.).
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
// must be a ref to struct with fields of Value type having struct tag
// with tags for given providers. Current values in cfg, if any, will be
// used as defaults.
//
// Providers will be called for each cfg field, in order, with next
// provider will be called only if previous providers won't provide a
// value for a current field.
//
// It is recommended to apply "from flags" provider(s) using separate call
// after all other providers will be applied - this way flag provider will
// be able to show values set by other providers as flag defaults in usage
// message.
//
// Returns error if any provider will try to set invalid value.
func ProvideStruct(cfg interface{}, providers ...Provider) error {
	var lastErr error
	forStruct(cfg, func(value Value, name string, tags Tags) {
		for _, provider := range providers {
			err, ok := provider.Provide(value, name, tags)
			if err != nil {
				lastErr = fmt.Errorf("cfg.%s %#q: %w", name, tags, err)
				break
			}
			if ok {
				break
			}
		}
	})
	return lastErr
}

type RequiredError struct{ Value }

func (*RequiredError) Error() string { return "value required" }

func WrapErr(err error, fs *flag.FlagSet, cfgs ...interface{}) error {
	if reqErr := new(RequiredError); errors.As(err, &reqErr) {
		var flagName string
		if fs != nil {
			fs.VisitAll(func(f *flag.Flag) {
				if f.Value == reqErr.Value {
					flagName = f.Name
				}
			})
		}
		return wrapErr(reqErr, "-"+flagName, cfgs...)
	}
	return err
}

func WrapPErr(err error, fs *pflag.FlagSet, cfgs ...interface{}) error {
	if reqErr := new(RequiredError); errors.As(err, &reqErr) {
		var flagName string
		if fs != nil {
			fs.VisitAll(func(f *pflag.Flag) {
				if f.Value == reqErr.Value {
					flagName = f.Name
				}
			})
		}
		return wrapErr(reqErr, "--"+flagName, cfgs...)
	}
	return err
}

func wrapErr(reqErr *RequiredError, flagName string, cfgs ...interface{}) error {
	var lastErr error
	for _, cfg := range cfgs {
		forStruct(cfg, func(value Value, name string, tags Tags) {
			if value == reqErr.Value {
				s := fmt.Sprintf("%s %s", flagName, tags)
				s = strings.TrimSpace(strings.Join(strings.Fields(s), " "))
				if s != "" {
					name = fmt.Sprintf("%s (%s)", name, s)
				}
				lastErr = fmt.Errorf("%s: %w", name, reqErr)
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
		if !implementsValue(f.Type) {
			panic(fmt.Sprintf("cfg.%s: must implements Value", f.Name))
		}
		f.Tag = reflect.StructTag(strings.ReplaceAll(string(f.Tag), "\n", " "))
		f.Tag = reflect.StructTag(strings.ReplaceAll(string(f.Tag), "\t", " "))
		value := val.Elem().FieldByName(f.Name).Addr().Interface().(Value)
		handle(value, f.Name, f.Tag)
	}
}
