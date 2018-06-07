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

// parses the nettop output
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
	// loop over the output lines from nettop
	storeInMap(outLines)
	// print the map
	for k, v := range statsMap {
		fmt.Printf("%s -> %v\n", k, v)
	}
}

// stores the output from nettop in the map
func storeInMap(outLines []string) {
	for _, line := range outLines {
		// split the line into words
		lineData := strings.Split(line, ",")
		// convert the numbers to ints and check for erros
		in, err := strconv.Atoi(lineData[2])
		if err != nil {
			log.Fatal(err)
		}
		out, err := strconv.Atoi(lineData[3])
		if err != nil {
			log.Fatal(err)
		}
		// make the tuple with the numbers
		tuple := bytesTuple{in, out}
		// store them in the map
		statsMap[lineData[1]] = &tuple
	}
}
