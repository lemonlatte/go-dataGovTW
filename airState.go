package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const DATA_URI = "http://opendata.epa.gov.tw/ws/Data/AQX/?format=json"

type PublishTime time.Time

func (t *PublishTime) MarshalJSON() ([]byte, error) {
	b := time.Time(*t).Format("2016-04-28 01:00")
	return []byte(b), nil
}

func (pt *PublishTime) UnmarshalJSON(b []byte) (err error) {
	t, err := time.Parse("2006-01-02 15:04", string(b[1:len(b)-1]))
	if err == nil {
		*pt = PublishTime(t)
	}
	return
}

type AirState struct {
	County      string
	SiteName    string
	PSI         string
	PM10        string
	PM2_5       string `json:"PM2.5"`
	O3          string
	CO          string
	NO2         string
	SO2         string
	PublishTime *PublishTime
}

func GetAirState() (err error) {
	resp, err := http.Get(DATA_URI)
	if err != nil {
		return err
	}
	as := make([]AirState, 0)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&as)
	if err == nil {
		fmt.Printf("%+v", as)
	}
	return
}
