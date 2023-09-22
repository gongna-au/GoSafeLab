package main

import (
	"log"

	"github.com/stacktitan/smb"
)

func main() {
	options := &smb.Options{
		Host:        "192.168.1.1",
		Port:        445,
		User:        "username",
		Domain:      "domain",
		Workstation: "",
		Password:    "password",
	}

	session, err := smb.NewSession(options, false)
	if err != nil {
		log.Fatalln("[!] Failed to create session:", err)
	}
	defer session.Logoff()

	if session.IsAuthenticated {
		log.Println("[+] Successfully authenticated")
	} else {
		log.Fatalln("[!] Failed to authenticate")
	}
}
