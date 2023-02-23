package geth

import (
	"github.com/zhugelianglongming/vm-geth/geth/address"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/xuperchain/xupercore/kernel/contract/bridge"
)

type stateDB struct {
	ctx         *bridge.Context
	cp          bridge.ContractCodeProvider
	addrAdaptor *address.Adaptor
}

func NewStateDB(ctx *bridge.Context, cp bridge.ContractCodeProvider) *stateDB {
	return &stateDB{
		ctx:         ctx,
		cp:          cp,
		addrAdaptor: new(address.Adaptor),
	}
}

func (s stateDB) CreateAccount(address common.Address) {
	s.addrAdaptor.BindContract(s.ctx.ContractName, address)
}

func (s stateDB) SubBalance(address common.Address, b *big.Int) {
	panic("implement me")
}

func (s stateDB) AddBalance(address common.Address, b *big.Int) {
	panic("implement me")
}

func (s stateDB) GetBalance(address common.Address) *big.Int {
	panic("implement me")
}

func (s stateDB) GetNonce(address common.Address) uint64 {
	panic("implement me")
}

func (s stateDB) SetNonce(address common.Address, u uint64) {
	panic("implement me")
}

func (s stateDB) GetCodeHash(address common.Address) common.Hash {
	panic("implement me")
}

func (s stateDB) GetCode(addr common.Address) []byte {
	xAddr := s.addrAdaptor.E2X(addr)
	if xAddr.Type != address.XAddressTypeContractName {
		return nil
	}
	contractName := xAddr.Address

	if s.ctx.ReadFromCache {
		code, _ := s.cp.GetContractCodeFromCache(contractName)
		return code
	}
	code, _ := s.cp.GetContractCode(contractName)
	return code
}

func (s stateDB) SetCode(addr common.Address, bytes []byte) {
	xAddr := s.addrAdaptor.E2X(addr)
	if xAddr.Type != address.XAddressTypeContractName {
		return
	}
	contractName := xAddr.Address

	_ = s.ctx.State.Put("contract", evmCodeKey(contractName), bytes)
}

func (s stateDB) GetCodeSize(address common.Address) int {
	panic("implement me")
}

func (s stateDB) AddRefund(u uint64) {
	panic("implement me")
}

func (s stateDB) SubRefund(u uint64) {
	panic("implement me")
}

func (s stateDB) GetRefund() uint64 {
	panic("implement me")
}

func (s stateDB) GetCommittedState(address common.Address, hash common.Hash) common.Hash {
	panic("implement me")
}

func (s stateDB) GetState(address common.Address, hash common.Hash) common.Hash {
	panic("implement me")
}

func (s stateDB) SetState(address common.Address, hash common.Hash, hash2 common.Hash) {
	panic("implement me")
}

func (s stateDB) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	panic("implement me")
}

func (s stateDB) SetTransientState(addr common.Address, key, value common.Hash) {
	panic("implement me")
}

func (s stateDB) Suicide(address common.Address) bool {
	panic("implement me")
}

func (s stateDB) HasSuicided(address common.Address) bool {
	panic("implement me")
}

func (s stateDB) Exist(address common.Address) bool {
	panic("implement me")
}

func (s stateDB) Empty(address common.Address) bool {
	panic("implement me")
}

func (s stateDB) AddressInAccessList(addr common.Address) bool {
	panic("implement me")
}

func (s stateDB) SlotInAccessList(addr common.Address, slot common.Hash) (addressOk bool, slotOk bool) {
	panic("implement me")
}

func (s stateDB) AddAddressToAccessList(addr common.Address) {
	panic("implement me")
}

func (s stateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	panic("implement me")
}

func (s stateDB) Prepare(rules params.Rules, sender, coinbase common.Address, dest *common.Address, precompiles []common.Address, txAccesses types.AccessList) {
	panic("implement me")
}

func (s stateDB) RevertToSnapshot(i int) {
	panic("implement me")
}

func (s stateDB) Snapshot() int {
	panic("implement me")
}

func (s stateDB) AddLog(log *types.Log) {
	panic("implement me")
}

func (s stateDB) AddPreimage(hash common.Hash, bytes []byte) {
	panic("implement me")
}

func (s stateDB) ForEachStorage(address common.Address, f func(common.Hash, common.Hash) bool) error {
	panic("implement me")
}

func evmCodeKey(contractName string) []byte {
	return []byte(contractName + ".code")
}
