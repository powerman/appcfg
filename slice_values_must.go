//go:generate mise exec -- genny -in=$GOFILE -out=gen.$GOFILE gen "DurationSlice=BoolSlice,StringArray,StringSlice,NotEmptyStringArray,NotEmptyStringSlice,EndpointSlice,IntSlice,Int64Slice,UintSlice,Uint64Slice,Float64Slice,PortSlice,ListenPortSlice,IPNetSlice,HostPortSlice"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

package appcfg

// MustDurationSlice returns DurationSlice initialized with given values or panics.
func MustDurationSlice(ss ...string) DurationSlice {
	if len(ss) == 0 {
		panic("require at least 1 arg")
	}
	var v DurationSlice
	for _, s := range ss {
		err := v.Set(s)
		if err != nil {
			panic(err)
		}
	}
	v.completed = true
	return v
}
