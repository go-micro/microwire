package transport

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/go-micro/microwire/v5/codec"
	"github.com/go-micro/microwire/v5/logger"
)

type Options struct {
	// Addrs is the list of intermediary addresses to connect to
	Addrs []string
	// Codec is the codec interface to use where headers are not supported
	// by the transport and the entire payload must be encoded
	Codec codec.Marshaler
	// Secure tells the transport to secure the connection.
	// In the case TLSConfig is not specified best effort self-signed
	// certs should be used
	Secure bool
	// TLSConfig to secure the connection. The assumption is that this
	// is mTLS keypair
	TLSConfig *tls.Config
	// Timeout sets the timeout for Send/Recv
	Timeout time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
	// Logger is the underline logger
	Logger logger.Logger
}

type DialOptions struct {
	// Tells the transport this is a streaming connection with
	// multiple calls to send/recv and that send may not even be called
	Stream bool
	// Timeout for dialing
	Timeout time.Duration

	// TODO: add tls options when dialing
	// Currently set in global options

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type ListenOptions struct {
	// TODO: add tls options when listening
	// Currently set in global options

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

// Addrs to use for transport.
func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

// Codec sets the codec used for encoding where the transport
// does not support message headers.
func Codec(c codec.Marshaler) Option {
	return func(o *Options) {
		o.Codec = c
	}
}

// Timeout sets the timeout for Send/Recv execution.
func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

// Use secure communication. If TLSConfig is not specified we
// use InsecureSkipVerify and generate a self signed cert.
func Secure(b bool) Option {
	return func(o *Options) {
		o.Secure = b
	}
}

// TLSConfig to be used for the transport.
func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

// Indicates whether this is a streaming connection.
func WithStream() DialOption {
	return func(o *DialOptions) {
		o.Stream = true
	}
}

func WithTimeout(d time.Duration) DialOption {
	return func(o *DialOptions) {
		o.Timeout = d
	}
}

// WithLogger sets the underline logger.
func WithLogger(l logger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

// NetListener Set net.Listener for httpTransport.
func NetListener(customListener net.Listener) ListenOption {
	return func(o *ListenOptions) {
		if customListener == nil {
			return
		}
		if o.Context == nil {
			o.Context = context.TODO()
		}
		o.Context = context.WithValue(o.Context, netListener{}, customListener)
	}
}
