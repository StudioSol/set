package set

import (
	"fmt"
	"testing"
)

func BenchmarkLinkedHashSet_Contains_vs_InArray(b *testing.B) {
	total := len(giantGenericSlice)
	step := total / 5
	sizes := []int{step, 2 * step, 3 * step, 4 * step, total}

	for _, n := range sizes {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			set := NewLinkedHashSet[string]()
			set.Add(giantGenericSlice[:n]...)

			foundTarget := giantGenericSlice[n/2]
			notFoundTarget := "___not_present___"

			b.Run("Found/Contains", func(b *testing.B) {
				b.ReportAllocs()
				var sink bool
				for i := 0; i < b.N; i++ {
					sink = set.Contains(foundTarget)
				}
				_ = sink
			})
			b.Run("Found/InArray", func(b *testing.B) {
				b.ReportAllocs()
				var sink bool
				for i := 0; i < b.N; i++ {
					sink = set.InArray(foundTarget)
				}
				_ = sink
			})
			b.Run("NotFound/Contains", func(b *testing.B) {
				b.ReportAllocs()
				var sink bool
				for i := 0; i < b.N; i++ {
					sink = set.Contains(notFoundTarget)
				}
				_ = sink
			})
			b.Run("NotFound/InArray", func(b *testing.B) {
				b.ReportAllocs()
				var sink bool
				for i := 0; i < b.N; i++ {
					sink = set.InArray(notFoundTarget)
				}
				_ = sink
			})
		})
	}
}
