// Copyright (c) 2019 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package windy

// Refer to:
// https://community.windy.com/topic/8168/report-you-weather-station-data-to-windy

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ebarkie/http/query"
)

// createRequest builds the HTTP request.
func (s Station) createRequest(obs ...query.Values) *http.Request {
	req, _ := http.NewRequest("GET", URL+"/"+s.Key, nil)

	// Create mandatory query parameters.
	q := req.URL.Query()
	//q.Add("station", s.ID)
	if !s.Time.IsZero() {
		q.Add("time", s.Time.In(time.UTC).Format(time.RFC3339))
	}

	// Add observations to query parameters.
	for _, o := range obs {
		for k, v := range o.Values() {
			q.Add(k, v)
		}
	}

	req.URL.RawQuery = q.Encode()

	return req
}

// Encode returns the request URL for the specified observations.  This
// is generally used for testing and debugging.
func (s Station) Encode(obs ...query.Values) string {
	return s.createRequest(obs...).URL.String()
}

// Upload uploads the specified observations.
func (s *Station) Upload(obs ...query.Values) (err error) {
	// Clear payload(s) after upload attempt.
	defer func() {
		for _, o := range obs {
			o.Clear()
		}
	}()

	// Initiate HTTP request.
	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(s.createRequest(obs...))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check HTTP return status code.
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP request returned non-OK status code %d", resp.StatusCode)
		return
	}

	return
}
