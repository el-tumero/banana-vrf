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

var PRIVATE_KEYS = [2]string{"567eade5964411e5c837c03de980e0e006cfab066f1faffee2b82dea5969a942", "8b44176a6734b87519422284f98c3ce7c979bd540ea823b0cf486b06c628e865"}

func GetTestPrivateKey(id int) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(PRIVATE_KEYS[id])
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func GetTestUser(ctx context.Context, id int) (*user.User, error) {
	priv, err := GetTestPrivateKey(id)
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
	auth, err := u.PrepareTransactorOpts(1_000_000)
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

func CreateTestUserAndDeployContract(ctx context.Context, id int) (*user.User, error) {
	u, err := GetTestUser(ctx, id)
	if err != nil {
		return nil, err
	}

	_, conAddr, err := DeployContract(ctx, u)
	if err != nil {
		return nil, err
	}

	err = u.AddContract(conAddr)
	if err != nil {
		return nil, err
	}

	return u, nil
}
