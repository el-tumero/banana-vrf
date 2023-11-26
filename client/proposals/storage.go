package proposals

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/el-tumero/banana-vrf-client/user"
)

var storage []*ProposalBigInt

func AddProposalToStorage(p *Proposal, u *user.User) {
	// check if there is already record like this
	if FindProposal(p.Vrf) {
		fmt.Println("Already in storage!")
		return
	}
	roundId, err := u.GetCurrRoundId()
	if err != nil {
		fmt.Println("Can't fetch roundId!")
		return
	}
	roundData, err := u.GetRoundData(roundId - 1)
	if err != nil {
		fmt.Println("Can't fetch roundData")
		return
	}

	// log.Println("add to storage...")

	// fmt.Println(u.GetAddress2().String())
	// fmt.Println(u.IsOperatorActive(u.GetAddress2()))

	if p.ValidateProposal(roundId, roundData.RandomNumber, u) {
		pu := &ProposalBigInt{
			Num:   new(big.Int).SetBytes(p.Vrf[16:48]),
			Round: p.Round,
			Vrf:   p.Vrf,
		}
		storage = append(storage, pu)
		// fmt.Println("APPEND")
	}
}

func GetStorage() []*ProposalBigInt {
	return storage
}

func FindProposal(vrf []byte) bool {
	for i := range storage {
		if bytes.Equal(storage[i].Vrf, vrf) {
			return true
		}
	}
	return false
}

func PrintStorage() {
	for i := 0; i < len(storage); i++ {
		fmt.Println(storage[i].Num)
	}
}

func DiscoverSmallest() *ProposalBigInt {
	smallest := storage[0]
	for _, p := range storage[1:] {

		if p.Num.Cmp(smallest.Num) == -1 {
			smallest = p
		}
	}
	return smallest
}

func FlushStorage() {
	storage = storage[:0]
}
