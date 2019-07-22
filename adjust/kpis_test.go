package adjust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var testCases = []struct {
	name    string
	path    string
	fixture string
	opts    *Options
}{
	{
		"No Options",
		"/kpis/v1/random.json",
		"kpis.json",
		nil,
	},
	{
		"With Options",
		"/kpis/v1/random.json",
		"kpis_options.json",
		&Options{
			AttributionSource: "dynamic",
			AttributionType:   "click",
			StartDate:         "2019-05-01",
			EndDate:           "2019-05-09",
			UTCOffset:         "00:00",
			EventKpis:         "tam0sc_events,xxwrmf_events,wk9hmc_events,tucsbx_events",
			Kpis:              "installs",
			Sandbox:           false,
			HumanReadableKpis: true,
			Grouping:          "day,networks,campaigns,adgroups,creatives",
			Reattributed:      "all",
		},
	},
}

func TestKPIService_List(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, mux, _, teardown := setup()
			t.Log(fmt.Sprintf("Setup Done %s", tc.name))
			defer teardown()
			mux.HandleFunc(tc.path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				w.WriteHeader(http.StatusOK)
				w.Write(loadFixture(tc.fixture))
			})
			got, _, err := client.KPI.List(context.Background(), tc.opts)
			if err != nil {
				t.Errorf("TestKPIService - %s : KPI.List returned error: %v", tc.name, err)
			}
			want := &KPI{}
			responseToInterface(loadFixture(tc.fixture), &want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("TestKPIService - %s : KPI.List = %+v, want %+v", tc.name, got, want)
			}
		})
	}
}
