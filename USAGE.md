# windy

```go
import "github.com/ebarkie/windy"
```

Package windy implements the Windy Stations upload protocol.

## Usage

```go
var URL = "https://stations.windy.com/pws/update"
```
URL is the base API URL.

#### type Station

```go
type Station struct {
	ID   string    // Station name
	Key  string    // Station key
	Time time.Time // Upload time (default is now)
}
```

Station represents a weather station.

#### func (Station) Encode

```go
func (s Station) Encode(obs ...query.Values) string
```
Encode returns the request URL for the specified observations. This is generally
used for testing and debugging.

#### func (*Station) Upload

```go
func (s *Station) Upload(obs ...query.Values) (err error)
```
Upload uploads the specified observations.

#### type Wx

```go
type Wx struct {
	query.Data
}
```

Wx represents weather observations.

#### func (*Wx) Bar

```go
func (w *Wx) Bar(in float64)
```
Bar is barometric pressure in inches.

#### func (*Wx) DewPoint

```go
func (w *Wx) DewPoint(f float64)
```
DewPoint is the outdoor dew point in degrees Fahrenheit.

#### func (*Wx) OutHumidity

```go
func (w *Wx) OutHumidity(p int)
```
OutHumidity is the humidity percentage (0-100).

#### func (*Wx) OutTemp

```go
func (w *Wx) OutTemp(f float64)
```
OutTemp is outdoor temperature in degrees Fahrenheit.

#### func (*Wx) RainRate

```go
func (w *Wx) RainRate(in float64)
```
RainRate is rain inches over the past hour or the accumulated rainfall for the
past 60 minutes in inches.

#### func (*Wx) UVIndex

```go
func (w *Wx) UVIndex(i float64)
```
UVIndex is the UltraViolet light index.

#### func (*Wx) WindDir

```go
func (w *Wx) WindDir(deg int)
```
WindDir is instantaneous wind direction in degrees (0-359).

#### func (*Wx) WindGustSpeed

```go
func (w *Wx) WindGustSpeed(mph float64)
```
WindGustSpeed is the software specific time period wind gust speed in miles per
hour.

#### func (*Wx) WindSpeed

```go
func (w *Wx) WindSpeed(mph float64)
```
WindSpeed is the instantaneous wind speed in miles per hour.
