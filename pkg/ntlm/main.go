package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{Dial: (&net.Dialer{}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConns:          10,
			IdleConnTimeout:       30 * time.Second},
	}

	req, err := http.NewRequest("GET", "http://ntlm-auth-server/resource", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("user", "password")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Process the response...
}
