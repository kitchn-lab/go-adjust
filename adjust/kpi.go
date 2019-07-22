package adjust

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// KPIService struct to hold individual service
type KPIService service

// KPI sturct to descripe response
type KPI struct {
	ResultParameters ResultParameters `json:"result_parameters"`
	ResultSet        ResultSet        `json:"result_set"`
}

// ResultParameters reflects input param
type ResultParameters struct {
	AttributionSource string      `json:"attribution_source"`
	AttributionType   string      `json:"attribution_type"`
	Events            []Event     `json:"events,omitempty"`
	Grouping          []string    `json:"grouping"`
	EndDate           string      `json:"end_date"`
	StartDate         string      `json:"start_date"`
	UTCOffset         string      `json:"utc_offset"`
	HumanReadableKpis bool        `json:"human_readable_kpis"`
	Kpis              []string    `json:"kpis"`
	Reattributed      string      `json:"reattributed"`
	Sandbox           bool        `json:"sandbox"`
	Trackers          []RPTracker `json:"trackers,omitempty"`
}

type Event struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

// RPTracker to present tracker in ResultParameters
type RPTracker struct {
	Token          string `json:"token"`
	Name           string `json:"name"`
	HasSubtrackers bool   `json:"has_subtrackers"`
}

// ResultSet to present information over current tracker data
type ResultSet struct {
	Token    string    `json:"token"`
	Name     string    `json:"name"`
	Currency string    `json:"currency"`
	Trackers []Tracker `json:"trackers,omitempty"`
	Dates    []Date    `json:"dates,omitempty"`
}

type Date struct {
	Date     string    `json:"date"`
	Networks []Network `json:"networks"`
}

type Network struct {
	Token     string     `json:"token"`
	Name      string     `json:"name"`
	Campaigns []Campaign `json:"campaigns"`
}

type Campaign struct {
	Token    string    `json:"token"`
	Name     string    `json:"name"`
	AdGroups []AdGroup `json:"adgroups"`
}

//ID converts the name to an ID "GB,IE_NonBrand_JG (284979913)"
func (c Campaign) ID() (int64, error) {
	idx := strings.LastIndex(c.Name, "(")
	sid := c.Name[idx+1 : len(c.Name)-1]
	return strconv.ParseInt(sid, 10, 64)
}

type AdGroup struct {
	Token     string     `json:"token"`
	Name      string     `json:"name"`
	Creatives []Creative `json:"creatives"`
}

//ID converts the name to an ID "Exact Match (285227483)"00
func (a AdGroup) ID() (int64, error) {
	idx := strings.LastIndex(a.Name, "(")
	sid := a.Name[idx+1 : len(a.Name)-1]
	return strconv.ParseInt(sid, 10, 64)
}

type Creative struct {
	Token     string    `json:"token"`
	Name      string    `json:"name"`
	KpiValues []float64 `json:"kpi_values"`
}

// Tracker holds its token and kpi values.
type Tracker struct {
	Token     string    `json:"token"`
	KPIValues []float64 `json:"kpi_values"`
}

// Options parameter you can add to the url
type Options struct {
	AttributionSource string `url:"attribution_source,omitempty"`
	AttributionType   string `url:"attribution_type,omitempty"`
	EventKpis         string `url:"event_kpis,omitempty"` // x,y,z format
	Grouping          string `url:"grouping,omitempty"`   // x,y,z format
	EndDate           string `url:"end_date,omitempty"`   // "2019-05-09"
	StartDate         string `url:"start_date,omitempty"` // "2019-05-09"
	UTCOffset         string `url:"utc_offset,omitempty"` // "00:00"
	HumanReadableKpis bool   `url:"human_readable_kpis,omitempty"`
	Kpis              string `url:"kpis,omitempty"`
	Reattributed      string `url:"reattributed,omitempty"`
	Sandbox           bool   `url:"sandbox,omitempty"`
}

// List function to get Kpis
func (s *KPIService) List(ctx context.Context, opt *Options) (*KPI, *Response, error) {
	u, err := addOptions(fmt.Sprintf("kpis/v1/%s.json", s.client.AppID), opt)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(u)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	kpis := &KPI{}

	resp, err := s.client.Do(ctx, req, kpis)
	if err != nil {
		return nil, resp, err
	}

	return kpis, resp, nil
}
