package sysvak

import (
	"fmt"
	"strings"
	"time"
)

// Query ...
type Query struct {
	From           time.Time
	To             time.Time
	Doses          []string
	Municipalities []string
	Genders        []string
	Ages           []string
}

// AgeRange ...
var AgeRange = map[int]string{
	1: "0-15",
	2: "16-44",
	3: "45-54",
	4: "55-64",
	5: "65-74",
	6: "75-84",
	7: "> 85",
}

// NewQuery creates a new query with reasonable defaults
func NewQuery() Query {
	return Query{
		From:           time.Now().Add(-24 * 7 * time.Hour),
		To:             time.Now(),
		Doses:          []string{"1", "2"},
		Municipalities: []string{},
		Genders:        []string{"K", "M"},
		Ages:           []string{"1", "2", "3", "4", "5", "6", "7"},
	}
}

// AsURL returns query as URL
func (q Query) AsURL() string {

	return baseURL + "?" +
		strings.Join([]string{
			"tabell=diagnose",
			"fordeling=geografi",
			"diagnoseKodeListe=COVID_19",
			fmt.Sprintf("fraDag=%s", q.From.Format("2006-01-02")),
			fmt.Sprintf("tilDag=%s", q.To.Format("2006-01-02")),
			fmt.Sprintf("doseKodeListe=%s", strings.Join(q.Doses, ",")),
			fmt.Sprintf("kommuneKodeListe=%s", strings.Join(q.Municipalities, ",")),
			fmt.Sprintf("kjonnKodeListe=%s", strings.Join(q.Genders, ",")),
			fmt.Sprintf("aldersgruppeKodeListe=%s", strings.Join(q.Ages, ",")),
		}, "&")
}
