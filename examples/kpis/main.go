package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/kitchn-lab/go-adjust/adjust"
)

func main() {
	pw := flag.String("password", "", "Adjust Password")
	email := flag.String("email", "", "Adjust Email")
	appID := flag.String("app_id", "", "Adjust app ID")
	flag.Parse()
	if *pw == "" || *email == "" || *appID == "" {
		panic("please set password, email and app_id flag")
	}
	client, err := adjust.NewClient(nil, *email, *pw, *appID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	opt := adjust.Options{}
	list, _, err := client.KPI.List(context.Background(), &opt)
	if err != nil {
		log.Fatalf("Kpis List error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&list)
	fmt.Println(string(res))
	fmt.Println("----------------")
}
