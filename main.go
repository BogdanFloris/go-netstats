// Copyright 2018 bogdanfloris
// Main program of the go-netstats application
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// bytesTuple is a struct that represents a tuple
// with two fields, bytesIn and bytesOut
type bytesTuple struct {
	bytesIn  int
	bytesOut int
}

var statsMap = make(map[string]*bytesTuple)

func main() {
	parseNettop()
}

func parseNettop() {
	/* Command to get the information about network usage
	flags:
		-P :collapse the results
		-J bytes_in, bytes_out: only get bytes_in and bytes_out columns
		-L 1 one output */
	out, err := exec.Command(
		"nettop", "-P", "-J bytes_in,bytes_out", "-L 1").Output()
	// check for the error
	if err != nil {
		log.Fatal(err)
	}
	// split the output by \n and disregard the first and last line
	outLines := strings.Split(string(out), "\n")[1:]
	outLines = outLines[:len(outLines)-1]
	for _, line := range outLines {
		lineData := strings.Split(line, ",")
		in, err := strconv.Atoi(lineData[2])
		if err != nil {
			log.Fatal(err)
		}
		out, err := strconv.Atoi(lineData[3])
		if err != nil {
			log.Fatal(err)
		}
		tuple := bytesTuple{in, out}
		statsMap[lineData[1]] = &tuple
	}
	for k, v := range statsMap {
		fmt.Printf("%s -> %v\n", k, v)
	}
}
