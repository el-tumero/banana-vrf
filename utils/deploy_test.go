package main_test

import (
	"context"
	"testing"

	"github.com/el-tumero/banana-vrf-client/user"
	. "github.com/el-tumero/banana-vrf-utils"
)

func TestDeploy(t *testing.T) {
	ctx := context.Background()
	priv, err := GetTestPrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	u, _ := user.NewFromPrivateKey(priv)
	err = u.ConnectToBlockchain(ctx, user.TEST_RPC)
	if err != nil {
		t.Fatal(err)
	}
	err = DeployContract(ctx, u)
	if err != nil {
		t.Fatal(err)
	}
}
