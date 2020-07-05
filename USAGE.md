# windy

```go
import "github.com/ebarkie/windy"
```

Package windy implements the Windy Stations protocol.

## Usage

```go
var URL = "https://stations.windy.com/pws/update"
```
URL is the base API URL.

#### type Obs

```go
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
```

Obs represents a weather observation.

#### type Req

```go
type Req struct {
	Key      string    `json:"-"`
	Stations []Station `json:"stations,omitempty"`
	Obss     []Obs     `json:"observations,omitempty"`
}
```

Req is a Windy API request.

#### func (Req) Body

```go
func (r Req) Body() string
```
Body returns the request Body.

#### func (Req) Encode

```go
func (r Req) Encode() string
```
Encode returns the request URL.

#### func (*Req) Upload

```go
func (r *Req) Upload() (err error)
```
Upload uploads to Windy.

#### type ShareOpt

```go
type ShareOpt uint8
```

ShareOpt is a share option.

```go
const (
	Open ShareOpt = iota
	OpenWindy
	Private
)
```
Possible share options.

#### func (ShareOpt) MarshalText

```go
func (o ShareOpt) MarshalText() (text []byte, err error)
```
MarshalText marshals the share option as text.

#### func (ShareOpt) String

```go
func (o ShareOpt) String() string
```
String returns the share option as a string.

#### type Station

```go
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
```

Station represents a weather station.
