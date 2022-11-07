package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Sprintf("domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
		if err := scanner.Err(); err != nil {
			log.Fatal("could not read form input: %v\n", err)
		}
	}

}

func checkDomain(domain string) {
	
}
