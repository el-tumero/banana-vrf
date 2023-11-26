package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/el-tumero/banana-vrf-client/contract"
	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// deploy script
	rpcAddr := flag.String("rpc", user.TEST_RPC, "rpc address")
	chainId := flag.Int64("chain_id", 1337, "chain id")
	privateKey := flag.String("priv", "567eade5964411e5c837c03de980e0e006cfab066f1faffee2b82dea5969a942", "private key address")

	flag.Parse()
	ctx := context.Background()

	priv, err := crypto.HexToECDSA(*privateKey)
	if err != nil {
		log.Fatal("private key conversion err", err)
	}

	u, err := user.NewFromPrivateKey(priv)
	if err != nil {
		log.Fatal("can't create user ", err)
	}
	user.CHAIN_ID = *chainId
	u.ConnectToBlockchain(ctx, *rpcAddr)

	_, _, err = DeployContract(ctx, u)
	if err != nil {
		log.Fatal("can't deploy contract ", err)
	}
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
