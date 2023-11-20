package main

import (
	"os/exec"
	"time"
)

var cmd *exec.Cmd

func RunLocalBlockchain() error {
	cmd = exec.Command("ganache", "-s", "hello")
	err := cmd.Start()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	return nil
}

func CloseLocalBlockchain() error {
	err := cmd.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}
