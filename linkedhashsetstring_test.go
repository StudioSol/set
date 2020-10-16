package set

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var giantStringSlice = make([]string, giantSliceLength)

func init() {
	for i := 0; i < giantSliceLength; i++ {
		giantStringSlice[i] = strconv.Itoa(i + 1)
	}
}

func TestLinkedHashSetStringAdd(t *testing.T) {
	t.Run("It should not store elements that are already on the Set", func(t *testing.T) {
		set := NewLinkedHashSetString()
		set.Add("0", "0")
		set.Add("0")
		require.Equal(t, set.Length(), 1)
	})

	t.Run("It should store elements with the correct constraints", func(t *testing.T) {
		set := NewLinkedHashSetString()
		set.Add("0", "1", "2", "99", "93", "32", "00", "01", "2")
		require.Equal(t, set.Length(), 8)
	})
}

func TestLinkedHashSetStringRemove(t *testing.T) {
	set := NewLinkedHashSetString()
	set.Add(giantStringSlice...)

	// first element
	set.Remove(giantStringSlice[0])
	set.Remove(giantStringSlice[0])
	set.Remove(giantStringSlice[0])
	set.Remove(giantStringSlice[0])

	// last element
	set.Remove(giantStringSlice[len(giantStringSlice)-1])

	// arbitrary elements
	set.Remove(giantStringSlice[1000], giantStringSlice[2000], giantStringSlice[3000])
	require.Equal(t, set.Length(), len(giantStringSlice)-5)
}

func TestLinkedHashSetStringIter(t *testing.T) {
	set := NewLinkedHashSetString()
	set.Add(giantStringSlice...)

	var (
		i                  int
		somethingWentWrong bool
	)
	for value := range set.Iter() {
		if value != giantStringSlice[i] {
			somethingWentWrong = true
			break
		}
		i++
	}
	require.False(t, somethingWentWrong)
	require.Equal(t, i, giantSliceLength)
}

func TestLinkedHashSetStringLength(t *testing.T) {
	t.Run("small set", func(t *testing.T) {
		set := NewLinkedHashSetString()
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
		set := NewLinkedHashSetString()
		set.Add(giantStringSlice...)
		require.Equal(t, set.Length(), len(giantStringSlice))
	})
}

func TestLinkedHashSetStringInArray(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		set := NewLinkedHashSetString("02", "04", "06", "08")
		require.True(t, set.InArray("02"))
		require.True(t, set.InArray("04"))
		require.True(t, set.InArray("06"))
		require.True(t, set.InArray("08"))
	})

	t.Run("not found", func(t *testing.T) {
		set := NewLinkedHashSetString("02", "04", "06", "08")
		require.False(t, set.InArray("01"))
		require.False(t, set.InArray("03"))
		require.False(t, set.InArray("05"))
		require.False(t, set.InArray("07"))
	})

	t.Run("empty", func(t *testing.T) {
		set := NewLinkedHashSetString()
		require.False(t, set.InArray("01"))
		require.False(t, set.InArray("03"))
		require.False(t, set.InArray("05"))
		require.False(t, set.InArray("07"))
	})
}

func TestLinkedHashSetStringAsSlice(t *testing.T) {
	expectedArray := []string{"0", "1", "2", "99", "93", "32", "00", "01"}
	set := NewLinkedHashSetString(expectedArray...)

	array := set.AsSlice()
	require.Len(t, array, len(expectedArray))

	for i, value := range array {
		require.Equal(t, value, expectedArray[i])
	}
}

func TestLinkedHashSetStringAsInterface(t *testing.T) {
	expectedArray := []string{"0", "1", "2", "99", "93", "32", "00", "01"}
	set := NewLinkedHashSetString(expectedArray...)

	array := set.AsInterface()
	require.Len(t, array, len(expectedArray))

	for i, value := range array {
		require.Equal(t, value.(string), expectedArray[i])
	}
}
