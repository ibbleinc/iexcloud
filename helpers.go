// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Date models a report date
type Date time.Time

// UnmarshalJSON implements the Unmarshaler interface for Date.
func (d *Date) UnmarshalJSON(data []byte) error {
	var aux string
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return fmt.Errorf("error unmarshaling date to string: %s", err)
	}
	t, err := time.Parse("2006-01-02", aux)
	if err != nil {
		return fmt.Errorf("error converting string to date: %s", err)
	}
	*d = Date(t)
	return nil
}

// MarshalJSON implements the Marshaler interface for Date.
func (d *Date) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	return json.Marshal(t.Format("2006-01-02"))
}

// PathRange refers to the date range used in the path of an endpoint.
type PathRange int

// Enum values for PathRange.
const (
	Mo1 PathRange = iota
	Mo3
	Mo6
	Yr1
	Yr2
	Yr5
	YTD
	Next
)

var pathRangeDescription = map[PathRange]string{
	Mo1:  "One month (default)",
	Mo3:  "Three months",
	Mo6:  "Six months",
	Yr1:  "One year",
	Yr2:  "Two years",
	Yr5:  "Five years",
	YTD:  "Year-to-data",
	Next: "Next upcoming",
}

// PathRanges maps the string keys from the JSON to the PathRange
// constant values.
var PathRanges = map[string]PathRange{
	"next": Next,
	"1m":   Mo1,
	"3m":   Mo3,
	"6m":   Mo6,
	"5y":   Yr5,
	"2y":   Yr2,
	"1y":   Yr1,
	"ytd":  YTD,
}

// PathRangeJSON maps a PathRange to the string used in the JSON.
var PathRangeJSON = map[PathRange]string{
	Mo1:  "1m",
	Mo3:  "3m",
	Mo6:  "6m",
	Yr1:  "1y",
	Yr2:  "2y",
	Yr5:  "5y",
	YTD:  "ytd",
	Next: "next",
}

// MarshalJSON implements the Marshaler interface for PathRange.
func (p *PathRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(PathRangeJSON[*p])
}

// UnmarshalJSON implements the Unmarshaler interface for PathRange.
func (p *PathRange) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("error unmarshaling path range, got %s", data)
	}
	return p.Set(s)
}

// Set sets the issue type using a string.
func (p *PathRange) Set(s string) error {
	// Ensure the provided string matches on the keys in the map.
	got, ok := PathRanges[s]
	if !ok {
		return fmt.Errorf("invalid issue type %q", s)
	}
	// Set the issue type to the value found in the map per the key.
	*p = got
	return nil
}

// String implements the Stringer interface for PathRange.
func (p PathRange) String() string {
	return pathRangeDescription[p]
}

// EpochTime refers to unix timestamps used for some fields in the API
type EpochTime time.Time

// MarshalJSON implements the Marshaler interface for EpochTime.
func (e EpochTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(time.Time(e).Unix())), nil
}

// UnmarshalJSON implements the Unmarshaler interface for EpochTime.
func (e *EpochTime) UnmarshalJSON(data []byte) (err error) {
	s := string(data)
	if s == "null" {
		return
	}
	ts, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	// Per docs: If the value is -1, IEX has not quoted the symbol in the trading day.
	if ts == -1 {
		return
	}

	*e = EpochTime(time.Unix(int64(ts)/1000, 0))
	return
}

// String implements the Stringer interface for EpochTime.
func (e EpochTime) String() string { return time.Time(e).String() }

// IssueType refers to the common issue type of the stock.
type IssueType int

const (
	blank IssueType = iota
	ad
	re
	ce
	si
	lp
	cs
	et
)

var issueTypeDescription = map[IssueType]string{
	ad:    "American Depository Receipt (ADR)",
	re:    "Real Estate Investment Trust (REIT)",
	ce:    "Closed end fund (Stock and Bond Fund)",
	si:    "Secondary Issue",
	lp:    "Limited Partnership",
	cs:    "Common Stock",
	et:    "Exchange Traded Fund (ETF)",
	blank: "Not available",
}

// IssueTypes maps the string keys from the JSON to the IssueType constant
// values.
var IssueTypes = map[string]IssueType{
	"ad":  ad,
	"re":  re,
	"ce":  ce,
	"cef": ce,
	"si":  si,
	"lp":  lp,
	"cs":  cs,
	"et":  et,
	"":    blank,
}

// IssueTypeJSON maps an IssueType to the string used in the JSON.
var IssueTypeJSON = map[IssueType]string{
	ad:    "ad",
	re:    "re",
	ce:    "ce",
	si:    "si",
	lp:    "lp",
	cs:    "cs",
	et:    "et",
	blank: "",
}

// String implements the Stringer interface for IssueType.
func (i IssueType) String() string {
	return issueTypeDescription[i]
}

// MarshalJSON implements the Marshaler interface for IssueType.
func (i *IssueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(IssueTypeJSON[*i])
}

// UnmarshalJSON implements the Unmarshaler interface for IssueType.
func (i *IssueType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("issueType should be a string, got %s", data)
	}
	return i.Set(s)
}

// Set sets the issue type using a string.
func (i *IssueType) Set(s string) error {
	// Ensure the provided string matches on the keys in the map.
	got, ok := IssueTypes[s]
	if !ok {
		return fmt.Errorf("invalid issue type %q", s)
	}
	// Set the issue type to the value found in the map per the key.
	*i = got
	return nil
}

// AnnounceTime refers to the time of earnings announcement.
type AnnounceTime int

const (
	bto AnnounceTime = iota
	dmt
	amc
)

var announceTimeDescription = map[AnnounceTime]string{
	bto: "Before open",
	dmt: "During trading",
	amc: "After close",
}

// AnnounceTimes maps the string keys from the JSON to the AnnounceType
// constant values.
var AnnounceTimes = map[string]AnnounceTime{
	"BTO": bto,
	"DMT": dmt,
	"AMC": amc,
}

// AnnounceTimeJSON maps an AnnounceTime to the string used in the JSON.
var AnnounceTimeJSON = map[AnnounceTime]string{
	bto: "BTO",
	dmt: "DMT",
	amc: "AMC",
}

// UnmarshalJSON implements the Unmarshaler interface for AnnounceTime.
func (a *AnnounceTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("announceTime should be a string, got %s", data)
	}
	return a.Set(s)
}

// Set sets the issue type using a string.
func (a *AnnounceTime) Set(s string) error {
	// Ensure the provided string matches on the keys in the map.
	got, ok := AnnounceTimes[s]
	if !ok {
		return fmt.Errorf("invalid issue type %q", s)
	}
	// Set the issue type to the value found in the map per the key.
	*a = got
	return nil
}

// MarshalJSON implements the Marshaler interface for AnnounceTime.
func (a *AnnounceTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(AnnounceTimeJSON[*a])
}

// String implements the Stringer interface for AnnounceTime.
func (a AnnounceTime) String() string {
	return announceTimeDescription[a]
}
