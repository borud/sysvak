package sysvak

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://statistikk.fhi.no/api/Sysvak/gruppering"

// Lookup ...
func Lookup(q Query) ([]Result, error) {
	r, err := http.Get(q.AsURL())
	if err != nil {
		return []Result{}, err
	}
	defer r.Body.Close()

	var result []Result
	return result, json.NewDecoder(r.Body).Decode(&result)
}
