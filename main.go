package main

import (
	"log"
	"os"

	"github.com/zeeraw/riksbank/cli"
)

func main() {
	tool := cli.New()
	err := tool.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
