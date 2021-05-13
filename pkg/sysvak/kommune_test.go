package sysvak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMunicipality(t *testing.T) {
	assert.Equal(t, "Trondheim", MunicipalityByCode["5001"].Name)

	assert.Equal(t, "Trondheim", SearchMunicipality("Trondheim")[0].Name)
	assert.Greater(t, len(SearchMunicipality("ro")), 0)
}
