package example

/*
A demo for usage.
It shows a package registers geth type InstanceCreator during its init.
*/

import (
	"plugin"

	"github.com/xuperchain/xupercore/kernel/contract/bridge"
)

type CreatorFunc = func(*bridge.InstanceCreatorConfig) (bridge.InstanceCreator, error)

const (
	typeGeth = "geth"
	driverName = "geth"
)

func init() {

	// load plugin
	p, err := plugin.Open("dir/vm-path.so")
	if err != nil {
		panic(err)
	}

	// find symbol
	creatorFunc, err := p.Lookup("NewEvmCreator")
	if err != nil {
		panic(err)
	}

	// convert type
	newEvmCreator, ok := creatorFunc.(CreatorFunc)
	if !ok {
		panic("unexpected type from vm-path module symbol")
	}

	// use it
	bridge.Register(typeGeth, driverName, newEvmCreator)
}