module github.com/asim/go-micro/plugins/registry/gossip/v3

go 1.16

require (
	go-micro.dev/v4 v4.0.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/hashicorp/memberlist v0.1.5
	github.com/mitchellh/hashstructure v1.1.0
)

replace github.com/asim/go-micro/v3 => ../../../../go-micro
