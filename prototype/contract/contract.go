package contract

import (
	"fmt"

	"github.com/holiman/uint256"
)

var blockNumber = 0
var randomNumber *uint256.Int
var isChanged = false

func DebugSetRandomNumber() {
	randomNumber = uint256.MustFromHex("0xcf75f41566b7465c315a1a50768931d389e05b34962628614b72b1d4642c77c6")
}

func DebugGetRandomNumber() *uint256.Int {
	return randomNumber
}

func ChangeRandomNumber(n *uint256.Int) {
	if isChanged {
		if randomNumber.Cmp(n) == 1 {
			randomNumber = n
		}
		return
	}
	randomNumber = n
	isChanged = true
	return
}

func GetRandomNumber() (*uint256.Int, error) {
	if !isChanged {
		return nil, fmt.Errorf("Random number not yet selected\n")
	}
	return randomNumber, nil
}

func NextBlock() {
	blockNumber++
	isChanged = false
}
