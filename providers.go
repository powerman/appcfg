package appcfg

import (
	"fmt"
	"os"
	"strings"
)

// Tags provide access to tags attached to some Value.
type Tags interface {
	// Get returns the value associated with key. If there is no such
	// key, Get returns the empty string.
	Get(key string) string
	// Lookup returns the value associated with key. If the key is
	// present the value (which may be empty) is returned. Otherwise
	// the returned value will be the empty string. The ok return
	// value reports whether the key is present.
	Lookup(key string) (value string, ok bool)
}

// Provider tries to find and set value based on it name and/or tags.
type Provider interface {
	Provide(value Value, name string, tags Tags) (provided bool, err error)
}

// FromEnv implements Provider using value from environment variable with
// name defined by tag "env" with optional prefix.
type FromEnv struct {
	prefix    string
	trimSpace bool
}

// NewFromEnv creates new FromEnv with optional prefix.
func NewFromEnv(prefix string, opts ...FromEnvOption) *FromEnv {
	f := &FromEnv{
		prefix: prefix,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

// Provide implements Provider.
func (f *FromEnv) Provide(value Value, _ string, tags Tags) (bool, error) {
	name := tags.Get("env")
	if name == "" {
		return false, nil
	}
	name = f.prefix + name
	if s, ok := os.LookupEnv(name); ok {
		if f.trimSpace {
			s = strings.TrimSpace(s)
		}
		err := value.Set(s)
		if err != nil {
			err = fmt.Errorf("$%s=%q: %w", name, s, err)
		}
		return true, err
	}
	return false, nil
}

// FromEnvOption is an option for NewFromEnv.
type FromEnvOption func(*FromEnv)

// FromEnvTrimSpace removes from environment variables value all leading
// and trailing white space, as defined by Unicode.
func FromEnvTrimSpace() FromEnvOption {
	return func(f *FromEnv) { f.trimSpace = true }
}
