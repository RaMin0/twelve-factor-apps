package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var numOf500s int

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		fmt.Printf("LOG: %s\n", line)

		statusString := regexp.MustCompile("\\d+").FindString(line)
		status, _ := strconv.Atoi(statusString)
		if status == http.StatusInternalServerError {
			numOf500s++
			if numOf500s%5 == 0 {
				fmt.Println("LOG: ##### Oops! #####")
			}
		}
	}
}
