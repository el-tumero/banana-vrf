package network

import (
	"log"

	"github.com/holiman/uint256"
)

// singleton

type Proposal struct {
	address string
	num     *uint256.Int
}

type Round struct {
	id              int
	contractAddress string
}

var currentRoundId = 0
var proposals []*Proposal
var round *Round

func Init() {
	r := &Round{
		0,
		"0x0",
	}
	round = r
}

func Propose(address string, n *uint256.Int) {
	p := &Proposal{
		address: address,
		num:     n,
	}
	proposals = append(proposals, p)
}

func NextRound(contractAddress string) {
	currentRoundId++
	r := &Round{
		currentRoundId,
		contractAddress,
	}
	round = r
	proposals = []*Proposal{}
}

func DiscoverSmallest() *Proposal {
	smallest := proposals[0]
	for _, p := range proposals[1:] {
		if p.num.Cmp(smallest.num) == -1 {
			smallest = p
		}
	}
	return smallest
}

func ShowProposals() {
	for _, p := range proposals {
		log.Println(p.address, "num:", p.num.Hex())
	}
}

func (p *Proposal) ShowNum() {
	log.Println(p.num.Hex())
}
