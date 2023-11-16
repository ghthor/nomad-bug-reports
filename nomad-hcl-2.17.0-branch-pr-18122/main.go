package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/nomad/jobspec2"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	must(err)
	defer f.Close()
	_, err = jobspec2.Parse("", f)
	fmt.Println(err)
}
