package utils

import (
	"math/rand"
)

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
