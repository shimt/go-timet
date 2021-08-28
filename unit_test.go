// Copyright 2021 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timet

import (
	"fmt"
	"testing"
	"time"
)

func ExampleNS() {
	d := NS(1)
	fmt.Println(d)
	fmt.Println(d == 1*time.Nanosecond)
	// Output:
	// 1ns
	// true
}

func TestNS(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"0", args{0}, 0 * time.Nanosecond},
		{"1", args{1}, 1 * time.Nanosecond},
		{"10", args{10}, 10 * time.Nanosecond},
		{"100", args{100}, 100 * time.Nanosecond},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NS(tt.args.n); got != tt.want {
				t.Errorf("NS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleUS() {
	d := US(1)
	fmt.Println(d)
	fmt.Println(d == 1*time.Microsecond)
	// Output:
	// 1µs
	// true
}

func TestUS(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"0", args{0}, 0 * time.Microsecond},
		{"1", args{1}, 1 * time.Microsecond},
		{"10", args{10}, 10 * time.Microsecond},
		{"100", args{100}, 100 * time.Microsecond},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := US(tt.args.n); got != tt.want {
				t.Errorf("US() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleMS() {
	d := MS(1)
	fmt.Println(d)
	fmt.Println(d == 1*time.Millisecond)
	// Output:
	// 1ms
	// true
}

func TestMS(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"0", args{0}, 0 * time.Millisecond},
		{"1", args{1}, 1 * time.Millisecond},
		{"10", args{10}, 10 * time.Millisecond},
		{"100", args{100}, 100 * time.Millisecond},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MS(tt.args.n); got != tt.want {
				t.Errorf("MS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleS() {
	d := S(1)
	fmt.Println(d)
	fmt.Println(d == 1*time.Second)
	// Output:
	// 1s
	// true
}

func TestS(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"0", args{0}, 0 * time.Second},
		{"1", args{1}, 1 * time.Second},
		{"10", args{10}, 10 * time.Second},
		{"100", args{100}, 100 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := S(tt.args.n); got != tt.want {
				t.Errorf("S() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleM() {
	d := M(1)
	fmt.Println(d)
	fmt.Println(d == 1*time.Minute)
	// Output:
	// 1m0s
	// true
}

func TestM(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"0", args{0}, 0 * time.Minute},
		{"1", args{1}, 1 * time.Minute},
		{"10", args{10}, 10 * time.Minute},
		{"100", args{100}, 100 * time.Minute},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := M(tt.args.n); got != tt.want {
				t.Errorf("M() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleH() {
	d := H(1)
	fmt.Println(d)
	fmt.Println(d == 1*time.Hour)
	// Output:
	// 1h0m0s
	// true
}

func TestH(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"0", args{0}, 0 * time.Hour},
		{"1", args{1}, 1 * time.Hour},
		{"10", args{10}, 10 * time.Hour},
		{"100", args{100}, 100 * time.Hour},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := H(tt.args.n); got != tt.want {
				t.Errorf("H() = %v, want %v", got, tt.want)
			}
		})
	}
}
