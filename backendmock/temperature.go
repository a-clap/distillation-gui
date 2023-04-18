package backendmock

import "math/rand"

func randomTemperature(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
