package address

import (
	"github.com/ethereum/go-ethereum/common"
)

// Adaptor is address adaptor for Geth EVM implementation
// which make sure to map random generated contract EVM address to contract name
type Adaptor struct {
	contractName string
	contractAddr common.Address
}

// BindContract binds newly generated EVM contract address to its name
func (a *Adaptor) BindContract(name string, addr common.Address) {
	a.contractName = name
	a.contractAddr = addr
}

// E2X transform EVM address to xuper address
func (a *Adaptor) E2X(addr common.Address) XAddress {
	if addr == a.contractAddr {
		xAddr := XAddress{
			Address: a.contractName,
			Type:    XAddressTypeContractName,
		}
		return xAddr
	}
	return NewEVMAddress(addr).xAddress()
}
