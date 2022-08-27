package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/powerman/check"
)

func TestMain(m *testing.M) { check.TestMain(m) }

var origExtCfg = extCfg

func testGetCfg(flags ...string) (interface{}, error) {
	extCfg = origExtCfg
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)

	err := initExtCfg(fs)
	if err != nil {
		return nil, err
	}
	fs.Parse(flags)
	return getCfg()
}

func errMatch(t *check.C, flags string, match string) {
	t.Helper()
	cfg, err := testGetCfg(strings.Fields(flags)...)
	t.Match(err, match)
	t.Nil(cfg)
}

func errMatchEnv(t *check.C, name, val, match string) {
	t.Helper()
	old, ok := os.LookupEnv(name)
	if ok {
		defer os.Setenv(name, old)
	} else {
		defer os.Unsetenv(name)
	}
	t.Nil(os.Setenv(name, val))
	errMatch(t, "", match)
}
