package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/digitalocean/godo"
)

// Wrapper around the Digital Ocean library
// only exposes functionality that we are using
type DigitalOceanClient struct {
	Domain    *string
	Token     string
	Context   context.Context
	APIClient *godo.Client
}

func createClient(domain *string, token string) (*DigitalOceanClient, error) {
	ctx := context.TODO()
	doClient := godo.NewFromToken(token)
	// Always succeeds :)
	return &DigitalOceanClient{Domain: domain, Context: ctx, Token: token, APIClient: doClient}, nil
}

// Get the root A record for the domain passed as a command line argument
func (c DigitalOceanClient) GetDomainRecord() int {
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	domainRecords, _, err := c.APIClient.Domains.RecordsByType(c.Context, *c.Domain, "A", opt)
	if err != nil {
		fmt.Println("could not fetch records for domain")
		fmt.Println(err)
		os.Exit(1)
	}
	var rootRecordId int
	for _, record := range domainRecords {
		// only want to modify the root domain (for now).
		if record.Name == "@" {
			rootRecordId = record.ID
			break
		}
	}
	return rootRecordId
}

// Update the root A record for the domain passed as a command line argument
func (c DigitalOceanClient) UpdateDomainRecord(publicIPAddress string, recordId int) {
	editRequest := &godo.DomainRecordEditRequest{
		Type: "A",
		Data: publicIPAddress,
		TTL:  1000,
	}
	newRecord, _, err := c.APIClient.Domains.EditRecord(c.Context, *c.Domain, recordId, editRequest)
	if err != nil {
		fmt.Println("Could not update domain record")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Record successfully updated")
	recordString, err := json.MarshalIndent(newRecord, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal json")
		os.Exit(1)
	}
	fmt.Println(string(recordString))
}
