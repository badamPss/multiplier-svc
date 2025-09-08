package main

import "math/rand"

func generateMultiplier(rtp float64) float64 {
	r := rand.Float64()
	multiplier := rtp / (1 - r)

	if multiplier > 10000 {
		return 10000
	}
	if multiplier < 1 {
		return 1
	}

	return multiplier
}
