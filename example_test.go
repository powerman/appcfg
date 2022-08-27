package appcfg_test

import (
	"flag"
	"log"
	"time"

	"github.com/powerman/appcfg"
)

// Intermediate config, used to gather and validate values from different
// external sources (environment variables, flags, etc.).
var extCfg = struct { // Type defines constraint types and providers for each exported field.
	Host      appcfg.NotEmptyString  `env:"HOST"`
	Port      appcfg.Port            `env:"PORT"`
	BindPorts appcfg.ListenPortSlice `env:"BIND_PORTS"`
	Timeout   appcfg.Duration
	Retries   appcfg.IntBetween `env:"RETRIES"`
	fs        *flag.FlagSet     // Needed just to avoid using globals and ease testing.
}{ // Values may define defaults for some fields and must setup some types.
	Port:      appcfg.MustPort("443"), // Set default.
	BindPorts: appcfg.MustListenPortSlice("80", "443"),
	Timeout:   appcfg.MustDuration("30s"), // Set default.
	Retries:   appcfg.NewIntBetween(1, 3), // Configure value constraints, no default.
}

// initExtCfg should be called before calling fs.Parse() - it'll gather
// configuration values from all sources except flags, and will setup
// flags on given fs.
func initExtCfg(fs *flag.FlagSet) error {
	const envPrefix = "EXAMPLE_"
	fromEnv := appcfg.NewFromEnv(envPrefix, appcfg.FromEnvTrimSpace())
	err := appcfg.ProvideStruct(&extCfg, fromEnv) // Set appCfg fields from environment.

	extCfg.fs = fs
	appcfg.AddFlag(fs, &extCfg.Host, "host", "host to connect")
	appcfg.AddFlag(fs, &extCfg.Port, "port", "port to connect")
	appcfg.AddFlag(fs, &extCfg.BindPorts, "bind.ports", "ports to bind")
	appcfg.AddFlag(fs, &extCfg.Timeout, "timeout", "connect timeout")

	return err
}

// config contains validated values in a way convenient for your app.
type config struct {
	Host      string        // Not "",  set by $EXAMPLE_HOST or -host, no default.
	Port      int           // 1…65535, set by $EXAMPLE_PORT or -port, default 443.
	BindPorts []int         // 1…65535, set by $EXAMPLE_BIND_PORTS or -bind.ports, default 80,443.
	Timeout   time.Duration // Any,     set by -timeout,               default 30s.
	Retries   int           // 1…3,     set by $EXAMPLE_RETRIES,       no default.
}

// getCfg checks is all required config values was provided and converts
// them into structure convenient for your app.
func getCfg() (cfg *config, err error) {
	cfg = &config{
		Host:      extCfg.Host.Value(&err), // Value set err if field was not set.
		Port:      extCfg.Port.Value(&err),
		BindPorts: extCfg.BindPorts.Value(&err),
		Timeout:   extCfg.Timeout.Value(&err),
		Retries:   extCfg.Retries.Value(&err),
	}
	if err != nil {
		return nil, appcfg.WrapErr(err, extCfg.fs, &extCfg)
	}
	return cfg, nil
}

func Example() {
	err := initExtCfg(flag.CommandLine)
	if err != nil {
		log.Print(err)
		return
	}

	flag.Parse()

	cfg, err := getCfg()
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("cfg: %#v", cfg)
}
