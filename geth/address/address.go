package address

import (
	"github.com/ethereum/go-ethereum/common"
)

type Address common.Address

func (a Address) Address() common.Address {
	return common.Address(a)
}
