package geth

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/xuperchain/xupercore/kernel/contract"
	"github.com/xuperchain/xupercore/kernel/contract/bridge"
	"github.com/xuperchain/xupercore/kernel/contract/bridge/pb"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/xuperchain/vm-geth/geth/address"
)

const (
	MethodInitialize = "initialize"
)

type EvmInstance struct {
	vm        *vm.EVM
	ctx       *bridge.Context
	cp        bridge.ContractCodeProvider
	code      []byte
	abi       []byte
	gasUsed   uint64
	fromCache bool
}

func (i EvmInstance) Exec() error {
	var (
		caller address.Address
		callee common.Address
		input  []byte
		gas    uint64
		value  *big.Int

		ret          []byte
		leftOverGas  uint64
		err          error
	)

	// TODO: params

	if i.ctx.Method == MethodInitialize {
		ret, _, leftOverGas, err = i.vm.Create(caller, i.code, gas, value)
	} else {
		ret, leftOverGas, err = i.vm.Call(caller, callee, input, gas, value)
	}
	if err != nil {
		return err
	}

	i.gasUsed = uint64(contract.MaxLimits.Cpu) - leftOverGas - i.vm.StateDB.GetRefund()
	i.ctx.Output = &pb.Response{
		Status: http.StatusOK,
		Body: ret,
	}
	return nil
}

func (i EvmInstance) ResourceUsed() contract.Limits {
	panic("implement me")
}

func (i EvmInstance) Release() {
	panic("implement me")
}

func (i EvmInstance) Abort(msg string) {
	panic("implement me")
}
