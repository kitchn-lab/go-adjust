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
	ok, err := adjust.ValidAccountCredentials(*email, *pw)
	if err != nil {
		log.Fatalf("Session error: %s", err)
		panic(err)
	}
	fmt.Println(ok)
	fmt.Println("----------------")
}
