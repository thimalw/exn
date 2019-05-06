// Package exchange provides a Go API for currency conversion using openexchangerates.org
package exchange

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const appID = "116da9ad934a4066b3883abbd60f8fac"
const baseURL = "https://openexchangerates.org/api/"

type latestResult struct {
	Rates map[string]float64
}

// Convert returns the amount in `value` converted from currency of `from` to
// the currency of `to`.
// `from` and `to` must be 3-letter currency codes.
func Convert(value float64, from, to string) (float64, error) {
	q := "latest.json?app_id=" + appID
	url := baseURL + q
	resp, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("Convert: %s", resp.Status)
	}

	var result latestResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0.0, err
	}

	rate := (result.Rates[to] / result.Rates[from]) * value

	return rate, nil
}
