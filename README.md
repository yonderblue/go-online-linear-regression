#go-online-linear-regression
[![Build Status](https://travis-ci.org/gaillard/go-online-linear-regression.png)](https://travis-ci.org/gaillard/go-online-linear-regression)
[![GoDoc](https://godoc.org/github.com/gaillard/go-online-linear-regression/v1?status.svg)](https://godoc.org/github.com/gaillard/go-online-linear-regression/v1)

Library to calculate online linear regression.

##Example
```go
package main

import (
	"fmt"
	"github.com/gaillard/go-online-linear-regression/v1"
)

func main() {
	//xDelta argument of 7 will cause the point added with r.Add(1.5, 4.4) to be removed before last fmt.Printf()
	r := regression.New(7)

	r.Add(1.5, 4.4)
	r.Add(2.9, 1.56)

	slope, intercept, stdError := r.CalculateWithStdError()
	fmt.Printf("slope %f\n", slope)
	fmt.Printf("intercept %f\n", intercept)
	fmt.Printf("standard error %f\n", stdError)

	r.Add(7.2, 10.5)
	r.Add(9, 7.6)

	slope, intercept, stdError = r.CalculateWithStdError()
	fmt.Printf("slope %f\n", slope)
	fmt.Printf("intercept %f\n", intercept)
	fmt.Printf("standard error %f\n", stdError)
}
```

Outputs
```shell
slope -2.028571
intercept 7.442857
standard error NaN
slope 1.188768
intercept -1.015158
standard error 3.720348
```
