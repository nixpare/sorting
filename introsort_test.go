package sorting

import "testing"

func TestIntroSort(t *testing.T) {
	testSortingAlgorithmStandard[[]integer](t, IntroSort, newRandomInteger, nil)
}

func BenchmarkIntroSort(b *testing.B) {
	b.Run("Reduced", func(b *testing.B) {
		benchmarkSortingAlgorithmReduced[[]integer](b, IntroSort, newRandomInteger)
	})

	b.Run("Standard", func(b *testing.B) {
		benchmarkSortingAlgorithmStandard[[]integer](b, IntroSort, newRandomInteger)
	})

	b.Run("Shuffle", func(b *testing.B) {
		benchmarkSortingAlgorithmShuffle[[]integer](b, IntroSort, inOrderInteger, testSize)
	})
}