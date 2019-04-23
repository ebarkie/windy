// Copyright (c) 2019 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package windy implements the Windy Stations protocol.
package windy

// URL is the base API URL.
var URL = "https://stations.windy.com/pws/update"

// Req is a Windy API request.
type Req struct {
	Key      string    `json:"-"`
	Stations []Station `json:"stations,omitempty"`
	Obss     []Obs     `json:"observations,omitempty"`
}

// Station represents a weather station.
type Station struct {
	// Station ID
	Station int32 `json:"station"`

	// Share option
	Share ShareOpt `json:"shareOption"`

	// User selected station name
	Name string `json:"name,omitempty"`

	// North-South position on Earth's surface
	Lat float64 `json:"latitude,omitempty"`

	// East-West position on Earth's surface
	Lon float64 `json:"longitude,omitempty"`

	// Height above the Earth's sea level
	Elev int `json:"elevation,omitempty"`

	// Temperature sensor height above the surface
	TempHeight int `json:"tempheight,omitempty"`

	// Wind sensor height above the surface
	WindHeight int `json:"windheight,omitempty"`
}

// ShareOpt is a share option.
type ShareOpt uint8

// Possible share options.
const (
	Open ShareOpt = iota
	OpenWindy
	Private
)

// MarshalText marshals the share option as text.
func (o ShareOpt) MarshalText() (text []byte, err error) {
	return []byte(o.String()), nil
}

// String returns the share option as a string.
func (o ShareOpt) String() string {
	opts := [...]string{
		"Open",
		"Only Windy",
		"Private",
	}

	if int(o) >= len(opts) {
		return opts[Private]
	}
	return opts[o]
}
