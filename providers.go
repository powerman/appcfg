package appcfg

import (
	"fmt"
	"os"
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
	prefix string
}

// NewFromEnv creates new FromEnv with optional prefix.
func NewFromEnv(prefix string) *FromEnv {
	return &FromEnv{
		prefix: prefix,
	}
}

// Provide implements Provider.
func (f *FromEnv) Provide(value Value, _ string, tags Tags) (bool, error) {
	name := tags.Get("env")
	if name == "" {
		return false, nil
	}
	name = f.prefix + name
	if s, ok := os.LookupEnv(name); ok {
		err := value.Set(s)
		if err != nil {
			err = fmt.Errorf("$%s=%q: %w", name, s, err)
		}
		return true, err
	}
	return false, nil
}
