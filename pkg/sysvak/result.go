package sysvak

import (
	"strings"
	"time"
)

// RawResult ...
type RawResult struct {
	Description string    `json:"tekst"`
	Count       int64     `json:"antall"`
	Where       string    `json:"fordeltPaa,omitempty"`
	Order       int64     `json:"rekkefolge,omitempty"`
	Date        time.Time `json:"provedato,omitempty"`
	Month       int64     `json:"manedNr,omitempty"`
}

// Result ...
type Result struct {
	Description string
	Gender      int
	Dose        int
	Age         int
	Where       string
	Date        time.Time
	Order       int64
	Month       int64
	Count       int64
}

// Parse and return normalized result
func (r *RawResult) Parse() Result {
	fields := strings.Split(r.Description, ", ")

	result := Result{}
	result.Description = r.Description
	result.Dose = StringToDose[fields[1]]
	result.Gender = StringToGender[fields[2]]
	result.Age = StringToAgeRange[fields[3]]
	result.Where = r.Where
	result.Date = r.Date
	result.Order = r.Order
	result.Month = r.Month
	result.Count = r.Count

	return result
}
