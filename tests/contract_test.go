package main_test

import (
	"context"
	"testing"

	"github.com/el-tumero/banana-vrf-client/user"
	. "github.com/el-tumero/banana-vrf-tests"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

func TestDeploy(t *testing.T) {
	if err := RunLocalBlockchain(); err != nil {
		t.Fatal(err)
	}
	defer CloseLocalBlockchain()

	ctx := context.Background()
	u, err := GetTestUser(ctx)
	if err != nil {
		t.Log(err)
		return
	}

	_, _, err = DeployContract(ctx, u)
	if err != nil {
		t.Log(err)
		return
	}
}

func TestVerify(t *testing.T) {
	if err := RunLocalBlockchain(); err != nil {
		t.Fatal(err)
	}
	defer CloseLocalBlockchain()

	ctx := context.Background()

	u, err := GetTestUser(ctx)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	contract, addr, err := DeployContract(ctx, u)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if err := u.AddContract(addr); err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	rn, err := u.GetPrevRandomNumber()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	rnu := uint256.MustFromBig(rn)
	hash := crypto.Keccak256(rnu.Bytes())
	t.Log(new(uint256.Int).SetBytes(hash).Hex())

	vrf, err := u.GenerateVrf(rnu)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	// hash := crypto.Keccak256(rnu.Bytes())
	// res := crypto.VerifySignature(u.GetPubkey(), hash, vrf[:64])

	r := [32]byte(vrf[0:32])
	s := [32]byte(vrf[32:64])

	signAddr, err := contract.VerifySignature(&bind.CallOpts{}, rnu.Bytes32(), vrf[64], r, s)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if u.GetAddress() != signAddr.String() {
		t.Log("Signature not valid!")
		t.Fail()
		return
	}
}

func TestVerify2(t *testing.T) {
	if err := RunLocalBlockchain(); err != nil {
		t.Fatal(err)
	}
	defer CloseLocalBlockchain()

	ctx := context.Background()
	u, err := CreateTestUserAndDeployContract(ctx)

	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	rn, err := u.GetPrevRandomNumber()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	rnu := uint256.MustFromBig(rn)
	vrf, err := u.GenerateVrf(rnu)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if res := u.VerifyRandomNumber(vrf); res != true {
		t.Log(err)
		t.Fail()
		return
	}
}

func TestSetRandomNumber(t *testing.T) {
	if err := RunLocalBlockchain(); err != nil {
		t.Fatal(err)
	}
	defer CloseLocalBlockchain()

	ctx := context.Background()
	u, err := CreateTestUserAndDeployContract(ctx)

	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	rn, err := u.GetPrevRandomNumber()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	rnu := uint256.MustFromBig(rn)
	vrf, err := u.GenerateVrf(rnu)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if err := u.SetRandomNumber(vrf); err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	t.Log()

	num, err := u.GetCurrRandomNumber()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if user.ConvertVrfToUint256(vrf).ToBig().Text(16) != num.Text(16) {
		t.Log("Wrong current rand number! (different than predicted)")
		t.Fail()
		return
	}

}
