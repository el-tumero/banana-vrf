package proposals

import "github.com/el-tumero/banana-vrf-client/mock"

var storage []*Proposal

func AddProposalToStorage(p *Proposal) {
	if p.ValidateProposal(mock.GetRoundNumber(), mock.GetRandomNumber()) {
		storage = append(storage, p)
	}
}

func FlushStorage() {
	storage = []*Proposal{}
}
