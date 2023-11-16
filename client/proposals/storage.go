package proposals

import (
	"github.com/el-tumero/banana-vrf-client/mock"
	"github.com/holiman/uint256"
)

var storage []*ProposalUint256

func AddProposalToStorage(p *Proposal) {
	if p.ValidateProposal(mock.GetRoundNumber(), mock.GetRandomNumber()) {
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
