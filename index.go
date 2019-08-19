package sorting

import (
	"math"
	"time"
)

func epochSeconds(date time.Time) int64 {
	return date.Unix()
}

func New(votes int32, created time.Time) float64 {
	s := float64(votes)
	createdseconds := epochSeconds(created)

	order := math.Log10(math.Max(math.Abs(s), 1))
	sign := 0

	if s > 0 {
		sign = 1
	} else if s < 0 {
		sign = -1
	}

	seconds := createdseconds - 1134028003

	return math.Round(float64(sign)*order + float64(seconds)/45000)
}

func Hot(votes int32, updated time.Time) float64 {
	s := float64(votes)
	updatedseconds := epochSeconds(updated)

	order := math.Log10(math.Max(math.Abs(s), 1))
	sign := 0

	if s > 0 {
		sign = 1
	} else if s < 0 {
		sign = -1
	}

	seconds := updatedseconds - 1134028003

	return math.Round(float64(sign)*order + float64(seconds)/45000)
}

func Engaging(votes int32, updated time.Time, entries int) float64 {
	s := float64(votes)
	updatedseconds := epochSeconds(updated)

	order := math.Log10(math.Max(math.Abs(s), 1)) * 10

	e := math.Log2(math.Max(math.Abs(float64(entries)), 1)) * 10

	sign := 0

	if s > 0 {
		sign = 1
	} else if s < 0 {
		sign = -1
	}

	seconds := updatedseconds - 1134028003

	return math.Round(float64(sign)*float64(e)*order + float64(seconds)/90000)
}
