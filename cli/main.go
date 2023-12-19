package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "banana",
		Usage: "Managing your BananaVRF client",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"it"},
				Usage:   "Initializes CLI",
				Action:  Init,
			},
			{
				Name:    "start",
				Aliases: []string{"str"},
				Usage:   "Starts BananaVRF client",
				Action:  Start,
			},
			{
				Name:    "stop",
				Aliases: []string{"stp"},
				Usage:   "Stops BananaVRF client",
				Action:  Stop,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Start(ctx *cli.Context) error {
	file, err := os.OpenFile("temp.txt", os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	cmd := exec.Command("nohup", "test_proc/test_proc", "&")
	err = cmd.Start()
	if err != nil {
		return err
	}
	pid := cmd.Process.Pid
	_, err = file.WriteString(strconv.Itoa(pid))
	if err != nil {
		return err
	}

	fmt.Println(pid)
	fmt.Println("Client started!")
	return nil
}

func Stop(ctx *cli.Context) error {
	file, err := os.Open("temp.txt")
	if err != nil {
		return err
	}

	buf := make([]byte, 10)
	_, err = file.Read(buf)
	if err != nil {
		return err
	}
	buf = bytes.Trim(buf, "\x00")

	cmd := exec.Command("kill", "-9", string(buf))
	err = cmd.Start()
	if err != nil {
		return err
	}
	fmt.Println("Client stopped!")

	return nil
}

func Init(ctx *cli.Context) error {
	file, err := os.Create("temp.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
