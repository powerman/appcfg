//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "Duration=Bool,String,NotEmptyString,Endpoint,Int,Int64,Uint,Uint64,Float64,Port,ListenPort"

package appcfg

// MustDuration returns Duration initialized with given value or panics.
func MustDuration(s string) Duration {
	var v Duration
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}
