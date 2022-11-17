package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
		if err := scanner.Err(); err != nil {
			log.Printf("could not read form input:%v \n", err)
		}
	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var sprRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Erorr %v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error :%v\n", err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spfi") {
			hasSPF = true
			sprRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, record := range dmarcRecords {
		hasDMARC = true
		dmarcRecord = record
		break
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord)
}
