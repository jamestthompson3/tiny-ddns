package main

import (
	"fmt"
	"os"
)

func main() {
	parsedDomain := parseArgs()
	fmt.Printf("\nFetching Info for: \n   %s\n", *parsedDomain)
	token := os.Getenv("DIGITALOCEAN_TOKEN")
	client, _ := createClient(parsedDomain, token)
	rootRecordId := client.GetDomainRecord()
	publicIP := getPublicIP()
	client.UpdateDomainRecord(publicIP, rootRecordId)
}
