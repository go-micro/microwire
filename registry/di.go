package registry

import (
	"fmt"
	"strings"

	mCli "github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct {
	Plugin    string
	Addresses string
}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin    = "registry"
	cliArgAddresses = "registry_address"
)

func ProvideFlags(
	config *Config,
	cliConfig *mCli.Config,
	c mCli.Cli,
) (*DiFlags, error) {
	if cliConfig.Cli.NoFlags {
		// Defined silently ignore that
		return &DiFlags{}, nil
	}

	result := &DiFlags{}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		mCli.Usage("Registry for discovery. etcd, mdns"),
		mCli.Default(config.Registry.Plugin),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		mCli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
		mCli.Usage("Comma-separated list of registry addresses"),
		mCli.Default(strings.Join(config.Registry.Addresses, ",")),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
		mCli.Destination(&result.Addresses),
	); err != nil {
		return nil, err
	}

	return result, nil
}

func ProvideConfig(
	_ di.DiConfig,
	flags *DiFlags,
	config *Config,
	cliConfig *mCli.Config,
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
	defConfig.Registry.Plugin = flags.Plugin
	defConfig.Registry.Addresses = strings.Split(flags.Addresses, ",")
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
) (Registry, error) {
	if !config.Registry.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Registry.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown registry: %v", err)
	}

	opts := []Option{}
	if len(config.Registry.Addresses) > 0 {
		opts = append(opts, Addrs(config.Registry.Addresses...))
	}

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
