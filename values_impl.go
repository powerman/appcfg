//go:generate -command genny sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" genny
//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "Duration=Bool,String,NotEmptyString,OneOfString,Endpoint,Int,Int64,Uint,Uint64,Float64,IntBetween,Port,ListenPort,IPNet"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

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
func (*Duration) Type() string {
	return "Duration"
}
