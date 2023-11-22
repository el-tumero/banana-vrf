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
	u, err := GetTestUser(ctx, 0)
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

	u, err := GetTestUser(ctx, 0)
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
	u, err := CreateTestUserAndDeployContract(ctx, 0)

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
	u, err := CreateTestUserAndDeployContract(ctx, 0)

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

	if err := u.SetRandomNumber(ctx, vrf); err != nil {
		t.Log(err)
		t.Fail()
		return
	}

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

func TestSetRandomNumberTwoUsers(t *testing.T) {
	if err := RunLocalBlockchain(); err != nil {
		t.Fatal(err)
	}
	defer CloseLocalBlockchain()

	ctx := context.Background()
	u0, err := CreateTestUserAndDeployContract(ctx, 0)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	u1, err := GetTestUser(ctx, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	err = u1.AddContract(u0.GetContractAddress())
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	rn, err := u0.GetPrevRandomNumber()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	rnu := uint256.MustFromBig(rn)

	vrf0, err := u0.GenerateVrf(rnu)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	vrf1, err := u1.GenerateVrf(rnu)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	n0 := user.ConvertVrfToUint256(vrf0)
	n1 := user.ConvertVrfToUint256(vrf1)
	// t.Log(n0, n1)

	compare := n0.Cmp(n1)
	_ = compare
	// -1 -> n0 < n1

	if err := u0.SetRandomNumber(ctx, vrf0); err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if err := u1.SetRandomNumber(ctx, vrf1); err == nil {
		t.Log("No expected error!")
		t.Fail()
		return
	}

	t.Log(err)

	num, err := u0.GetCurrRandomNumber()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if n0.ToBig().Cmp(num) != 0 {
		t.Log("Wrong current rand number! (different than predicted)")
		t.Fail()
		return
	}

	round, err := u0.GetRoundData(1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if round.State != 1 {
		t.Log("Wrong round state!")
		t.Fail()
		return
	}

	t.Log(round)

}

func TestNextBlock(t *testing.T) {
	if err := RunLocalBlockchain(); err != nil {
		t.Fatal(err)
	}
	defer CloseLocalBlockchain()

	ctx := context.Background()
	u, err := GetTestUser(ctx, 0)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if err := NextBlock(ctx); err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	// time.Sleep(10 * time.Second)
	blockNumber, err := u.GetBlockNumber(ctx)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if blockNumber != 1 {
		t.Log("Block number is expected to be 1")
		t.Fail()
		return
	}

	t.Log("blockNumber:", blockNumber)
}
