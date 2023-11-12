package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/el-tumero/banana-vrf-prototype/client"
	"github.com/el-tumero/banana-vrf-prototype/contract"
	"github.com/holiman/uint256"
)

func main() {
	c, err := client.New()
	if err != nil {
		log.Fatal(err)
	}
	wlt := c.GetWalletAddress()
	rng := c.GetRngAddress()

	log.Print(wlt, "   ", rng)
	log.Println()

	sig, err := c.Sign([]byte{1, 2, 3, 4, 5})
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(len(sig)) // 65 bytes * 8 = 520 bits
	sigHex := hex.EncodeToString(sig)
	log.Println(sigHex)

	slcSig := sig[16:48] // 32 bytes * 8 = 256 bits
	n := new(uint256.Int).SetBytes(slcSig)
	contract.ChangeRandomNumber(n)

	rn, _ := contract.GetRandomNumber()
	fmt.Println(rn.String())
}
