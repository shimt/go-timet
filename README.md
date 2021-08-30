# go-timet
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fshimt%2Fgo-timet.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fshimt%2Fgo-timet?ref=badge_shield)


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


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fshimt%2Fgo-timet.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fshimt%2Fgo-timet?ref=badge_large)