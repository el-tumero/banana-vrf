package client

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
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

	return &client{
		rngPrivKey:    rngpk,
		walletPrivKey: wltpk,
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

func BytesToBigInt(data []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(data)
	return n
}
