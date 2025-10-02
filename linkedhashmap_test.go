package set

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Run("When the key exists", func(t *testing.T) {
		set := newLinkedHashMap()
		value := rand.Int()
		set.Put("test", value)

		v, exists := set.Get("test")
		require.Equal(t, value, v)
		require.True(t, exists)
	})

	t.Run("When the key not exists", func(t *testing.T) {
		set := newLinkedHashMap()
		result, exists := set.Get("bla")
		require.Nil(t, result)
		require.False(t, exists)
	})
}

func TestPut(t *testing.T) {
	t.Run("valid key", func(t *testing.T) {
		set := newLinkedHashMap()
		value := rand.Int()
		set.Put("test", value)

		v, exists := set.Get("test")
		require.Equal(t, value, v)
		require.True(t, exists)
	})

	t.Run("invalid key", func(t *testing.T) {
		set := newLinkedHashMap()

		value := rand.Int()
		set.Put(nil, value)

		require.Equal(t, 0, set.Length())
	})
}

func TestRemove(t *testing.T) {
	testNumbers := []int{1, 2, 3}
	t.Run("first value", func(t *testing.T) {
		set := newLinkedHashMap()
		for _, number := range testNumbers {
			set.Put(number, number)
		}

		set.Remove(1)
		require.Equal(t, 2, set.Length())

		v, exists := set.Get(1)
		require.Nil(t, v)
		require.False(t, exists)
	})

	t.Run("last value", func(t *testing.T) {
		set := newLinkedHashMap()
		for _, number := range testNumbers {
			set.Put(number, number)
		}

		set.Remove(3)
		require.Equal(t, 2, set.Length())

		v, exists := set.Get(3)
		require.Nil(t, v)
		require.False(t, exists)
	})

	t.Run("middle value", func(t *testing.T) {
		set := newLinkedHashMap()
		for _, number := range testNumbers {
			set.Put(number, number)
		}

		set.Remove(2)
		require.Equal(t, 2, set.Length())

		v, exists := set.Get(2)
		require.Nil(t, v)
		require.False(t, exists)
	})

	t.Run("single value", func(t *testing.T) {
		set := newLinkedHashMap()
		set.Put(1, 1)
		set.Remove(1)
		require.Equal(t, 0, set.Length())

		v, exists := set.Get(1)
		require.Nil(t, v)
		require.False(t, exists)
	})
}
