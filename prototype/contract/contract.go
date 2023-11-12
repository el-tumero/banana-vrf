package contract

import (
	"fmt"

	"github.com/holiman/uint256"
)

var blockNumber = 0
var randomNumber *uint256.Int
var isChanged = false

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
