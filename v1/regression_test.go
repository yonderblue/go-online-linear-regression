package regression

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Example() {
	r := New(math.MaxFloat64)

	r.Add(1.5, 4.4)
	r.Add(2.9, 1.56)

	slope, intercept := r.Calculate()
	fmt.Printf("slope %f\n", slope)
	fmt.Printf("intercept %f\n", intercept)

	r.Add(7.2, 10.5)
	r.Add(9, 7.6)

	slope, intercept = r.Calculate()
	fmt.Printf("slope %f\n", slope)
	fmt.Printf("intercept %f\n ", intercept)

	//Output:
	//slope -2.028571
	//intercept 7.442857
	//slope 0.834403
	//intercept 1.717827
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	slope, intercept := r.Calculate()
	if !math.IsNaN(slope) || !math.IsNaN(intercept) {
		t.Error(slope, intercept)
	}
}

func TestSingle(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	r.Add(1.5, 4.4)
	slope, intercept := r.Calculate()
	if !math.IsNaN(slope) || !math.IsNaN(intercept) {
		t.Error(slope, intercept)
	}
}

func TestZeroDuration(t *testing.T) {
	t.Parallel()

	r := New(0)
	r.Add(0, 1)
	slope, intercept := r.Calculate()
	if !math.IsNaN(slope) || !math.IsNaN(intercept) {
		t.Error(slope, intercept)
	}
}

func TestRemoval(t *testing.T) {
	t.Parallel()

	r := New(1)
	r.Add(1, 1)
	r.Add(2, 2)
	slope, intercept := r.Calculate()
	if floatToString(slope, 1) != "1.0" || floatToString(intercept, 1) != "0.0" {
		t.Error(slope, intercept)
	}

	//x 1 should be removed on this Add()
	r.Add(3, 1)
	slope, intercept = r.Calculate()
	if floatToString(slope, 1) != "-1.0" || floatToString(intercept, 1) != "4.0" {
		t.Error(slope, intercept)
	}

	//same time, no removal
	r.Add(3, 1)
	slope, intercept = r.Calculate()
	if floatToString(slope, 1) != "-1.0" || floatToString(intercept, 1) != "4.0" {
		t.Error(slope, intercept)
	}
}

func TestRepeatX(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	r.Add(1, 1)
	r.Add(1, 2)
	r.Add(2, 3)
	r.Add(2, 4)
	slope, intercept := r.Calculate()
	if floatToString(slope, 1) != "2.0" || floatToString(intercept, 1) != "-0.5" {
		t.Error(slope, intercept)
	}
}

func TestRepeatY(t *testing.T) {
	t.Parallel()

	r := New(math.MaxFloat64)
	r.Add(1, 1)
	r.Add(2, 1)
	r.Add(3, 2)
	r.Add(4, 2)
	slope, intercept := r.Calculate()
	if floatToString(slope, 1) != "0.4" || floatToString(intercept, 1) != "0.5" {
		t.Error(slope, intercept)
	}
}

func floatToString(float float64, digitsAfterDecimal int) string {
	return strconv.FormatFloat(float, 'f', digitsAfterDecimal, 64)
}
