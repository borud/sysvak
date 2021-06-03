package sysvak

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://statistikk.fhi.no/api/Sysvak/gruppering"

// Lookup ...
func Lookup(q Query) ([]Result, error) {
	url, err := q.AsURL()
	if err != nil {
		return []Result{}, err
	}
	r, err := http.Get(url)
	if err != nil {
		return []Result{}, err
	}
	defer r.Body.Close()

	var rawResult []RawResult
	err = json.NewDecoder(r.Body).Decode(&rawResult)
	if err != nil {
		return nil, err
	}

	var result = make([]Result, len(rawResult))
	for n, rr := range rawResult {
		result[n] = rr.Parse()
	}

	return result, nil
}
