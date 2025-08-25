//go:generate mise exec -- genny -in=$GOFILE -out=gen.$GOFILE gen "DurationSlice=BoolSlice,StringArray,StringSlice,NotEmptyStringArray,NotEmptyStringSlice,OneOfStringSlice,EndpointSlice,IntSlice,Int64Slice,UintSlice,Uint64Slice,Float64Slice,IntBetweenSlice,PortSlice,ListenPortSlice,IPNetSlice,HostPortSlice"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

package appcfg

import (
	"fmt"
)

var _ Value = &DurationSlice{}

// String implements flag.Value interface.
func (v *DurationSlice) String() string {
	if v == nil || v.values == nil {
		return ""
	}
	return fmt.Sprint(v.values)
}

// Set implements flag.Value interface.
func (v *DurationSlice) Set(s string) error {
	if v.completed {
		v.completed = false
		v.values = nil
	}
	err := v.set(s)
	if err != nil {
		v.values = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *DurationSlice) Get() any {
	if v.values == nil {
		return nil
	}
	v.completed = true
	return v.values
}

// Type implements pflag.Value interface.
func (*DurationSlice) Type() string {
	return "DurationSlice"
}
