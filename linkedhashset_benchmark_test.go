package set

import "testing"

func BenchmarkLinkedHashSet_Contains_vs_InArray(b *testing.B) {
	set := NewLinkedHashSet[string]()
	set.Add(giantGenericSlice...)

	foundTarget := giantGenericSlice[len(giantGenericSlice)/2]
	notFoundTarget := "___not_present___"

	b.Run("Found", func(b *testing.B) {
		b.Run("Contains", func(b *testing.B) {
			b.ReportAllocs()
			var sink bool
			for i := 0; i < b.N; i++ {
				sink = set.Contains(foundTarget)
			}
			_ = sink
		})
		b.Run("InArray", func(b *testing.B) {
			b.ReportAllocs()
			var sink bool
			for i := 0; i < b.N; i++ {
				sink = set.InArray(foundTarget)
			}
			_ = sink
		})
	})

	b.Run("NotFound", func(b *testing.B) {
		b.Run("Contains", func(b *testing.B) {
			b.ReportAllocs()
			var sink bool
			for i := 0; i < b.N; i++ {
				sink = set.Contains(notFoundTarget)
			}
			_ = sink
		})
		b.Run("InArray", func(b *testing.B) {
			b.ReportAllocs()
			var sink bool
			for i := 0; i < b.N; i++ {
				sink = set.InArray(notFoundTarget)
			}
			_ = sink
		})
	})
}
