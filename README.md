# go-netstats
A simple command-line bandwidth monitor for macOS written in Go.
# Description
go-netstats is a simple command-line application that gets data about network usage from the nettop program, parses it and displays the amount of data that was sent and received over the network.
# Usage
First build the executable file:
```
go build netstats.go
```
Then to run the program:
```
./netstats
There are two flags:
    -c: Starts a continuous stream, meaning the the program is run in a infinite loop and outputs the result every 5 seconds.
    -s int: Modifies the default 5 seconds sleep timer.
```
# Example Output
```
Total received bytes: 7.2 MB
Total sent bytes: 2.4 MB
```
