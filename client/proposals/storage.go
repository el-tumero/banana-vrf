package proposals

import (
	"fmt"

	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/holiman/uint256"
)

var storage []*ProposalUint256

func AddProposalToStorage(p *Proposal, u *user.User) {
	// temp values

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

	// fmt.Println(u.GetAddress2().String())
	// fmt.Println(u.IsOperatorActive(u.GetAddress2()))

	if p.ValidateProposal(roundId, roundData.RandomNumber, u) {
		pu := &ProposalUint256{
			Num:   new(uint256.Int).SetBytes(p.Vrf[16:48]),
			Round: p.Round,
			Vrf:   p.Vrf,
		}
		storage = append(storage, pu)
	}
}

func GetStorage() []*ProposalUint256 {
	return storage
}

func PrintStorage() {
	for i := 0; i < len(storage); i++ {
		fmt.Println(storage[i].Num)
	}
}

func DiscoverSmallest() *ProposalUint256 {
	smallest := storage[0]
	for _, p := range storage[1:] {

		if p.Num.Cmp(smallest.Num) == -1 {
			smallest = p
		}
	}
	return smallest
}

func FlushStorage() {
	storage = []*ProposalUint256{}
}
