//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "Duration=String,NotEmptyString,OneOfString,Endpoint,Port,ListenPort"

package appcfg

import (
	"fmt"
)

var _ Value = &Duration{}

// String implements flag.Value interface.
func (v *Duration) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *Duration) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *Duration) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *Duration) Type() string {
	return "Duration"
}
