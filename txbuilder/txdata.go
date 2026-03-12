package txbuilder

import (
	"github.com/holiman/uint256"
	"github.com/theQRL/go-qrl/common"
	"github.com/theQRL/go-qrl/core/types"
)

type TxMetadata struct {
	GasTipCap  *uint256.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *uint256.Int // a.k.a. maxFeePerGas
	Gas        uint64
	To         *common.Address
	Value      *uint256.Int
	Data       []byte
	AccessList types.AccessList
}
