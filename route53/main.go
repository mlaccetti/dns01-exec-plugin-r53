// Copyright 2017 <???> All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
  "encoding/json"
  "io"
  "io/ioutil"
  "os"
)

type Config struct {
  Username string `json:"username"`
  Token    string `json:"token"`
}

func main() {
  apiVersion := os.Getenv("APIVERSION")
  command := os.Getenv("COMMAND")
  domain := os.Getenv("DOMAIN")
  fqdn := os.Getenv("FQDN")
  token := os.Getenv("TOKEN")

  // AWS
  zoneID := os.Getenv("ZONEID")

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

  c, err := newClient(zoneID)
  if err != nil {
    io.WriteString(os.Stderr, "Error creating google DNS client"+err.Error())
    os.Exit(1)
  }

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
