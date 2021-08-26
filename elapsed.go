// Copyright 2021 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timet

import "time"

// ElapsedTime represents an active elapsed time  instance.
type ElapsedTime struct {
	st time.Time
	et time.Time
}

// Start starts measure.
func (s *ElapsedTime) Start() {
	s.st = time.Now()
}

// Start ends measure.
func (s *ElapsedTime) Stop() {
	s.et = time.Now()
}

// Time returns start time and end time.
func (s *ElapsedTime) Time() (time.Time, time.Time) {
	return s.st, s.et
}

// Elapsed returns elapsed time.
func (s *ElapsedTime) Elapsed() time.Duration {
	if s.et.After(time.Time{}) {
		return s.et.Sub(s.st)
	}
	return time.Until(s.st)
}

// String returns a string representing the elapsed time.
func (s *ElapsedTime) String() string {
	return s.Elapsed().String()
}
