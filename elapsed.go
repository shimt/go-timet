// Copyright 2021 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timet

import (
	"encoding/json"
	"fmt"
	"time"
)

// ElapsedTime represents an active elapsed time  instance.
type ElapsedTime struct {
	Begin time.Time
	End   time.Time
}

// Start starts measure.
func (s *ElapsedTime) Start() {
	s.Begin = time.Now()
}

// Stop ends measure.
func (s *ElapsedTime) Stop() {
	s.End = time.Now()
}

// Time returns start time and end time.
func (s *ElapsedTime) Time() (time.Time, time.Time) {
	return s.Begin, s.End
}

// Elapsed returns elapsed time.
func (s *ElapsedTime) Elapsed() time.Duration {
	if s.Begin.Equal(s.End) {
		return time.Duration(0)
	}

	if s.End.After(time.Time{}) {
		return s.End.Sub(s.Begin)
	}
	return time.Until(s.Begin)
}

// String returns a string representing the elapsed time.
func (s *ElapsedTime) String() string {
	return s.Elapsed().String()
}

// ---- json ----

var _ json.Marshaler = &ElapsedTime{}
var _ json.Unmarshaler = &ElapsedTime{}

// MarshalJSON implements the json.Marshaler interface.
func (s *ElapsedTime) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 2)
	if s.Begin.After(time.Time{}) {
		m["start"] = s.Begin.Format(time.RFC3339Nano)
	}
	if s.End.After(time.Time{}) {
		m["stop"] = s.End.Format(time.RFC3339Nano)
	}
	return json.Marshal(m)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *ElapsedTime) UnmarshalJSON(b []byte) error {
	m := make(map[string]string, 2)
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	if v, ok := m["start"]; ok {
		t, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return fmt.Errorf("start: %w", err)
		}
		s.Begin = t
	}

	if v, ok := m["stop"]; ok {
		t, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return fmt.Errorf("stop: %w", err)
		}
		s.End = t
	}

	return nil
}
