package user

import (
	"context"
	"math/big"
	"strings"

	"github.com/el-tumero/banana-vrf-client/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

var localAbi abi.ABI

func (u *User) EstimatorAddAbi(contractAbi string) error {
	abi, err := abi.JSON(strings.NewReader(contract.ContractABI))
	if err != nil {
		return err
	}
	localAbi = abi
	return nil
}

func (u *User) EstimateSetRandomNumber(ctx context.Context, v uint8, r [32]byte, s [32]byte) error {
	data, err := localAbi.Pack("setRandomNumber", v, r, s)
	if err != nil {
		return err
	}

	msg := ethereum.CallMsg{
		From:  crypto.PubkeyToAddress(u.privateKey.PublicKey),
		To:    &(u.contractAddr),
		Value: big.NewInt(0),
		Data:  data,
	}

	_, err = u.blc.EstimateGas(ctx, msg)
	if err != nil {
		return err
	}

	return nil

}

func (u *User) EstimateFinalizeRound(ctx context.Context) error {
	data, err := localAbi.Pack("nextRound")
	if err != nil {
		return err
	}

	msg := ethereum.CallMsg{
		From:  crypto.PubkeyToAddress(u.privateKey.PublicKey),
		To:    &(u.contractAddr),
		Value: big.NewInt(0),
		Data:  data,
	}

	_, err = u.blc.EstimateGas(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) EstimateAddStake(ctx context.Context, amount *uint256.Int) error {
	data, err := localAbi.Pack("addStake")
	if err != nil {
		return err
	}

	msg := ethereum.CallMsg{
		From:  crypto.PubkeyToAddress(u.privateKey.PublicKey),
		To:    &(u.contractAddr),
		Value: amount.ToBig(),
		Data:  data,
	}

	_, err = u.blc.EstimateGas(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
