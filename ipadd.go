package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getPublicIP() string {
  resp, err := http.Get("https://ipecho.net/plain")
  if err != nil {
    fmt.Println("Could not query IP address")
    os.Exit(1)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("Could not parse IP addr respone body")
    os.Exit(1)
  }
  return string(body)
}
