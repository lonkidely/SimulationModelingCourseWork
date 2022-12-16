package utils

import (
	"math"
	"math/rand"
)

func Exponential(lambda float64) float64 {
	return rand.ExpFloat64() / lambda
}

func Normal(mu, sigma float64) float64 {
	return math.Abs(rand.NormFloat64()*sigma + mu)
}

func Uniform(l, r float64) float64 {
	return l + rand.Float64()*(r-l)
}
