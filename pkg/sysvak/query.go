package sysvak

import (
	"fmt"
	"strconv"
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
func (q Query) AsURL() (string, error) {

	// This is a bit silly because the API has some warts
	doses := make([]string, len(q.Doses))
	for n, d := range q.Doses {
		intVal, err := strconv.ParseInt(d, 10, 64)
		if err != nil {
			return "", err
		}
		doses[n] = fmt.Sprintf("%02d", intVal)
	}

	return baseURL + "?" +
		strings.Join([]string{
			"tabell=diagnose",
			"fordeling=geografi",
			"diagnoseKodeListe=COVID_19",
			fmt.Sprintf("fraDag=%s", q.From.Format("2006-01-02")),
			fmt.Sprintf("tilDag=%s", q.To.Format("2006-01-02")),
			fmt.Sprintf("doseKodeListe=%s", strings.Join(doses, ",")),
			fmt.Sprintf("kommuneKodeListe=%s", strings.Join(q.Municipalities, ",")),
			fmt.Sprintf("kjonnKodeListe=%s", strings.Join(q.Genders, ",")),
			fmt.Sprintf("aldersgruppeKodeListe=%s", strings.Join(q.Ages, ",")),
		}, "&"), nil
}
