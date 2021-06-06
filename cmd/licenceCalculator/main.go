package main

import (
	"os"
	"fmt"
	"log"

	"github.com/ryanchew3/flexera-coding-challenge/pkg/file"
	"github.com/ryanchew3/flexera-coding-challenge/pkg/checker"
)

func main() {
	f := os.Getenv("IN_FILE")
	fmt.Println(f)

	d, err := file.ReadFile(f)
	if err != nil {
		log.Println(err.Error())
		return
	}

	out := checker.ProcessChecks(d)
	fmt.Println(out)
}
