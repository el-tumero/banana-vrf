package user_test

import (
	"context"
	"testing"

	"github.com/el-tumero/banana-vrf-client/mock"
	"github.com/el-tumero/banana-vrf-client/proposals"
	"github.com/el-tumero/banana-vrf-client/user"
)

func TestGenerateVrf(t *testing.T) {
	u, err := user.New()
	if err != nil {
		t.Fatal(err)
	}

	vrf, err := u.GenerateVrf(mock.GetRandomNumber())
	if err != nil {
		t.Fatal(err)
	}

	if len(vrf) != 65 {
		t.FailNow()
	}
}

func TestPrepareAndValidate(t *testing.T) {
	u, err := user.New()
	if err != nil {
		t.Fatal(err)
	}

	// contract stuff
	mock.Init()
	mock.AddPubkeyToStakers([65]byte(u.GetPubkey()), 15)

	vrf, err := u.GenerateVrf(mock.GetRandomNumber())
	if err != nil {
		t.Fatal(err)
	}

	p := &proposals.Proposal{
		Round: 1,
		Vrf:   vrf,
	}

	p.ValidateProposal(1, mock.GetRandomNumber())

	prepared, err := p.Prepare()
	if err != nil {
		t.Fatal(err)
	}

	if len(prepared) != proposals.PROPOSAL_LEN {
		t.Fatal("Wrong proposal lenght!")
	}

	pc, err := proposals.CastBytes(prepared)
	if err != nil {
		t.Fatal(err)
	}

	res := pc.ValidateProposal(1, mock.GetRandomNumber())

	if res == false {
		t.Fatal("Not valid proposal!")
	}
}

func TestGetStake(t *testing.T) {
	res := mock.GetStake([65]byte{1, 4, 3, 2, 4})
	if res != 0 {
		t.FailNow()
	}
}

func TestConnectToBlockchain(t *testing.T) {
	ctx := context.Background()
	u, err := user.New()
	if err != nil {
		t.Fatal(err)
	}

	err = u.ConnectToBlockchain(ctx, user.TEST_RPC)
	if err != nil {
		t.Fatal(err)
	}

}
