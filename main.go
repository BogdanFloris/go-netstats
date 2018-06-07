// Copyright 2018 bogdanfloris
// Main program of the go-netstats application
package main

import (
	"fmt"
	"log"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

// bytesTuple is a struct that represents a tuple
// with two fields, bytesIn and bytesOut
type bytesTuple struct {
	bytesIn  int64
	bytesOut int64
}

var statsMap = make(map[string]*bytesTuple)

func main() {
	parseNettop()
	// print the map
	for k, v := range statsMap {
		fmt.Printf("%s -> %v\n", k, *v)
	}
	// get the sums
	sumIn, sumOut := sumUsage()
	fmt.Printf("Total received bytes: %s\n", humanReadbleByteCount(sumIn))
	fmt.Printf("Total sent bytes: %s\n", humanReadbleByteCount(sumOut))
}

func sumUsage() (int64, int64) {
	var sumBytesIn int64
	var sumBytesOut int64
	for _, v := range statsMap {
		sumBytesIn += v.bytesIn
		sumBytesOut += v.bytesOut
	}
	return sumBytesIn, sumBytesOut
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
}

// stores the output from nettop in the map
func storeInMap(outLines []string) {
	for _, line := range outLines {
		// split the line into words
		lineData := strings.Split(line, ",")
		// convert the numbers to ints and check for erros
		in, err := strconv.ParseInt(lineData[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		out, err := strconv.ParseInt(lineData[3], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		// make the tuple with the numbers
		tuple := bytesTuple{in, out}
		// store them in the map
		statsMap[lineData[1]] = &tuple
	}
}

func humanReadbleByteCount(bytes int64) string {
	var unit int64 = 1024
	if bytes < unit {
		return string(bytes) + "B"
	}
	exp := int64(math.Log(float64(bytes)) / math.Log(float64(unit)))
	pre := "kMGTPE"[exp-1]
	return fmt.Sprintf("%.1f %cB", float64(bytes)/math.Pow(float64(unit), float64(exp)), pre)
}
