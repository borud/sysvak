package sysvak

import "time"

// Result ...
type Result struct {
	Description string    `json:"tekst"`
	Count       int64     `json:"antall"`
	Where       string    `json:"fordeltPaa,omitempty"`
	Order       int64     `json:"rekkefolge,omitempty"`
	Date        time.Time `json:"provedato,omitempty"`
	Month       int64     `json:"manedNr,omitempty"`
}
