package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	statuses := []int{
		http.StatusOK,
		http.StatusNotFound,
		http.StatusInternalServerError,
	}

	for {
		randomIndex := rand.Intn(len(statuses))
		randomStatus := statuses[randomIndex]
		randomStatusText := http.StatusText(randomStatus)
		fmt.Fprintf(os.Stdout, "Responding with %d (%s)\n",
			randomStatus, randomStatusText)
		time.Sleep(time.Second)
	}
}
