package randomtools

// PrepareProbability normalize value to 0.0..1.0
func PrepareProbability(in float64) float64 {
	if in < 0 {
		return 0
	}
	if in > 1 {
		return 1
	}

	return in
}
