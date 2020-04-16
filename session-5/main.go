package main

import (
	"fmt"
	"net/http"
)

var (
	processes = []*Process{
		&Process{},
		&Process{},
		&Process{},
	}
	currentProcessIndex = 0

	sharedVisits int
)

type Process struct{}

func (p *Process) IncrementVisits() { sharedVisits++ }

func loadBalancer(w http.ResponseWriter, r *http.Request) {
	currentProcess := processes[currentProcessIndex]

	currentProcess.IncrementVisits()
	fmt.Fprintf(w, "Process: %d\n", currentProcessIndex)
	fmt.Fprintf(w, "Visits:  %d", sharedVisits)

	currentProcessIndex++
	if currentProcessIndex == len(processes) {
		currentProcessIndex = 0
	}
}

func main() {
	http.HandleFunc("/", loadBalancer)
	http.ListenAndServe("localhost:3000", nil)
}
