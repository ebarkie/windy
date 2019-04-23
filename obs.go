// Copyright (c) 2019 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package windy

import "time"

// Obs represents a weather observation.
type Obs struct {
	// Station ID
	Station int32 `json:"station"`

	// Observation time
	Time *time.Time `json:"time,omitempty"`

	// Air temperature in °C
	Temp *float64 `json:"temp,omitempty"`

	// Air temperature in °F
	TempF *float64 `json:"tempf,omitempty"`

	// Wind speed in m/s
	Wind *float64 `json:"wind,omitempty"`

	// Wind speed in MPH
	WindSpeedMPH *float64 `json:"windspeedmph,omitempty"`

	// Instantaneous wind direction in degrees
	WindDir *int `json:"winddir,omitempty"`

	// Current wind gust in m/s
	Gust *float64 `json:"gust,omitempty"`

	// Current wind gust in MPH
	WindGustMPH *float64 `json:"windgustmph,omitempty"`

	// Relative humidity
	RH *float64 `json:"rh,omitempty"`

	// Dew point in °C
	DewPoint *float64 `json:"dewpoint,omitempty"`

	// Atmospheric pressure in Pa
	Pressure *float64 `json:"pressure,omitempty"`

	// Atmospheric pressure in inches of Hg
	BaromIn *float64 `json:"baromin,omitempty"`

	// Precipitation over the past hour in mm
	Precip *float64 `json:"precip,omitempty"`

	// Precipitation over the past hour in in
	RainIn *float64 `json:"rainin,omitempty"`

	// UltraViolet light index
	UV *float64 `json:"uv,omitempty"`
}
