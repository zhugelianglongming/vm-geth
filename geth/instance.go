package geth

import "github.com/xuperchain/xupercore/kernel/contract"

type EvmInstance struct {

}

func (e EvmInstance) Exec() error {
	panic("implement me")
}

func (e EvmInstance) ResourceUsed() contract.Limits {
	panic("implement me")
}

func (e EvmInstance) Release() {
	panic("implement me")
}

func (e EvmInstance) Abort(msg string) {
	panic("implement me")
}
