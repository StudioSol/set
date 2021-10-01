package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const giantSliceLength = 100000

var giantINT64Slice = make([]int64, giantSliceLength)

func init() {
	for i := 0; i < giantSliceLength; i++ {
		giantINT64Slice[i] = int64(i + 1)
	}
}

func TestLinkedHashSetINT64Add(t *testing.T) {
	t.Run("Given LinkedHashSetINT64.Add", func(t *testing.T) {
		t.Run("It should not store elements that are already on the Set", func(t *testing.T) {
			set := NewLinkedHashSetINT64()
			set.Add(0, 0)
			set.Add(0)
			require.Equal(t, set.Length(), 1)
		})
		t.Run("It should store elements with the correct constraints", func(t *testing.T) {
			set := NewLinkedHashSetINT64()
			set.Add(0, 1, 2, 99, 93, 32, 00, 01, 2)
			require.Equal(t, set.Length(), 6)
		})
	})
}

func TestLinkedHashSetINT64Remove(t *testing.T) {
	t.Run("When a big list is given", func(t *testing.T) {
		set := NewLinkedHashSetINT64()
		set.Add(giantINT64Slice...)
		t.Run("It should remove elements from a Set", func(t *testing.T) {
			// first element
			set.Remove(giantINT64Slice[0])
			set.Remove(giantINT64Slice[0])
			set.Remove(giantINT64Slice[0])
			set.Remove(giantINT64Slice[0])
			// last element
			set.Remove(giantINT64Slice[len(giantINT64Slice)-1])
			// arbitrary elements
			set.Remove(giantINT64Slice[1000], giantINT64Slice[2000], giantINT64Slice[3000])
			require.Equal(t, set.Length(), len(giantINT64Slice)-5)
		})
	})
	t.Run("When list with one item is given", func(t *testing.T) {
		set := NewLinkedHashSetINT64()
		set.Add(1)
		t.Run("It should remove the element from the set", func(t *testing.T) {
			set.Remove(1)
			require.Equal(t, set.Length(), 0)
		})
	})
}

func TestLinkedHashSetINT64Iter(t *testing.T) {
	t.Run("It should iterate over all elements of the set respecting the insertion order", func(t *testing.T) {
		set := NewLinkedHashSetINT64()
		set.Add(giantINT64Slice...)
		var (
			i                  int
			somethingWentWrong bool
		)
		for value := range set.Iter() {
			if value != giantINT64Slice[i] {
				somethingWentWrong = true
				break
			}
			i++
		}
		require.False(t, somethingWentWrong)
		require.Equal(t, i, giantSliceLength)
	})
}

func TestLinkedHashSetINT64Length(t *testing.T) {
	t.Run("It should return the correct length of the Set", func(t *testing.T) {
		set := NewLinkedHashSetINT64()
		set.Add(0, 1, 2, 99, 93, 32, 00, 01, 2)
		require.Equal(t, set.Length(), 6)
		set.Remove(1)
		require.Equal(t, set.Length(), 5)
		set.Remove(2, 99, 94)
		require.Equal(t, set.Length(), 3)
		set.Add(94)
		require.Equal(t, set.Length(), 4)
	})

	t.Run("It should return the correct length of the Set no matter the length of the Set", func(t *testing.T) {
		set := NewLinkedHashSetINT64()
		set.Add(giantINT64Slice...)
		require.Equal(t, set.Length(), len(giantINT64Slice))
	})
}

func TestInArray(t *testing.T) {
	t.Run("When the element is in the list", func(t *testing.T) {
		set := NewLinkedHashSetINT64(2, 4, 6, 8)
		require.True(t, set.InArray(2))
		require.True(t, set.InArray(4))
		require.True(t, set.InArray(6))
		require.True(t, set.InArray(8))
	})
	t.Run("When the element is not in the list", func(t *testing.T) {
		set := NewLinkedHashSetINT64(2, 4, 6, 8)
		require.False(t, set.InArray(1))
		require.False(t, set.InArray(3))
		require.False(t, set.InArray(5))
		require.False(t, set.InArray(7))
	})
	t.Run("When the list is empty", func(t *testing.T) {
		set := NewLinkedHashSetINT64()
		require.False(t, set.InArray(1))
		require.False(t, set.InArray(3))
		require.False(t, set.InArray(5))
		require.False(t, set.InArray(7))
	})
}

func TestAsSlice(t *testing.T) {
	t.Run("It should return the correct slice", func(t *testing.T) {
		expectedArray := []int64{2, 4, 6, 8}
		set := NewLinkedHashSetINT64(expectedArray...)

		array := set.AsSlice()
		require.Len(t, array, len(expectedArray))

		for i, value := range array {
			require.Equal(t, value, expectedArray[i])
		}
	})
}

func TestAsInterface(t *testing.T) {
	t.Run("It should return the correct slice", func(t *testing.T) {
		expectedArray := []int64{2, 4, 6, 8}
		set := NewLinkedHashSetINT64(expectedArray...)

		array := set.AsInterface()
		require.Len(t, array, len(expectedArray))

		for i, value := range array {
			require.Equal(t, value.(int64), expectedArray[i])
		}
	})
}
