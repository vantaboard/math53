package helpers

func Smooth(min float64, max float64, n int) []float64 {
	step := (max - min) / float64(n)
	ts := make([]float64, n)
	for i := 0; i < n; i++ {
		ts[i] = min + float64(i)*step
	}
	return ts
}
