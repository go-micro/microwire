package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	"github.com/google/wire"
)

var DiBrokerSet = wire.NewSet(
	ProvideBrokerConfigStore,
	mBroker.DiSet,
)

var DiRegistrySet = wire.NewSet(
	ProvideRegistryConfigStore,
	mRegistry.DiSet,
)

var DiTransportSet = wire.NewSet(
	ProvideTransportConfigStore,
	mTransport.DiSet,
)

var DiAllComponentsSuperSet = wire.NewSet(
	DiBrokerSet,
	DiRegistrySet,
	DiTransportSet,
)

var DiAllComponentProvidersSet = wire.NewSet(
	mBroker.Provide,
	mRegistry.Provide,
	mTransport.Provide,
)

var DiCliSet = wire.NewSet(
	ProvideCliConfigStore,
	ProvideCLI,
	ProvideCliArgs,
	ProvideInitializedCLI,
)

var DiMicroServiceSet = wire.NewSet(
	ProvideMicroOpts,
	ProvideMicroService,
)
