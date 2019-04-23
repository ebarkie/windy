// Copyright (c) 2019 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package windy

// Refer to:
// https://community.windy.com/topic/8168/report-you-weather-station-data-to-windy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// createRequest builds the HTTP request.
func (r Req) createRequest() *http.Request {
	req, _ := http.NewRequest("POST", URL+"/"+r.Key, bytes.NewBufferString(r.Body()))
	req.Header.Set("Content-Type", "application/json")

	return req
}

// Encode returns the request URL.
func (r Req) Encode() string {
	return r.createRequest().URL.String()
}

// Body returns the request Body.
func (r Req) Body() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// Upload uploads to Windy.
func (r *Req) Upload() (err error) {
	// Clear observations after upload attempt.
	defer func() {
		r.Obss = []Obs{}
	}()

	// Initiate HTTP request.
	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(r.createRequest())
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
