package client

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/el-tumero/banana-vrf-prototype/contract"
	"github.com/el-tumero/banana-vrf-prototype/network"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

type client struct {
	rngPrivKey    *ecdsa.PrivateKey
	rngAddress    string
	walletPrivKey *ecdsa.PrivateKey
	walletAddress string
}

func New() (*client, error) {
	rngpk, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	wltpk, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	rngpub := crypto.PubkeyToAddress(rngpk.PublicKey).String()
	wltpub := crypto.PubkeyToAddress(wltpk.PublicKey).String()

	return &client{
		rngPrivKey:    rngpk,
		rngAddress:    rngpub,
		walletPrivKey: wltpk,
		walletAddress: wltpub,
	}, nil
}

func (c *client) GetWalletAddress() string {
	addr := crypto.PubkeyToAddress(c.walletPrivKey.PublicKey)
	return addr.String()
}

func (c *client) GetRngAddress() string {
	addr := crypto.PubkeyToAddress(c.rngPrivKey.PublicKey)
	return addr.String()
}

func (c *client) Sign(data []byte) ([]byte, error) {
	hashed := crypto.Keccak256(data)
	sig, err := crypto.Sign(hashed, c.rngPrivKey)
	if err != nil {
		return nil, nil
	}
	return sig, nil
}

func (c *client) Propose() error {
	rn := contract.DebugGetRandomNumber()
	// if err != nil {
	// 	return err
	// }

	sig, err := c.Sign(rn.Bytes())
	if err != nil {
		return err
	}

	slcSig := sig[16:48]
	n := new(uint256.Int).SetBytes(slcSig)

	network.Propose(c.rngAddress, n)
	return nil
}

func BytesToBigInt(data []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(data)
	return n
}
