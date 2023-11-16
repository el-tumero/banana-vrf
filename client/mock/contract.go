package mock

import "github.com/holiman/uint256"

var randomNumber = uint256.MustFromHex("0xcf75f41566b7465c315a1a50768931d389e05b34962628614b72b1d4642c77c6")
var roundNumber uint32 = 1

// list of staking pubkeys

var stakers map[[65]byte]int

func Init() {
	stakers = make(map[[65]byte]int)
}

func AddPubkeyToStakers(pubkey [65]byte, stake int) {
	stakers[pubkey] = stake
}

func GetStake(pubkey [65]byte) int {
	return stakers[pubkey]
}

func AddNewRandomNumber(rn *uint256.Int) {
	randomNumber = rn
	roundNumber++
}

func GetRoundNumber() uint32 {
	return roundNumber
}

func GetRandomNumber() *uint256.Int {
	return randomNumber
}
