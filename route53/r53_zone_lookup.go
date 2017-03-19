package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

func getZoneId(r53 *route53.Route53, domain string) (string, error) {
	params := &route53.ListHostedZonesByNameInput{
		DNSName:      &domain,
		MaxItems:     aws.String("1"),
	}
	resp, err := r53.ListHostedZonesByName(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return "", err
	}

	// Pretty-print the response data.
	fmt.Println(resp)

	if len(resp.HostedZones) > 0 {
		if strings.Contains(*resp.HostedZones[0].Name, domain) {
			fmt.Println("Found our domain.")
			return strings.Split(*resp.HostedZones[0].Id, "/")[2], nil
		}
	}

	return "", errors.New("Could not find the specified domain in Route53.")
}
