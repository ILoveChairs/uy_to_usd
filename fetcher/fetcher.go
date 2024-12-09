package fetcher

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func FetchConversionRate() (float64, error) {
	resp, err := http.Get("https://www.google.com/finance/quote/USD-UYU")
	if err != nil {
		return 0.0, errors.New("fetch was not possible")
	}
	if resp.StatusCode != 200 {
		return 0.0, errors.New("fetch status code is not 200")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, errors.New("failed to read response body")
	}

	convRateStr := extractConversionRate(string(body))
	convRate, err := strconv.ParseFloat(convRateStr, 64)
	if err != nil {
		return 0.0, errors.New("failed to parse conversion rate")
	}

	return convRate, nil
}

// <div class="YMlKec fxKbKc">43,4939</div>
func extractConversionRate(body string) string {
	const key string = "<div class=\"YMlKec fxKbKc\">"
	startIdx := strings.Index(body, key)
	endIdx := strings.Index(body[startIdx:], "</div>")
	return body[startIdx+len(key) : startIdx+endIdx]
}
