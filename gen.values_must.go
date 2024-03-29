// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny


package appcfg

// MustBool returns Bool initialized with given value or panics.
func MustBool(s string) Bool {
	var v Bool
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustString returns String initialized with given value or panics.
func MustString(s string) String {
	var v String
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustNotEmptyString returns NotEmptyString initialized with given value or panics.
func MustNotEmptyString(s string) NotEmptyString {
	var v NotEmptyString
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustEndpoint returns Endpoint initialized with given value or panics.
func MustEndpoint(s string) Endpoint {
	var v Endpoint
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustInt returns Int initialized with given value or panics.
func MustInt(s string) Int {
	var v Int
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustInt64 returns Int64 initialized with given value or panics.
func MustInt64(s string) Int64 {
	var v Int64
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustUint returns Uint initialized with given value or panics.
func MustUint(s string) Uint {
	var v Uint
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustUint64 returns Uint64 initialized with given value or panics.
func MustUint64(s string) Uint64 {
	var v Uint64
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustFloat64 returns Float64 initialized with given value or panics.
func MustFloat64(s string) Float64 {
	var v Float64
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustPort returns Port initialized with given value or panics.
func MustPort(s string) Port {
	var v Port
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustListenPort returns ListenPort initialized with given value or panics.
func MustListenPort(s string) ListenPort {
	var v ListenPort
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustIPNet returns IPNet initialized with given value or panics.
func MustIPNet(s string) IPNet {
	var v IPNet
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}


// MustHostPort returns HostPort initialized with given value or panics.
func MustHostPort(s string) HostPort {
	var v HostPort
	err := v.Set(s)
	if err != nil {
		panic(err)
	}
	return v
}
