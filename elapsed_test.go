// Copyright 2021 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timet

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		panic(err)
	}

	return t
}

func TestElapsedTime_Start(t *testing.T) {
	et := &ElapsedTime{}

	et.Start()
	assert.InDelta(t, time.Until(et.Begin), 0, float64(MS(1)))
}

func TestElapsedTime_Stop(t *testing.T) {
	et := &ElapsedTime{}

	et.Stop()
	assert.InDelta(t, time.Until(et.End), 0, float64(MS(1)))
}

func TestElapsedTime_Time(t *testing.T) {
	et := &ElapsedTime{}
	et.Start()
	et.Stop()

	s, e := et.Time()

	assert.Equal(t, et.Begin, s)
	assert.Equal(t, et.End, e)
}

func TestElapsedTime_Elapsed(t *testing.T) {
	t.Run("start-top", func(t *testing.T) {
		et := ElapsedTime{}
		et.Start()
		et.Stop()
		assert.Equal(t, et.End.Sub(et.Begin), et.Elapsed())
	})

	t.Run("start-only", func(t *testing.T) {
		et := ElapsedTime{}
		et.Start()
		assert.InDelta(t, time.Until(et.Begin), et.Elapsed(), float64(MS(1)))
	})

	type fields struct {
		begin time.Time
		end   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{"<empty>", fields{}, S(0)},
		{"1s", fields{
			begin: newTime("2021-08-27T23:23:23+09:00"),
			end:   newTime("2021-08-27T23:23:24+09:00"),
		}, 1 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ElapsedTime{
				Begin: tt.fields.begin,
				End:   tt.fields.end,
			}
			if got := s.Elapsed(); got != tt.want {
				t.Errorf("ElapsedTime.Elapsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElapsedTime_String(t *testing.T) {
	type fields struct {
		begin time.Time
		end   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"<empty>", fields{}, "0s"},
		{"1s", fields{
			begin: newTime("2021-08-27T23:23:23+09:00"),
			end:   newTime("2021-08-27T23:23:24+09:00"),
		}, "1s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ElapsedTime{
				Begin: tt.fields.begin,
				End:   tt.fields.end,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("ElapsedTime.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElapsedTime_MarshalJSON(t *testing.T) {
	type fields struct {
		begin time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"<empty>", fields{}, "{}", false},
		{"begin and end", fields{
			begin: newTime("2021-08-27T23:23:23.1+09:00"),
			end:   newTime("2021-08-27T23:23:24.2+09:00"),
		}, `{"begin":"2021-08-27T23:23:23.1+09:00","end":"2021-08-27T23:23:24.2+09:00"}`, false},
		{"begin only", fields{
			begin: newTime("2021-08-27T23:23:23.1+09:00"),
		}, `{"begin":"2021-08-27T23:23:23.1+09:00"}`, false},
		{"end only", fields{
			end: newTime("2021-08-27T23:23:24.2+09:00"),
		}, `{"end":"2021-08-27T23:23:24.2+09:00"}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ElapsedTime{
				Begin: tt.fields.begin,
				End:   tt.fields.end,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ElapsedTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, []byte(tt.want)) {
				t.Errorf("ElapsedTime.MarshalJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestElapsedTime_UnmarshalJSON(t *testing.T) {
	type fields struct {
		begin time.Time
		end   time.Time
	}
	tests := []struct {
		name       string
		json       string
		wantFields fields
		wantErr    bool
	}{
		{"<empty>", "{}", fields{time.Time{}, time.Time{}}, false},
		{"<error>", "[]", fields{time.Time{}, time.Time{}}, true},
		{"begin error", `{"begin": "A"}`, fields{time.Time{}, time.Time{}}, true},
		{"end error", `{"end": "B"}`, fields{time.Time{}, time.Time{}}, true},
		{"begin-end", `{"end":"2021-08-27T23:23:24.2+09:00","begin":"2021-08-27T23:23:23.1+09:00"}`, fields{
			begin: newTime("2021-08-27T23:23:23.1+09:00"),
			end:   newTime("2021-08-27T23:23:24.2+09:00"),
		}, false},
		{"begin only", `{"begin":"2021-08-27T23:23:23.1+09:00"}`, fields{
			begin: newTime("2021-08-27T23:23:23.1+09:00"),
		}, false},
		{"end only", `{"end":"2021-08-27T23:23:24.2+09:00"}`, fields{
			end: newTime("2021-08-27T23:23:24.2+09:00"),
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ElapsedTime{}
			if err := s.UnmarshalJSON([]byte(tt.json)); (err != nil) != tt.wantErr {
				t.Errorf("ElapsedTime.Unmarshaler() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantFields.begin, s.Begin, "check ElapsedTime.Begin")
			assert.Equal(t, tt.wantFields.end, s.End, "check ElapsedTime.End")
		})
	}
}
