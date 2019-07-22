package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kitchn-lab/go-adjust/adjust"
)

func main() {
	pw := flag.String("password", "", "Adjust Password")
	email := flag.String("email", "", "Adjust Email")
	flag.Parse()
	if *pw == "" || *email == "" {
		panic("please set password, email and app_id flag")
	}
	session, err := adjust.CreateSession(*email, *pw)
	if err != nil {
		log.Fatalf("Session error: %s", err)
		panic(err)
	}
	fmt.Println(string(session))
	fmt.Println("----------------")
}
