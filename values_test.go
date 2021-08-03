package appcfg_test

import (
	"testing"

	"github.com/powerman/check"

	"github.com/powerman/appcfg"
)

func TestBool(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	var v appcfg.Bool
	t.Equal(v.Type(), "Bool")

	t.Equal(v.String(), "")
	t.Nil(v.Get())
	var err error
	t.False(v.Value(&err))
	t.Match(err, "required")

	t.Nil(v.Set("0"))

	t.Equal(v.String(), "false")
	t.NotNil(v.Get())
	err = nil
	t.False(v.Value(&err))
	t.Nil(err)

	v = appcfg.MustBool("1")

	t.Equal(v.String(), "true")
	t.NotNil(v.Get())
	err = nil
	t.True(v.Value(&err))
	t.Nil(err)

	t.PanicMatch(func() { v = appcfg.MustBool("") }, "invalid")
}

func TestInt(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	var v appcfg.Int
	t.Equal(v.Type(), "Int")

	t.Equal(v.String(), "")
	t.Nil(v.Get())
	var err error
	t.Zero(v.Value(&err))
	t.Match(err, "required")

	t.Nil(v.Set("0"))

	t.Equal(v.String(), "0")
	t.NotNil(v.Get())
	err = nil
	t.Zero(v.Value(&err))
	t.Nil(err)

	v = appcfg.MustInt("-0b111")

	t.Equal(v.String(), "-7")
	t.NotNil(v.Get())
	err = nil
	t.Equal(v.Value(&err), -7)
	t.Nil(err)

	t.PanicMatch(func() { v = appcfg.MustInt("") }, "invalid")
}

func TestUint(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	var v appcfg.Uint
	t.Equal(v.Type(), "Uint")

	t.Equal(v.String(), "")
	t.Nil(v.Get())
	var err error
	t.Zero(v.Value(&err))
	t.Match(err, "required")

	t.Nil(v.Set("0"))

	t.Equal(v.String(), "0")
	t.NotNil(v.Get())
	err = nil
	t.Zero(v.Value(&err))
	t.Nil(err)

	v = appcfg.MustUint("0b111")

	t.Equal(v.String(), "7")
	t.NotNil(v.Get())
	err = nil
	t.Equal(v.Value(&err), uint(7))
	t.Nil(err)

	t.PanicMatch(func() { v = appcfg.MustUint("") }, "invalid")
	t.PanicMatch(func() { v = appcfg.MustUint("-1") }, "invalid")
}

func TestIPNet(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	var v appcfg.IPNet
	t.Equal(v.Type(), "IPNet")

	t.Equal(v.String(), "")
	t.Nil(v.Get())
	var err error
	t.Zero(v.Value(&err))
	t.Match(err, "required")

	t.Nil(v.Set("0.0.0.0/0"))

	t.Equal(v.String(), "0.0.0.0/0")
	t.NotNil(v.Get())
	err = nil
	t.NotZero(v.Value(&err))
	t.Nil(err)

	v = appcfg.MustIPNet("192.168.2.1/24")

	t.Equal(v.String(), "192.168.2.0/24")
	t.NotNil(v.Get())
	err = nil
	t.Equal(v.Value(&err).String(), "192.168.2.0/24")
	t.Nil(err)

	t.PanicMatch(func() { v = appcfg.MustIPNet("") }, "invalid")
}

func TestHostPort(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	var v appcfg.HostPort
	t.Equal(v.Type(), "HostPort")

	t.Equal(v.String(), "")
	t.Nil(v.Get())
	var err error
	t.Zero(v.Value(&err))
	t.Match(err, "required")

	t.Nil(v.Set("0.0.0.0:0"))

	t.Equal(v.String(), "0.0.0.0:0")
	t.NotNil(v.Get())
	err = nil
	host, port := v.Value(&err)
	t.Equal(host, "0.0.0.0")
	t.Equal(port, 0)
	t.Nil(err)

	v = appcfg.MustHostPort("localhost:80")

	t.Equal(v.String(), "localhost:80")
	t.NotNil(v.Get())
	err = nil
	host, port = v.Value(&err)
	t.Equal(host, "localhost")
	t.Equal(port, 80)
	t.Nil(err)

	t.PanicMatch(func() { v = appcfg.MustHostPort("localhost") }, "missing port")
	t.PanicMatch(func() { v = appcfg.MustHostPort("localhost:") }, "port: .* parsing")
	t.PanicMatch(func() { v = appcfg.MustHostPort("localhost:http") }, "port: .* parsing")
	t.PanicMatch(func() { v = appcfg.MustHostPort(":80") }, "no host")
}
