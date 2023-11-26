package proposals

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
)

const PROPOSAL_LEN = 69
const MINIMAL_STAKE = 100

type Proposal struct {
	Round uint32
	Vrf   []byte
}

type ProposalBigInt struct {
	Num   *big.Int
	Round uint32
	Vrf   []byte
}

var LastLocalProposal *Proposal
var Candidate bool = false

func CastBytes(data []byte) (*Proposal, error) {
	if len(data) > PROPOSAL_LEN {
		return nil, fmt.Errorf("input too long! (%d)", len(data))
	}

	var p *Proposal = &Proposal{}
	p.Round = binary.BigEndian.Uint32(data[0:4])
	p.Vrf = data[4:]

	return p, nil
}

func (p *Proposal) Prepare() ([]byte, error) {
	// 4 -> round, 65 -> signature
	rb := make([]byte, 4)
	binary.BigEndian.PutUint32(rb, p.Round)
	out := append(rb, p.Vrf...)
	return out, nil
}

func Propose(conn *websocket.Conn, roundId uint32, vrf []byte) error {
	p := &Proposal{
		Round: roundId,
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

	LastLocalProposal = p

	return nil
}

func (p *Proposal) ValidateProposal(round uint32, rnd *big.Int, u *user.User) bool {
	if p.Round != round {
		return false
	}

	// hashed := crypto.Keccak256(rnd.Bytes())
	hashed, err := user.Keccak256Abi(rnd)
	if err != nil {
		fmt.Println("Hash err")
		return false
	}

	pubkey, err := crypto.SigToPub(hashed, p.Vrf)
	if err != nil {
		fmt.Println("SigToPub err ", err)
		return false
	}
	res := u.IsOperatorActive(crypto.PubkeyToAddress(*pubkey))
	if !res {
		fmt.Println("Not operator!")
		return false
	}

	return true
}
