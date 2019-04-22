// Copyright (c) 2019 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package windy implements the Windy Stations upload protocol.
package windy

import "time"

// URL is the base API URL.
var URL = "https://stations.windy.com/pws/update"

// Station represents a weather station.
type Station struct {
	ID   string    // Station name
	Key  string    // Station key
	Time time.Time // Upload time (default is now)
}
