package main

import (
	"fmt"
)

func main() {
	var userpass = "admin"
	var clientID = "044b8ad6006845c29446b2f18e5b5909"
	var host = "https://vnextTrainingps.dev.logisensebilling.com"

	client, err := NewClient(host, userpass, userpass, clientID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(client)
}
