package geth

import (
	"github.com/xuperchain/xupercore/kernel/contract/bridge"
)

type EvmCreator struct {
}

func NewEvmCreator(_ *bridge.InstanceCreatorConfig) (bridge.InstanceCreator, error) {
	return new(EvmCreator), nil
}

func (e EvmCreator) CreateInstance(_ *bridge.Context, _ bridge.ContractCodeProvider) (bridge.Instance, error) {
	return new(EvmInstance), nil
}

func (e EvmCreator) RemoveCache(name string) {
	panic("implement me")
}
