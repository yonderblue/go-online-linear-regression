package regression

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Example() {
	r := New(7)

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

	//Output:
	//slope -2.028571
	//intercept 7.442857
	//standard error NaN
	//slope 1.188768
	//intercept -1.015158
	//standard error 3.720348
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	slope, intercept, stdError := r.CalculateWithStdError()
	if !math.IsNaN(slope) || !math.IsNaN(intercept) || !math.IsNaN(stdError) {
		t.Errorf("slope, intercept, stdError wasn't NaN, was %v, %v, %v", slope, intercept, stdError)
	}
}

func TestSingle(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	r.Add(1.5, 4.4)
	slope, intercept, stdError := r.CalculateWithStdError()
	if !math.IsNaN(slope) || !math.IsNaN(intercept) || !math.IsNaN(stdError) {
		t.Errorf("slope, intercept, stdError wasn't NaN, was %v, %v, %v", slope, intercept, stdError)
	}
}

func TestZeroDuration(t *testing.T) {
	t.Parallel()

	r := New(0)
	r.Add(0, 1)
	slope, intercept, stdError := r.CalculateWithStdError()
	if !math.IsNaN(slope) || !math.IsNaN(intercept) || !math.IsNaN(stdError) {
		t.Errorf("slope, intercept, stdError wasn't NaN, was %v, %v, %v", slope, intercept, stdError)
	}
}

func TestRemoval(t *testing.T) {
	t.Parallel()

	r := New(1)
	r.Add(1, 1)
	r.Add(2, 2)
	slope, intercept, stdError := r.CalculateWithStdError()
	if floatToString(slope, 1) != "1.0" || floatToString(intercept, 1) != "0.0" || !math.IsNaN(stdError) {
		t.Errorf("slope, intercept, stdError wasn't 1.0, 0.0, NaN, was %v, %v, %v", slope, intercept, stdError)
	}

	//x 1 should be removed on this Add()
	r.Add(3, 1)
	slope, intercept, stdError = r.CalculateWithStdError()
	if floatToString(slope, 1) != "-1.0" || floatToString(intercept, 1) != "4.0" || !math.IsNaN(stdError) {
		t.Errorf("slope, intercept, stdError wasn't -1.0, 4.0, NaN, was %v, %v, %v", slope, intercept, stdError)
	}

	//same time, no removal
	r.Add(3, 1)
	slope, intercept, stdError = r.CalculateWithStdError()
	if floatToString(slope, 1) != "-1.0" || floatToString(intercept, 1) != "4.0" || floatToString(stdError, 1) != "0.0" {
		t.Errorf("slope, intercept, stdError wasn't -1.0, 4.0, 0, was %v, %v, %v", slope, intercept, stdError)
	}
}

func TestRepeatX(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	r.Add(1, 1)
	r.Add(1, 2)
	r.Add(2, 3)
	r.Add(2, 4)
	slope, intercept, stdError := r.CalculateWithStdError()
	if floatToString(slope, 1) != "2.0" || floatToString(intercept, 1) != "-0.5" || floatToString(stdError, 1) != "0.7" {
		t.Errorf("slope, intercept, stdError wasn't 2.0, -0.5, 0.7, was %v, %v, %v", slope, intercept, stdError)
	}
}

func TestRepeatY(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	r.Add(1, 1)
	r.Add(2, 1)
	r.Add(3, 2)
	r.Add(4, 2)
	slope, intercept, stdError := r.CalculateWithStdError()
	if floatToString(slope, 1) != "0.4" || floatToString(intercept, 1) != "0.5" || floatToString(stdError, 1) != "0.3" {
		t.Errorf("slope, intercept, stdError wasn't 0.4, -0.5, 0.3, was %v, %v, %v", slope, intercept, stdError)
	}
}

func TestAddOld(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	r.Add(1, 1)

	defer func() {
		err := recover()
		if err == nil {
			t.Error("did not panic when Add()ing with x less than last x")
		}
	}()
	r.Add(0, 1)
}

func TestMultipleCalc(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	r.Add(1, 1)
	r.Add(2, 2)

	for i := 0; i < 3; i++ {
		slope, intercept, stdError := r.CalculateWithStdError()
		if floatToString(slope, 1) != "1.0" || floatToString(intercept, 1) != "0.0" || !math.IsNaN(stdError) {
			t.Errorf("slope, intercept wasn't 1.0, 0.0, NaN, was %v, %v, %v", slope, intercept, stdError)
		}
	}
}

func floatToString(float float64, digitsAfterDecimal int) string {
	return strconv.FormatFloat(float, 'f', digitsAfterDecimal, 64)
}
