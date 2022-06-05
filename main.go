package main

import (
	"gochain/cli"
	"gochain/model"
)

func main() {
	bc := model.CreateBlockchain()
	defer bc.Close()

	cli := cli.CLI{Chain: bc}
	cli.Run()
}
