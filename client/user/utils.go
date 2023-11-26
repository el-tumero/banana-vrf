package user

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

func Keccak256Abi(data *big.Int) ([]byte, error) {
	uint256Ty, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	args := abi.Arguments{
		{
			Type: uint256Ty,
		},
	}
	encoded, err := args.Pack(
		data,
	)
	if err != nil {
		return nil, err
	}

	hashed := crypto.Keccak256(encoded)
	return hashed, nil
}
