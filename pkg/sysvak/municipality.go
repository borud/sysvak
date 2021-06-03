package sysvak

import (
	_ "embed"
	"encoding/csv"
	"strings"

	"io"
	"log"
)

// MunicipalityCode is the 4 digit municipality code for Norwegian municipalities
type MunicipalityCode string

// Municipality ...
type Municipality struct {
	Code string
	Name string
}

//go:embed "kommuner.csv"
var municipalities string

// MunicipalityByCode maps from Norwegian municipality codes to Municipality struct
var MunicipalityByCode = make(map[string]Municipality)

func init() {
	// This is supposedly ISO 8859-1 according to SSB ... but it isn't
	r := csv.NewReader(strings.NewReader(municipalities))
	r.Comma = ';'

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error parsing CSV file: %v", err)
		}

		code := rec[0]
		name := rec[1]

		MunicipalityByCode[code] = Municipality{
			Code: code,
			Name: name,
		}
	}
}

// SearchMunicipality returns []Municipality for which s matches subscrint of Name
func SearchMunicipality(s string) []Municipality {
	var matches []Municipality

	search := strings.ToLower(s)

	for _, v := range MunicipalityByCode {
		if strings.Contains(strings.ToLower(v.Name), search) {
			matches = append(matches, v)
		}
	}
	return matches
}
