package main

import (
	"log"

	"github.com/el-tumero/banana-vrf-prototype/client"
	"github.com/el-tumero/banana-vrf-prototype/contract"
	"github.com/el-tumero/banana-vrf-prototype/network"
)

func main() {
	c0, err := client.New()
	if err != nil {
		log.Fatal(err)
	}

	c1, err := client.New()
	if err != nil {
		log.Fatal(err)
	}

	network.Init()
	contract.DebugSetRandomNumber()

	err = c0.Propose()
	if err != nil {
		log.Fatal(err)
	}

	err = c1.Propose()
	if err != nil {
		log.Fatal(err)
	}

	network.ShowProposals()

	smol := network.DiscoverSmallest()

	smol.ShowNum()

}
