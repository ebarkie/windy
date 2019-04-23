// Copyright (c) 2019 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package windy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadEncode(t *testing.T) {
	a := assert.New(t)

	r := &Req{Key: "deadbeef"}
	a.Equal("https://stations.windy.com/pws/update/deadbeef", r.Encode())

	o := Obs{}
	baromin := 29.86
	o.BaromIn = &baromin
	r.Obss = []Obs{o}
	a.Equal(`{"observations":[{"station":0,"baromin":29.86}]}`, r.Body())

	tempf := 86.5
	o.TempF = &tempf
	r.Obss = []Obs{o}
	a.Equal(`{"observations":[{"station":0,"tempf":86.5,"baromin":29.86}]}`, r.Body())
}
