package main

import (
	"log"
	"os"

	"github.com/zeeraw/riksbank/cli"
)

func main() {
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
