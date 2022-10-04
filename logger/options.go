package logger

import (
	"context"
	"io"
)

type Option func(*Options)

type Options struct {
	// Config is the configuration of this component.
	Config *Config
	// The logging level the logger should log at. default is `InfoLevel`
	Level Level
	// fields to always be logged
	Fields map[string]interface{}
	// It's common to set this to a file, or leave it default which is `os.Stderr`
	Out io.Writer
	// Caller skip frame count for file:line info
	CallerSkipCount int
	// Alternative options
	Context context.Context
}

// WithConfig sets the config to Options.
func WithConfig(n *Config) Option {
	return func(o *Options) {
		o.Config = n
	}
}

// WithFields set default fields for the logger.
func WithFields(fields map[string]interface{}) Option {
	return func(args *Options) {
		args.Fields = fields
	}
}

// WithLevel set default level for the logger.
func WithLevel(level Level) Option {
	return func(args *Options) {
		args.Level = level
	}
}

// WithOutput set default output writer for the logger.
func WithOutput(out io.Writer) Option {
	return func(args *Options) {
		args.Out = out
	}
}

// WithCallerSkipCount set frame count to skip.
func WithCallerSkipCount(c int) Option {
	return func(args *Options) {
		args.CallerSkipCount = c
	}
}

func SetOption(k, v interface{}) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

func ConfigToOpts(c *Config) []Option {
	lvl, err := GetLevel(c.Level)
	if err != nil {
		lvl = InfoLevel
	}

	opts := make([]Option, 4)
	opts[0] = WithConfig(c)
	opts[1] = WithLevel(lvl)
	opts[2] = WithFields(c.Fields)
	opts[3] = WithCallerSkipCount(c.CallerSkipCount)

	return opts
}
