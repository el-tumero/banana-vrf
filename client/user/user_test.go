package user_test

import (
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

	vrf, err := u.GenerateVrf(mock.GENESIS_RANDOM_NUMBER)
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

	vrf, err := u.GenerateVrf(mock.GENESIS_RANDOM_NUMBER)
	if err != nil {
		t.Fatal(err)
	}

	p := &proposals.Proposal{
		Round: 1,
		Vrf:   vrf,
	}

	p.ValidateProposal(1, mock.GENESIS_RANDOM_NUMBER)

	prepared, err := p.Prepare()
	if err != nil {
		t.Fatal(err)
	}

	if len(prepared) != 69 {
		t.Fatal("Wrong proposal lenght!")
	}

	pc, err := proposals.CastBytes(prepared)
	if err != nil {
		t.Fatal(err)
	}

	res := pc.ValidateProposal(1, mock.GENESIS_RANDOM_NUMBER)

	if res == false {
		t.Fatal("Not valid proposal!")
	}
}
