package main

import (
	"os"
	"testing"
	"time"

	"github.com/powerman/check"
)

func Test(tt *testing.T) {
	t := check.T(tt)
	os.Clearenv()
	want := &config{
		Host:      "example.com",
		Port:      443,
		BindPorts: []int{80, 443},
		Timeout:   30 * time.Second,
		Retries:   1,
	}

	t.Run("required", func(tt *testing.T) {
		t := check.T(tt)
		errMatch(t, "", `^Retries .* required`)
		t.Setenv("EXAMPLE_RETRIES", "  1  ")
		errMatch(t, "", `^Host .* required`)
	})
	t.Run("default", func(tt *testing.T) {
		t := check.T(tt)
		t.Setenv("EXAMPLE_RETRIES", "  1  ")
		t.Setenv("EXAMPLE_HOST", " example.com")
		cfg, err := testGetCfg()
		t.Nil(err)
		t.DeepEqual(cfg, want)
	})
	t.Run("constraint", func(tt *testing.T) {
		t := check.T(tt)
		t.Setenv("EXAMPLE_RETRIES", "  1  ")
		t.Setenv("EXAMPLE_HOST", " example.com")
		errMatchEnv(t, "EXAMPLE_HOST", "", `^Host .* empty`)
		errMatchEnv(t, "EXAMPLE_PORT", "0", `^Port .* between`)
		errMatchEnv(t, "EXAMPLE_BIND_PORTS", "1,-1,2", `^BindPorts .* between`)
		errMatchEnv(t, "EXAMPLE_RETRIES", "5", `^Retries .* between`)
		errMatch(t, "-host=", `^Host .* required`)
		errMatch(t, "-port=", `^Port .* required`)
		errMatch(t, "-timeout=x", `^Timeout .* required`)
	})
	t.Run("env", func(tt *testing.T) {
		t := check.T(tt)
		t.Setenv("EXAMPLE_HOST", "example2.com")
		t.Setenv("EXAMPLE_PORT", "2443")
		t.Setenv("EXAMPLE_BIND_PORTS", "2080,2443")
		t.Setenv("EXAMPLE_RETRIES", "2")
		cfg, err := testGetCfg()
		t.Nil(err)
		want.Host = "example2.com"
		want.Port = 2443
		want.BindPorts = []int{2080, 2443}
		want.Retries = 2
		t.DeepEqual(cfg, want)
	})
	t.Run("flag", func(tt *testing.T) {
		t := check.T(tt)
		t.Setenv("EXAMPLE_HOST", "example2.com")
		t.Setenv("EXAMPLE_PORT", "2443")
		t.Setenv("EXAMPLE_BIND_PORTS", "2080,2443")
		t.Setenv("EXAMPLE_RETRIES", "2")
		cfg, err := testGetCfg(
			"-host=example3.com",
			"-port=3443",
			"-bind.ports=3080,3443",
			"-timeout=3s",
		)
		t.Nil(err)
		want.Host = "example3.com"
		want.Port = 3443
		want.BindPorts = []int{3080, 3443}
		want.Timeout = 3 * time.Second
		t.DeepEqual(cfg, want)
	})
}
