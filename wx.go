// Copyright (c) 2019 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package windy

import (
	"github.com/ebarkie/http/query"
	"github.com/ebarkie/weatherlink/units"
)

// Wx represents weather observations.
type Wx struct {
	query.Data
}

// Bar is barometric pressure in inches.
func (w *Wx) Bar(in float64) {
	w.SetFloat("pressure", units.Pressure(in*units.Inches).Hectopascals())
}

// DewPoint is the outdoor dew point in degrees Fahrenheit.
func (w *Wx) DewPoint(f float64) {
	w.SetFloat("dewpoint", units.Fahrenheit(f).Celsius())
}

// OutHumidity is the humidity percentage (0-100).
func (w *Wx) OutHumidity(p int) {
	w.SetInt("rh", p)
}

// OutTemp is outdoor temperature in degrees Fahrenheit.
func (w *Wx) OutTemp(f float64) {
	w.SetFloat("temp", units.Fahrenheit(f).Celsius())
}

// RainRate is rain inches over the past hour or the accumulated
// rainfall for the past 60 minutes in inches.
func (w *Wx) RainRate(in float64) {
	w.SetFloat("precip", units.Length(in*units.Inches).Millimeters())
}

// UVIndex is the UltraViolet light index.
func (w *Wx) UVIndex(i float64) {
	w.SetFloat("uv", i)
}

// WindDir is instantaneous wind direction in degrees (0-359).
func (w *Wx) WindDir(deg int) {
	w.SetInt("winddir", deg)
}

// WindGustSpeed is the software specific time period wind gust
// speed in miles per hour.
func (w *Wx) WindGustSpeed(mph float64) {
	w.SetFloat("gust", units.Speed(mph*units.MPH).MPS())
}

// WindSpeed is the instantaneous wind speed in miles per hour.
func (w *Wx) WindSpeed(mph float64) {
	w.SetFloat("wind", units.Speed(mph*units.MPH).MPS())
}
