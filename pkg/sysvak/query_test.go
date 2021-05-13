package sysvak

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToURL(t *testing.T) {
	assert.Equal(t,
		"https://statistikk.fhi.no/api/Sysvak/gruppering?"+
			"tabell=diagnose&"+
			"fordeling=geografi&"+
			"diagnoseKodeListe=COVID_19&"+
			"fraDag=2021-01-01&"+
			"tilDag=2021-01-07&"+
			"doseKodeListe=01,02&"+
			"kommuneKodeListe=5001&"+
			"kjonnKodeListe=K,M&"+
			"aldersgruppeKodeListe=1,2,3,4,5,6,7",
		Query{
			From:           time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
			To:             time.Date(2021, 1, 7, 0, 0, 0, 0, time.Local),
			Doses:          []string{"01,02"},
			Municipalities: []string{"5001"},
			Genders:        []string{"K", "M"},
			Ages:           []string{"1", "2", "3", "4", "5", "6", "7"},
		}.AsURL())
}
