package main_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	"github.com/el-tumero/banana-vrf-client/user"
	. "github.com/el-tumero/banana-vrf-tests"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	vrf, err := u.GenerateVrf(rn)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	// hash := crypto.Keccak256(rnu.Bytes())
	// res := crypto.VerifySignature(u.GetPubkey(), hash, vrf[:64])

	r := [32]byte(vrf[0:32])
	s := [32]byte(vrf[32:64])

	signAddr, err := contract.VerifySignature(&bind.CallOpts{}, [32]byte(rn.Bytes()), vrf[64], r, s)
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

	vrf, err := u.GenerateVrf(rn)
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

	vrf, err := u.GenerateVrf(rn)
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

	vrf0, err := u0.GenerateVrf(rn)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	vrf1, err := u1.GenerateVrf(rn)
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

// helper func
func FailAndClose(t *testing.T, args ...any) {
	CloseLocalBlockchain()
	_, _, line, _ := runtime.Caller(1)
	fmt.Println("[Error]", line, "line:", args[0])
	t.FailNow()
}

func TestTwoRoundsScenario(t *testing.T) {
	if err := RunLocalBlockchain(); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	u0, err := CreateTestUserAndDeployContract(ctx, 0)
	if err != nil {
		FailAndClose(t, err)
	}

	u1, err := GetTestUser(ctx, 1)
	if err != nil {
		FailAndClose(t, err)
	}

	err = u1.AddContract(u0.GetContractAddress())
	if err != nil {
		FailAndClose(t, err)
	}

	// Add stake
	if err := u0.AddStake(ctx, uint256.NewInt(150)); err != nil {
		FailAndClose(t, err)
	}
	if err := u1.AddStake(ctx, uint256.NewInt(150)); err != nil {
		FailAndClose(t, err)
	}

	// Check stake
	stake0, err := u0.CheckStake(u0.GetAddress2())
	if err != nil {
		FailAndClose(t, err)
	}
	if stake0.Int64() != int64(150) {
		FailAndClose(t, "Expected 150!")
	}
	stake1, err := u1.CheckStake(u1.GetAddress2())
	if err != nil {
		FailAndClose(t, err)
	}
	if stake1.Int64() != int64(150) {
		FailAndClose(t, "Expected 150!")
	}

	// Check operator status
	if !u0.IsOperatorActive(u0.GetAddress2()) {
		FailAndClose(t, "Expected true - received false")
	}
	if !u1.IsOperatorActive(u1.GetAddress2()) {
		FailAndClose(t, "Expected true - received false")
	}

	rn, err := u0.GetPrevRandomNumber()
	if err != nil {
		FailAndClose(t, err)
	}

	vrf0, err := u0.GenerateVrf(rn)
	if err != nil {
		FailAndClose(t, err)
	}
	vrf1, err := u1.GenerateVrf(rn)
	if err != nil {
		FailAndClose(t, err)
	}

	err = u1.SetRandomNumber(ctx, vrf1)
	if err != nil {
		FailAndClose(t, err)
	}
	err = u0.SetRandomNumber(ctx, vrf0)
	if err != nil {
		FailAndClose(t, err)
	}

	round, err := u0.GetRoundData(1)
	if err != nil {
		FailAndClose(t, err)
	}

	// t.Log(round.Proposer.String())
	// t.Log(u0.GetAddress())
	// t.Log(u1.GetAddress())

	if round.Proposer.String() != u0.GetAddress() {
		FailAndClose(t, "Different proposer than expected!")
	}

	for i := 0; i < 5; i++ {
		NextBlock(ctx)
	}

	blockNumber, _ := u0.GetBlockNumber(ctx)
	t.Log("blockNumber:", blockNumber)

	if err := u1.FinalizeRound(ctx); err == nil {
		FailAndClose(t, "No expected error!")
	}
	if err := u0.FinalizeRound(ctx); err != nil {
		FailAndClose(t, err)
	}

	rn2, err := u0.GetPrevRandomNumber()
	if err != nil {
		FailAndClose(t, err)
	}

	vrf0, err = u0.GenerateVrf(rn2)
	if err != nil {
		FailAndClose(t, err)
	}
	vrf1, err = u1.GenerateVrf(rn2)
	if err != nil {
		FailAndClose(t, err)
	}

	// before
	// round, _ = u0.GetRoundData(2)
	// t.Log(round)

	if user.ConvertVrfToUint256(vrf0).Cmp(user.ConvertVrfToUint256(vrf1)) != -1 {
		FailAndClose(t, "Expected different value!")
	}

	err = u0.SetRandomNumber(ctx, vrf0)
	if err != nil {
		FailAndClose(t, err)
	}
	err = u1.SetRandomNumber(ctx, vrf1)
	if err == nil {
		FailAndClose(t, err)
	}

	// after
	// round, _ = u0.GetRoundData(2)
	// t.Log(round)

	CloseLocalBlockchain()
}
