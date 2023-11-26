package main_test

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/el-tumero/banana-vrf-client/user"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestHash(t *testing.T) {
	prev, _ := big.NewInt(0).SetString("268924301244112151111193779238987776559164197311682689858984696294031985235", 10)
	uint256Ty, _ := abi.NewType("uint256", "", nil)

	args := abi.Arguments{
		{
			Type: uint256Ty,
		},
	}
	encoded, _ := args.Pack(
		prev,
	)

	hash := crypto.Keccak256(encoded)

	if strings.Compare(hex.EncodeToString(hash), "f278c634217dfdeec27aa49cd1aadf83700e5c5be82df35bb0ce5129625d9a90") != 0 {
		t.Fatal("Wrong hash!")
	}
}

func TestHash2(t *testing.T) {
	prev, _ := big.NewInt(0).SetString("110674241778222555159086117039744400750421889117878875670049083117269247344050", 10)
	hash, _ := user.Keccak256Abi(prev)
	if strings.Compare(hex.EncodeToString(hash), "4d5f7dd6652bc469c41cc53d0c515d0254469f469272633b5b34ad32e2a10a0c") != 0 {
		t.Fatal("Wrong hash!")
	}
}
