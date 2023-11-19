package user

import (
	"context"
	"crypto/ecdsa"

	"github.com/el-tumero/banana-vrf-client/proposals"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	"github.com/holiman/uint256"
)

const TEST_RPC = "http://127.0.0.1:8545/"

type User struct {
	address    string
	privateKey *ecdsa.PrivateKey
	blc        *ethclient.Client
}

func New() (*User, error) {
	pk, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	var user *User = &User{}
	user.privateKey = pk
	user.address = crypto.PubkeyToAddress(pk.PublicKey).String()
	return user, nil

}

func NewFromPrivateKey() (*User, error) {
	// TO BE IMPLEMENTED
	return nil, nil
}

func (u *User) sign(data []byte) ([]byte, error) {
	hashed := crypto.Keccak256(data)
	sig, err := crypto.Sign(hashed, u.privateKey)
	if err != nil {
		return nil, nil
	}
	return sig, nil
}

func (u *User) GetAddress() string {
	return u.address
}

func (u *User) GetPubkey() []byte {
	return crypto.FromECDSAPub(&u.privateKey.PublicKey)
}

func (u *User) GenerateVrf(root *uint256.Int) ([]byte, error) {
	sig, err := u.sign(root.Bytes())
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (u *User) Propose(conn *websocket.Conn, vrf []byte) error {
	p := &proposals.Proposal{
		Round: 1,
		Vrf:   vrf,
	}

	prepared, err := p.Prepare()
	if err != nil {
		return err
	}

	// send to websocket server
	err = conn.WriteMessage(1, prepared)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) ConnectToBlockchain(ctx context.Context, url string) error {
	c, err := ethclient.Dial(url)
	if err != nil {
		return err
	}

	_, err = c.BlockNumber(ctx)
	if err != nil {
		return err
	}

	u.blc = c
	return nil
}

func (u *User) GetPrevRandomNumber() {

}
