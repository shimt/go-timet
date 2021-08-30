# go-timet

[![Go Reference](https://pkg.go.dev/badge/github.com/shimt/go-timet.svg)](https://pkg.go.dev/github.com/shimt/go-timet) [![GitHub Actions](https://github.com/shimt/go-timet/actions/workflows/main.yml/badge.svg)](https://github.com/shimt/go-timet/actions/workflows/main.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/shimt/go-timet)](https://goreportcard.com/report/github.com/shimt/go-timet) [![codecov](https://codecov.io/gh/shimt/go-timet/branch/master/graph/badge.svg?token=W60LVKYO9B)](https://codecov.io/gh/shimt/go-timet)

go-timet is time package template collection.

- I've coded many times.
- I've checked the specification many times.
- I've checked the implementation many times.
- I've made mistakes many times.

## Contents

### Elapsed time

A structure and methods for measuring elapsed time.

- Create an instance with ElapsedTime{}.
- Start() to start measurement.
- Stop() to end measurement.
- Get the elapsed time with ElapsedTime().

### Unit of duration

A set of functions for creating units of duration.

| unit             | function | equal to              |
| ---------------- | -------- | --------------------- |
| nanosecond (ns)  | NS(n)    | n \* time.Nanosecond  |
| microsecond (μs) | US(n)    | n \* time.Microsecond |
| millisecond (ms) | MS(n)    | n \* time.Millisecond |
| second (s)       | S(n)     | n \* time.Second      |
| minute (m)       | M(n)     | n \* time.Minute      |
| hour (h)         | H(n)     | n \* time.Hour        |
