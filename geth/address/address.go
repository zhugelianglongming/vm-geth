package address

import (
	"strings"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xuperchain/xupercore/kernel/permission/acl/utils"
	"github.com/xuperchain/xupercore/lib/crypto/hash"
)

const (
	EVMAddressTypePrefixContractName    = "1111"
	EVMAddressTypePrefixContractAccount = "1112"

	typePrefixLen = 4
)

// padding for contract name mapping to EVM address
const padding = "-"

type EVMAddress struct {
	addr common.Address
}

func NewEVMAddress(addr common.Address) EVMAddress {
	return EVMAddress{
		addr: addr,
	}
}

func (a EVMAddress) Address() common.Address {
	return a.addr
}

func (a EVMAddress) GetType() int {
	typePrefix := a.String()[:typePrefixLen]
	switch typePrefix {
	case EVMAddressTypePrefixContractName:
		return XAddressTypeContractName
	case EVMAddressTypePrefixContractAccount:
		return XAddressTypeContractAccount
	default:
		return XAddressTypeDefaultAccount
	}
}

func (a EVMAddress) String() string {
	return string(a.addr.Bytes())
}

func (a EVMAddress) xAddress() XAddress {
	switch a.GetType() {
	case XAddressTypeContractName:
		return a.xAddressContractName()
	case XAddressTypeContractAccount:
		return a.xAddressContractAccount()
	default:
		return a.xAddressDefaultAccount()
	}
}

// xAddressContractName converts EVM address to contract name.
// Contract name is suffix of EVM, which immediately following padding.
// !!!Note: there is no padding for a 16-word contract name
func (a EVMAddress) xAddressContractName() XAddress {
	lastPaddingIdx := strings.LastIndex(a.String(), padding)
	var contractName string
	if lastPaddingIdx == -1 {
		contractName = a.String()[typePrefixLen:]
	} else {
		contractName = a.String()[lastPaddingIdx+1:]
	}

	xAddr := XAddress{
		Address: contractName,
		Type:    XAddressTypeContractName,
	}
	return xAddr
}

// xAddressContractAccount converts EVM address to contract account.
// Contract account number is 16-word-suffix of EVM address.
func (a EVMAddress) xAddressContractAccount() XAddress {
	accountNumber := a.String()[typePrefixLen:]
	accountName := utils.MakeAccountKey("xuper", accountNumber)

	xAddr := XAddress{
		Address: accountName,
		Type:    XAddressTypeContractAccount,
	}
	return xAddr
}

// xAddressDefaultAccount converts EVM address to xuper default account
func (a EVMAddress) xAddressDefaultAccount() XAddress {

	version := []byte(xAddressVersion)
	publicKeyHash := a.addr.Bytes()

	key := append(version, publicKeyHash...)
	checkCode := hash.DoubleSha256(key)

	account := append(key, checkCode[:4]...)

	xAddr := XAddress{
		Address: base58.Encode(account),
		Type:    XAddressTypeDefaultAccount,
	}
	return xAddr
}

const (
	XAddressTypeDefaultAccount = iota
	XAddressTypeContractAccount
	XAddressTypeContractName
)

const xAddressVersion = uint8(1)

// XAddress xuper address
type XAddress struct {
	Address string
	Type    int
}
