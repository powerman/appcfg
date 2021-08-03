//go:generate -command genny sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" genny
//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "Duration=Bool,String,NotEmptyString,Endpoint,Int,Int64,Uint,Uint64,Float64,Port,ListenPort,IPNet,HostPort"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

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
