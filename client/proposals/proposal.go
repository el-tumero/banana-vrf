package proposals

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

const PROPOSAL_LEN = 69
const MINIMAL_STAKE = 10

type Proposal struct {
	Round uint32
	Vrf   []byte
}

type ProposalUint256 struct {
	Num   *uint256.Int
	Round uint32
	Vrf   []byte
}

func CastBytes(data []byte) (*Proposal, error) {
	if len(data) > PROPOSAL_LEN {
		return nil, fmt.Errorf("Input too long! (%d)", len(data))
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

func (p *Proposal) ValidateProposal(round uint32, rnd *uint256.Int) bool {
	if p.Round != round {
		return false
	}
	// for round 1 -> get round 0 random number
	hashed := crypto.Keccak256(rnd.Bytes())

	pubkey, err := crypto.Ecrecover(hashed, p.Vrf)
	if err != nil {
		return false
	}

	// check if user stakes more than minimal stake
	// if mock.GetStake([65]byte(pubkey)) < MINIMAL_STAKE {
	// 	return false
	// }

	isValid := crypto.VerifySignature(pubkey, hashed, p.Vrf[:64])

	return isValid
}
