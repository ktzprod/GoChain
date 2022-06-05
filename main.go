package main

import (
	"gochain/cli"
	"gochain/model"
)

func main() {
	bc := model.CreateBlockchain()
	defer bc.Close()

	cli := cli.CLI{bc}
	cli.Run()
}
