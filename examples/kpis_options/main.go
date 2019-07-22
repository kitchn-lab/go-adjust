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
	opt := adjust.Options{
		AttributionSource: "dynamic",
		AttributionType:   "click",
		StartDate:         "2019-05-01",
		EndDate:           "2019-05-09",
		UTCOffset:         "00:00",
		EventKpis:         "tam0sc_events,xxwrmf_events,wk9hmc_events,tucsbx_events",
		Kpis:              "installs",
		Sandbox:           false,
		HumanReadableKpis: false,
		Grouping:          "day,networks,campaigns,adgroups,creatives",
		Reattributed:      "all",
	}
	resp, _, err := client.KPI.List(context.Background(), &opt)
	if err != nil {
		log.Fatalf("Kpis List error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&resp)
	fmt.Println(string(res))
	fmt.Println("----------------")
	for _, date := range resp.ResultSet.Dates {
		for _, network := range date.Networks {
			for _, campaign := range network.Campaigns {
				id, err := campaign.ID()
				if err != nil {
					panic(err)
				}
				fmt.Println("Campaign Name: ", campaign.Name)
				fmt.Println("Campaign ID: ", id)
				for _, adGroup := range campaign.AdGroups {
					id, err := adGroup.ID()
					if err != nil {
						panic(err)
					}
					fmt.Println("AdGroup Name: ", adGroup.Name)
					fmt.Println("AdGroup ID: ", id)
					for _, creative := range adGroup.Creatives {
						fmt.Println("Keyword: ", creative.Name)
						for i, kpi := range creative.KpiValues {
							fmt.Println("KPI KEY: ", resp.ResultParameters.Kpis[i])
							fmt.Println("KPI VALUE: ", kpi)
						}
					}
				}
			}
		}
	}
}
