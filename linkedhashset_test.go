package set

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var giantGenericSlice = make([]string, giantSliceLength)

type testStruct struct {
	key   string
	value int
}

func init() {
	for i := 0; i < giantSliceLength; i++ {
		giantGenericSlice[i] = strconv.Itoa(i + 1)
	}
}

func TestLinkedHashSetAdd(t *testing.T) {
	t.Run("It should not store elements that are already on the Set", func(t *testing.T) {
		set := NewLinkedHashSet[string]()
		set.Add("0", "0")
		set.Add("0")
		require.Equal(t, set.Length(), 1)
	})

	t.Run("It should store elements with the correct constraints", func(t *testing.T) {
		set := NewLinkedHashSet[string]()
		set.Add("0", "1", "2", "99", "93", "32", "00", "01", "2")
		require.Equal(t, set.Length(), 8)
	})

	t.Run("I should store elements typeds by structs", func(t *testing.T) {
		set := NewLinkedHashSet[testStruct]()
		item1 := testStruct{
			key:   "a",
			value: 1,
		}
		item2 := testStruct{
			key:   "b",
			value: 2,
		}
		item3 := testStruct{
			key:   "c",
			value: 3,
		}
		set.Add(item1, item2, item3, item1)
		require.Equal(t, set.Length(), 3)
	})
}

func TestLinkedHashSetRemove(t *testing.T) {
	set := NewLinkedHashSet[string]()
	set.Add(giantGenericSlice...)

	// first element
	set.Remove(giantGenericSlice[0])
	set.Remove(giantGenericSlice[0])
	set.Remove(giantGenericSlice[0])
	set.Remove(giantGenericSlice[0])

	// last element
	set.Remove(giantGenericSlice[len(giantGenericSlice)-1])

	// arbitrary elements
	set.Remove(giantGenericSlice[1000], giantGenericSlice[2000], giantGenericSlice[3000])
	require.Equal(t, set.Length(), len(giantGenericSlice)-5)
}

func TestLinkedHashSetIter(t *testing.T) {
	set := NewLinkedHashSet[string]()
	set.Add(giantGenericSlice...)

	var (
		i                  int
		somethingWentWrong bool
	)
	for value := range set.Iter() {
		if value != giantGenericSlice[i] {
			somethingWentWrong = true
			break
		}
		i++
	}
	require.False(t, somethingWentWrong)
	require.Equal(t, i, giantSliceLength)
}

func TestLinkedHashSetLength(t *testing.T) {
	t.Run("small set", func(t *testing.T) {
		set := NewLinkedHashSet[string]()
		set.Add("0", "1", "2", "99", "93", "32", "00", "01", "2")
		require.Equal(t, set.Length(), 8)
		set.Remove("1")
		require.Equal(t, set.Length(), 7)
		set.Remove("2", "99", "94")
		require.Equal(t, set.Length(), 5)
		set.Add("94")
		require.Equal(t, set.Length(), 6)
	})

	t.Run("big set", func(t *testing.T) {
		set := NewLinkedHashSet[string]()
		set.Add(giantGenericSlice...)
		require.Equal(t, set.Length(), len(giantGenericSlice))
	})
}

func TestLinkedHashSetInArray(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		item1 := testStruct{
			key:   "a",
			value: 1,
		}
		item2 := testStruct{
			key:   "b",
			value: 2,
		}
		item3 := testStruct{
			key:   "c",
			value: 3,
		}
		set := NewLinkedHashSet(item1, item2, item3, item1)

		require.True(t, set.InArray(item1))
		require.True(t, set.InArray(item2))
		require.True(t, set.InArray(item3))
	})

	t.Run("not found", func(t *testing.T) {
		set := NewLinkedHashSet("02", "04", "06", "08")
		require.False(t, set.InArray("01"))
		require.False(t, set.InArray("03"))
		require.False(t, set.InArray("05"))
		require.False(t, set.InArray("07"))
	})

	t.Run("empty", func(t *testing.T) {
		set := NewLinkedHashSet[string]()
		require.False(t, set.InArray("01"))
		require.False(t, set.InArray("03"))
		require.False(t, set.InArray("05"))
		require.False(t, set.InArray("07"))
	})
}

func TestLinkedHashSetAsSlice(t *testing.T) {
	item1 := testStruct{
		key:   "a",
		value: 1,
	}
	item2 := testStruct{
		key:   "b",
		value: 2,
	}
	item3 := testStruct{
		key:   "c",
		value: 3,
	}
	expectedArray := []testStruct{item1, item2, item3}

	set := NewLinkedHashSet(expectedArray...)

	array := set.AsSlice()
	require.Len(t, array, len(expectedArray))

	for i, value := range array {
		require.Equal(t, value, expectedArray[i])
	}
}

func TestLinkedHashSetAsInterface(t *testing.T) {
	item1 := testStruct{
		key:   "a",
		value: 1,
	}
	item2 := testStruct{
		key:   "b",
		value: 2,
	}
	item3 := testStruct{
		key:   "c",
		value: 3,
	}
	expectedArray := []testStruct{item1, item2, item3}
	set := NewLinkedHashSet(expectedArray...)

	array := set.AsInterface()
	require.Len(t, array, len(expectedArray))

	for i, value := range array {
		require.Equal(t, value.(testStruct), expectedArray[i])
	}
}

func TestLinkedHashSetContains(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		set := NewLinkedHashSet("02", "04", "06", "08")

		require.True(t, set.Contains("02"))
		require.True(t, set.Contains("04"))
		require.True(t, set.Contains("06"))
		require.True(t, set.Contains("08"))
	})

	t.Run("not found", func(t *testing.T) {
		set := NewLinkedHashSet("02", "04", "06", "08")
		require.False(t, set.Contains("01"))
		require.False(t, set.Contains("03"))
		require.False(t, set.Contains("05"))
		require.False(t, set.Contains("07"))
	})

	t.Run("empty", func(t *testing.T) {
		set := NewLinkedHashSet[string]()
		require.False(t, set.Contains("01"))
		require.False(t, set.Contains("03"))
		require.False(t, set.Contains("05"))
		require.False(t, set.Contains("07"))
	})
}
