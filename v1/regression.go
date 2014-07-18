//Package regression provides online linear regression calculation.
package regression

import (
	"github.com/gaillard/go-queue/queue"
	"math"
)

//Regression represents a queue of past points. Use New() to initialize.
type Regression struct {
	xSum, ySum, xxSum, xySum, xDelta float64
	points                           *queue.Queue
	lastSlopeCalc, lastInterceptCalc float64

	//here so multiple calcs calls per add calls wont hurt performance
	lastCalcFresh bool

	//here for performance to avoid point.Back() calls
	lastX float64
}

type point struct {
	x, y, xy, xx float64
}

//New returns a Regression that keeps points back as far as xDelta from the last
//added point.
func New(xDelta float64) *Regression {
	return &Regression{xDelta: xDelta, points: queue.New(), lastX: math.Inf(-1)}
}

//Calculate returns the slope and intercept of a best fit line to the added
//points. Returns a cached value if called between adds.
func (r *Regression) Calculate() (slope, intercept float64) {
	if r.lastCalcFresh {
		slope = r.lastSlopeCalc
		intercept = r.lastInterceptCalc
		return
	}

	n := float64(r.points.Len())

	xSumOverN := r.xSum / n //here to only calc once for performance
	slope = (r.xySum - xSumOverN*r.ySum) / (r.xxSum - xSumOverN*r.xSum)
	intercept = (r.ySum - slope*r.xSum) / n

	r.lastSlopeCalc = slope
	r.lastInterceptCalc = intercept
	r.lastCalcFresh = true
	return
}

//Add adds the new x and y as a point into the queue. Will not clear old points
//until calculation. Panics if given an x value less than the last.
func (r *Regression) Add(x, y float64) {
	r.lastCalcFresh = false

	if x < r.lastX {
		panic("adding with x less than the last add is not allowed")
	}
	r.lastX = x

	//storing pointers instead of values only for performance
	newPoint := &point{x, y, x * y, x * x}
	r.points.PushBack(newPoint)
	r.xSum += newPoint.x
	r.ySum += newPoint.y
	r.xxSum += newPoint.xx
	r.xySum += newPoint.xy

	//here to only calc once for performance
	oldestXAllowed := r.lastX - r.xDelta

	for {
		pointGeneric := r.points.Front()
		if pointGeneric == nil {
			break
		}

		point := pointGeneric.(*point)
		if point.x >= oldestXAllowed {
			break
		}

		r.xSum -= point.x
		r.ySum -= point.y
		r.xxSum -= point.xx
		r.xySum -= point.xy

		r.points.PopFront()
	}
}
