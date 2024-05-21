package main

import (
	"github.com/ngyewch/mise-helper/cmd"
	"log"
	"os"
)

func main() {
	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
