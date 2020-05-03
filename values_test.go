package appcfg

import (
	"testing"

	"github.com/powerman/check"
)

func TestBool(tt *testing.T) {
	t := check.T(tt)

	var v Bool
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

	v = MustBool("1")

	t.Equal(v.String(), "true")
	t.NotNil(v.Get())
	err = nil
	t.True(v.Value(&err))
	t.Nil(err)

	t.PanicMatch(func() { v = MustBool("") }, "invalid")
}

func TestInt(tt *testing.T) {
	t := check.T(tt)

	var v Int
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

	v = MustInt("-0b111")

	t.Equal(v.String(), "-7")
	t.NotNil(v.Get())
	err = nil
	t.Equal(v.Value(&err), -7)
	t.Nil(err)

	t.PanicMatch(func() { v = MustInt("") }, "invalid")
}

func TestUint(tt *testing.T) {
	t := check.T(tt)

	var v Uint
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

	v = MustUint("0b111")

	t.Equal(v.String(), "7")
	t.NotNil(v.Get())
	err = nil
	t.Equal(v.Value(&err), uint(7))
	t.Nil(err)

	t.PanicMatch(func() { v = MustUint("") }, "invalid")
	t.PanicMatch(func() { v = MustUint("-1") }, "invalid")
}
