package user

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/el-tumero/banana-vrf-client/contract"
	"github.com/el-tumero/banana-vrf-client/proposals"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	"github.com/holiman/uint256"
)

const TEST_RPC = "http://127.0.0.1:8545/"
const CONTRACT_ADDR = "0x631d896D88F9f02668DFDFFC20fA3cCCD12e4bD1"
const TEST_CHAIN_ID = 1337

type User struct {
	address      string
	privateKey   *ecdsa.PrivateKey
	blc          *ethclient.Client
	contract     *contract.Contract
	contractAddr common.Address
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

func NewFromPrivateKey(priv *ecdsa.PrivateKey) (*User, error) {
	var user *User = &User{}
	user.privateKey = priv
	user.address = crypto.PubkeyToAddress(priv.PublicKey).String()
	return user, nil
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

func (u *User) AddContract(addr common.Address) error {
	vrfHost, err := contract.NewContract(addr, u.blc)
	if err != nil {
		return err
	}
	u.contract = vrfHost
	u.contractAddr = addr
	return nil
}

func (u *User) GetPrevRandomNumber() (*big.Int, error) {
	data, err := u.contract.GetPreviousRandomNumber(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *User) GetCurrRandomNumber() (*big.Int, error) {
	data, err := u.contract.GetCurrentRandomNumber(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *User) PrepareTransactorOpts(limit uint64) (*bind.TransactOpts, error) {
	nonce, err := u.blc.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(u.privateKey.PublicKey))
	if err != nil {
		return nil, err
	}
	gasPrice, err := u.blc.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	t, err := bind.NewKeyedTransactorWithChainID(u.privateKey, big.NewInt(TEST_CHAIN_ID))
	if err != nil {
		return nil, err
	}

	t.Nonce = big.NewInt(int64(nonce))
	t.Value = big.NewInt(0) // wei
	t.GasLimit = limit      // in units
	t.GasPrice = gasPrice

	return t, nil
}

func (u *User) GetBlockchainClient() *ethclient.Client {
	return u.blc
}

func (u *User) VerifyRandomNumber(vrf []byte) bool {

	r := [32]byte(vrf[0:32])
	s := [32]byte(vrf[32:64])
	v := vrf[64]

	isVerified, err := u.contract.VerifyProposal(&bind.CallOpts{}, v, r, s)
	if err != nil {
		return false
	}

	return isVerified
}

func (u *User) SetRandomNumber(ctx context.Context, vrf []byte) error {
	r := [32]byte(vrf[0:32])
	s := [32]byte(vrf[32:64])
	v := vrf[64]

	tran, err := u.PrepareTransactorOpts(300_000)
	if err != nil {
		return err
	}

	tx, err := u.contract.SetRandomNumber(tran, v, r, s)
	if err != nil {
		return err
	}

	// TODO: wait for failure or success
	recp, err := u.blc.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return err
	}

	if recp.Status != 1 {
		return fmt.Errorf("tx failed")
	}

	return nil
}

func (u *User) FinalizeRound(ctx context.Context) error {
	tran, err := u.PrepareTransactorOpts(300_000)
	if err != nil {
		return err
	}

	tx, err := u.contract.NextRound(tran)
	if err != nil {
		return err
	}

	recp, err := u.blc.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return err
	}

	if recp.Status != 1 {
		return fmt.Errorf("tx failed")
	}

	return nil
}

func (u *User) GetRoundData(id uint32) (*contract.VRFHostRound, error) {
	round, err := u.contract.GetRound(&bind.CallOpts{}, id)
	if err != nil {
		return nil, err
	}
	return &round, nil
}

func (u *User) GetContractAddress() common.Address {
	return u.contractAddr
}

func (u *User) GetBlockNumber(ctx context.Context) (uint64, error) {
	blockNumber, err := u.blc.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}

func ConvertVrfToUint256(vrf []byte) *uint256.Int {
	return new(uint256.Int).SetBytes(vrf[16:48])
}
