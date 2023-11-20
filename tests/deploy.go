package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/el-tumero/banana-vrf-client/contract"
	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetTestPrivateKey() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA("567eade5964411e5c837c03de980e0e006cfab066f1faffee2b82dea5969a942")
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func GetTestUser(ctx context.Context) (*user.User, error) {
	priv, err := GetTestPrivateKey()
	if err != nil {
		return nil, err
	}
	u, _ := user.NewFromPrivateKey(priv)
	err = u.ConnectToBlockchain(ctx, user.TEST_RPC)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func DeployContract(ctx context.Context, u *user.User) (*contract.Contract, common.Address, error) {
	auth, err := u.PrepareTransactorOpts()
	if err != nil {
		return nil, common.Address{}, nil
	}

	addr, _, instance, err := contract.DeployContract(auth, u.GetBlockchainClient())
	if err != nil {
		return nil, common.Address{}, err
	}

	fmt.Println("contract addr:", addr)
	return instance, addr, err
}
