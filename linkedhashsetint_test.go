package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var giantINTSlice = make([]int, giantSliceLength)

func init() {
	for i := 0; i < giantSliceLength; i++ {
		giantINTSlice[i] = i + 1
	}
}

func TestLinkedHashSetINTAdd(t *testing.T) {
	t.Run("Given LinkedHashSetINT.Add", func(t *testing.T) {
		t.Run("It should not store elements that are already on the Set", func(t *testing.T) {
			set := NewLinkedHashSetINT()
			set.Add(0, 0)
			set.Add(0)
			require.Equal(t, set.Length(), 1)
		})
		t.Run("It should store elements with the correct constraints", func(t *testing.T) {
			set := NewLinkedHashSetINT()
			set.Add(0, 1, 2, 99, 93, 32, 00, 01, 2)
			require.Equal(t, set.Length(), 6)
		})
	})
}

func TestLinkedHashSetINTRemove(t *testing.T) {
	t.Run("Given LinkedHashSetINT.Remove", func(t *testing.T) {
		t.Run("When a big list is given", func(t *testing.T) {
			set := NewLinkedHashSetINT()
			set.Add(giantINTSlice...)
			t.Run("It should remove elements from a Set", func(t *testing.T) {
				// first element
				set.Remove(giantINTSlice[0])
				set.Remove(giantINTSlice[0])
				set.Remove(giantINTSlice[0])
				set.Remove(giantINTSlice[0])
				// last element
				set.Remove(giantINTSlice[len(giantINTSlice)-1])
				// arbitrary elements
				set.Remove(giantINTSlice[1000], giantINTSlice[2000], giantINTSlice[3000])
				require.Equal(t, set.Length(), len(giantINTSlice)-5)
			})
		})
		t.Run("When list with one item is given", func(t *testing.T) {
			set := NewLinkedHashSetINT()
			set.Add(1)
			t.Run("It should remove the element from the set", func(t *testing.T) {
				set.Remove(1)
				require.Equal(t, set.Length(), 0)
			})
		})
	})
}

func TestLinkedHashSetINTIter(t *testing.T) {
	t.Run("Given LinkedHashSetINT.Iter", func(t *testing.T) {
		t.Run("It should iterate over all elements of the set respecting the insertion order", func(t *testing.T) {
			set := NewLinkedHashSetINT()
			set.Add(giantINTSlice...)
			var (
				i                  int
				somethingWentWrong bool
			)
			for value := range set.Iter() {
				if value != giantINTSlice[i] {
					somethingWentWrong = true
					break
				}
				i++
			}
			require.False(t, somethingWentWrong)
			require.Equal(t, i, giantSliceLength)
		})
	})
}

func TestLinkedHashSetINTLength(t *testing.T) {
	t.Run("It should return the correct length of the Set", func(t *testing.T) {
		set := NewLinkedHashSetINT()
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
		set := NewLinkedHashSetINT()
		set.Add(giantINTSlice...)
		require.Equal(t, set.Length(), len(giantINTSlice))
	})
}

func TestIntInArray(t *testing.T) {
	t.Run("When the element is in the list", func(t *testing.T) {
		set := NewLinkedHashSetINT(2, 4, 6, 8)
		require.True(t, set.InArray(2))
		require.True(t, set.InArray(4))
		require.True(t, set.InArray(6))
		require.True(t, set.InArray(8))
	})
	t.Run("When the element is not in the list", func(t *testing.T) {
		set := NewLinkedHashSetINT(2, 4, 6, 8)
		require.False(t, set.InArray(1))
		require.False(t, set.InArray(3))
		require.False(t, set.InArray(5))
		require.False(t, set.InArray(7))
	})
	t.Run("When the list is empty", func(t *testing.T) {
		set := NewLinkedHashSetINT()
		require.False(t, set.InArray(1))
		require.False(t, set.InArray(3))
		require.False(t, set.InArray(5))
		require.False(t, set.InArray(7))
	})
}

func TestIntAsSlice(t *testing.T) {
	t.Run("It should return the correct slice", func(t *testing.T) {
		expectedArray := []int{2, 4, 6, 8}
		set := NewLinkedHashSetINT(expectedArray...)

		array := set.AsSlice()
		require.Len(t, array, len(expectedArray))

		for i, value := range array {
			require.Equal(t, value, expectedArray[i])
		}
	})
}

func TestIntAsInterface(t *testing.T) {
	t.Run("It should return the correct slice", func(t *testing.T) {
		expectedArray := []int{2, 4, 6, 8}
		set := NewLinkedHashSetINT(expectedArray...)

		array := set.AsInterface()
		require.Len(t, array, len(expectedArray))

		for i, value := range array {
			require.Equal(t, value.(int), expectedArray[i])
		}
	})
}
