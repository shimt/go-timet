// Copyright 2021 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timet

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}

	return t
}

func TestElapsedTime_Start(t *testing.T) {
	et := &ElapsedTime{}

	et.Start()
	assert.InDelta(t, time.Until(et.st), 0, float64(MS(1)))
}

func TestElapsedTime_Stop(t *testing.T) {
	et := &ElapsedTime{}

	et.Stop()
	assert.InDelta(t, time.Until(et.et), 0, float64(MS(1)))
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
	t.Run("start-top", func(t *testing.T) {
		et := ElapsedTime{}
		et.Start()
		et.Stop()
		assert.Equal(t, et.et.Sub(et.st), et.Elapsed())
	})

	t.Run("start-only", func(t *testing.T) {
		et := ElapsedTime{}
		et.Start()
		assert.InDelta(t, time.Until(et.st), et.Elapsed(), float64(MS(1)))
	})

	type fields struct {
		st time.Time
		et time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{"<empty>", fields{}, S(0)},
		{"1s", fields{
			st: newTime("2021-08-27T23:23:23+09:00"),
			et: newTime("2021-08-27T23:23:24+09:00"),
		}, 1 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ElapsedTime{
				st: tt.fields.st,
				et: tt.fields.et,
			}
			if got := s.Elapsed(); got != tt.want {
				t.Errorf("ElapsedTime.Elapsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElapsedTime_String(t *testing.T) {
	type fields struct {
		st time.Time
		et time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"<empty>", fields{}, "0s"},
		{"1s", fields{
			st: newTime("2021-08-27T23:23:23+09:00"),
			et: newTime("2021-08-27T23:23:24+09:00"),
		}, "1s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ElapsedTime{
				st: tt.fields.st,
				et: tt.fields.et,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("ElapsedTime.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
