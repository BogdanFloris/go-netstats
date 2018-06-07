// Copyright 2018 bogdanfloris
// Main program of the go-netstats application
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("nettop", "-P", "-J bytes_in,bytes_out", "-L 1").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
