// Code generated with jinja2 templates. DO NOT EDIT.

package cache

import (
	"fmt"

	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct{}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin = "cache"
)

func ProvideFlags(
	config *Config,
	cliConfig *cli.Config,
	c cli.Cli,
) (DiFlags, error) {
	if cliConfig.Cli.NoFlags {
		// Defined silently ignore that
		return DiFlags{}, nil
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		cli.Usage("Cache k/v store"),
		cli.Default(config.Cache.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	return DiFlags{}, nil
}

func ProvideConfig(
	_ di.DiConfig,
	flags DiFlags,
	config *Config,
	c cli.Cli,
	cliConfig *cli.Config,
	configor config.Config,
) (DiConfig, error) {
	defConfig := NewConfig()

	if configor != nil {
		if err := configor.Scan(defConfig); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	if cliConfig.Cli.NoFlags {
		// Dont parse flags if NoFlags has been given
		return DiConfig{}, nil
	}

	defConfig = NewConfig()
	if f, ok := c.Get(cliArgPlugin); ok {
		defConfig.Cache.Plugin = cli.FlagValue(f, defConfig.Cache.Plugin)
	}

	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func ProvideConfigNoFlags(
	config *Config,
	configor config.Config,
) (DiConfig, error) {
	defConfig := NewConfig()

	if configor != nil {
		if err := configor.Scan(defConfig); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,

	config *Config,
) (Cache, error) {
	if !config.Cache.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Cache.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown cache: %v", err)
	}

	opts := []Option{}

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
