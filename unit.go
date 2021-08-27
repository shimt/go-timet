// Copyright 2021 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timet

import "time"

// NS returns n nano second duration.
//
// This is the same as "n * time.Nanosecond".
func NS(n int64) time.Duration {
	return time.Duration(n) * time.Nanosecond
}

// US returns n micro second duration.
//
// This is the same as "n * time.Microsecond".
func US(n int64) time.Duration {
	return time.Duration(n) * time.Microsecond
}

// MS returns n milli second duration.
//
// This is the same as "n * time.Millisecond".
func MS(n int64) time.Duration {
	return time.Duration(n) * time.Millisecond
}

// S returns n second duration.
//
// This is the same as "n * time.Second".
func S(n int64) time.Duration {
	return time.Duration(n) * time.Second
}

// M returns n minute duration.
//
// This is the same as "n * time.Minute".
func M(n int64) time.Duration {
	return time.Duration(n) * time.Minute
}

// H returns n hour duration.
//
// This is the same as "n * time.Hour".
func H(n int64) time.Duration {
	return time.Duration(n) * time.Hour
}
