// Copyright 2021 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timet

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestElapsedTime_Start(t *testing.T) {
	et := &ElapsedTime{}

	et.Start()
	assert.InDelta(t, time.Until(et.st), 0, float64(1*time.Millisecond))
}

func TestElapsedTime_Stop(t *testing.T) {
	et := &ElapsedTime{}

	et.Stop()
	assert.InDelta(t, time.Until(et.et), 0, float64(1*time.Millisecond))
}

func TestElapsedTime_Time(t *testing.T) {
	et := &ElapsedTime{}
	et.Start()
	et.Stop()

	s, e := et.Time()

	assert.Equal(t, et.st, s)
	assert.Equal(t, et.et, e)
}

func TestElapsedTime_Elapsed(t *testing.T) {
	et := &ElapsedTime{}
	et.Start()
	et.Stop()

	assert.Equal(t, et.et.Sub(et.st), et.Elapsed())
}
