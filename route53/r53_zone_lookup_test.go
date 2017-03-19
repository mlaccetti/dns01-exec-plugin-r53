package main

import (
	"fmt"
	"os"
	"testing"
)

func TestZoneLookup(*testing.T) {
	route53client := r53(Config{Token:os.Getenv("AWS_TOKEN"), Username:os.Getenv("AWS_USERNAME")})

	zoneId, err := getZoneId(route53client, "laccetti.com")
	if err != nil {
		fmt.Printf("Could not get zone ID: %s\n", err.Error())
		return
	}

	fmt.Printf("Found zone ID: %s", zoneId)
}
