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

	s := Station{ID: "f00d", Key: "deadbeef"}

	wx := &Wx{}
	wx.Bar(29.86)
	a.Equal("https://stations.windy.com/pws/update/deadbeef?pressure=1011.171652636", s.Encode(wx))

	wx.OutTemp(20)
	a.Equal("https://stations.windy.com/pws/update/deadbeef?pressure=1011.171652636&temp=-6.666666666666667", s.Encode(wx))
}
