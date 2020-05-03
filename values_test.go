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
