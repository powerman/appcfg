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
