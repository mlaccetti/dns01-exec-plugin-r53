package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

type Config struct {
  Username string `json:"username"`
  Token    string `json:"token"`
}

func r53(config Config) (*route53.Route53) {
	awsCredentials := credentials.NewStaticCredentials(config.Username, config.Token, "")
	awsConfig := aws.NewConfig()
	awsConfig.Credentials = awsCredentials
	sess := session.Must(session.NewSession(awsConfig))

	return route53.New(sess)
}

func main() {
  apiVersion := os.Getenv("APIVERSION")
  command := os.Getenv("COMMAND")
  domain := os.Getenv("DOMAIN")
  fqdn := os.Getenv("FQDN")
  token := os.Getenv("TOKEN")

  if apiVersion != "v1" {
    os.Exit(3)
  }

  data, err := ioutil.ReadAll(os.Stdin)
  if err != nil {
    io.WriteString(os.Stderr, err.Error())
    os.Exit(1)
  }

  var config Config
  err = json.Unmarshal(data, &config)
  if err != nil {
    io.WriteString(os.Stderr, err.Error())
    os.Exit(2)
  }

	r53client := r53(config)

	zoneId, err := getZoneId(r53client, domain)
	if err != nil {
		io.WriteString(os.Stderr, "Error retreiving zone ID from AWS: " + err.Error())
		os.Exit(1)
	}

  c := newClient(r53client , zoneId)
  switch command {
  case "CREATE":
    err = c.create(domain, fqdn, token)
  case "DELETE":
    err = c.delete(domain, fqdn, token)
  }

  if err != nil {
    io.WriteString(os.Stderr, err.Error())
    os.Exit(1)
  }
  os.Exit(0)
}
