// Package appcfg helps implement Clean Architecture adapter for
// application configuration.
//
// Application configuration values may come from different sources
// (defaults, flags, env, config files, services like consul, etc.).
// These sources are external for your application and use their own data
// formats. Same value may comes from several sources in some priority
// order, in some cases some sources may not be available (like flags in
// tests, but tests still may needs access to configuration values), etc.
// At same time some configuration value (like "port" or "timeout") has
// same format and constraints in your application no matter from which
// source it comes.
//
// Because of all that complexity it makes sense to handle configuration
// in same way as any other external data and implement adapter (in terms
// of Clean Architecture) which will convert configuration received from
// external sources into data structure convenient for your application.
//
// For this we'll need to define two data structures (first suitable for
// receiving values from external sources, second suitable for your
// application) and a conversion logic to create second using values from
// first one.
//
// First data structure should be able to apply same constraints to some
// value no matter from which source it comes from. Luckily, we already
// have suitable interface for this task: flag.Value and the likes. Using
// compatible interface is a requirement to be able to attach flags to
// these values, but as same time this interface allows us to accept
// values as plain strings, which makes it general enough for any sources.
//
// This package provides both a lot of types of Value interface (just like
// flag and other similar packages like pflag) to be used in first
// structure, and also some functions to help loading data from different
// sources (like environment variables) into such Value typed values.
//
// Provided Value interface has more strict semantics than flag.Value, to
// be able to distinguish between zero and unset values - as it is
// important to know is required configuration value was provided or not.
//
// See example to see how to use this package, and also check
// https://github.com/powerman/appcfg/tree/master/examples
// to see how to test such configuration.
package appcfg
