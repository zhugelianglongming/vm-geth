package geth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/xuperchain/xupercore/kernel/contract/bridge"
)

type EvmCreator struct {
	blockCtx vm.BlockContext
	chainCfg *params.ChainConfig
	vmConfig vm.Config
}

func NewEvmCreator(_ *bridge.InstanceCreatorConfig) (bridge.InstanceCreator, error) {
	return new(EvmCreator), nil
}

func (e EvmCreator) CreateInstance(ctx *bridge.Context, cp bridge.ContractCodeProvider) (bridge.Instance, error) {
	txCtx := vm.TxContext{
		GasPrice: big.NewInt(1),
	}
	evm := vm.NewEVM(e.blockCtx, txCtx, NewStateDB(ctx, cp), e.chainCfg, e.vmConfig)
	return &EvmInstance{
		vm:        evm,
		ctx:       ctx,
		cp:        cp,
		fromCache: ctx.ReadFromCache,
	}, nil
}

func (e EvmCreator) RemoveCache(_ string) {
	// nothing to remove yet
}
