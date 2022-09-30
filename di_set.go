package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	mRegistry "github.com/go-micro/microwire/registry"
	mStore "github.com/go-micro/microwire/store"
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

var DiStoreSet = wire.NewSet(
	ProvideStoreConfigStore,
	mStore.DiSet,
)

var DiTransportSet = wire.NewSet(
	ProvideTransportConfigStore,
	mTransport.DiSet,
)

// DiAllComponentsSuperSet is a set of all things components need, except the components themself.
var DiAllComponentsSuperSet = wire.NewSet(
	DiBrokerSet,
	DiRegistrySet,
	DiStoreSet,
	DiTransportSet,
)

var DiCliSet = wire.NewSet(
	ProvideCliConfigStore,
	ProvideCLI,
	ProvideCliArgs,
	ProvideInitializedCLI,
	mCli.DiSet,
)

var DiMicroServiceSet = wire.NewSet(
	ProvideMicroOpts,
	ProvideMicroService,
)
