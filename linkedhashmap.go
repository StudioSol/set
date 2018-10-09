package set

import (
	"fmt"
	"hash/fnv"
)

type entry struct {
	key    interface{}
	value  interface{}
	before *entry
	after  *entry
}

// linkedHashMap stores data in key-value pairs while maintaining insertion order
//
// - Uses a doubly linked list to maintain insertion order
type linkedHashMap struct {
	table         map[uint64]*entry
	header        *entry
	last          *entry
	currentLength int
}

// Length returns the length of the linked hash map
func (l *linkedHashMap) Length() int {
	return l.currentLength
}

// Put puts a new entry into the linked hash map
func (l *linkedHashMap) Put(key, value interface{}) {
	if key == nil {
		return
	}

	hash := l.hash(key)
	newEntry := &entry{
		key:   key,
		value: value,
	}

	if _, ok := l.table[hash]; ok {
		return
	}

	l.table[hash] = newEntry
	l.currentLength++

	// put to first position
	if l.header == nil {
		l.header = newEntry
		l.last = newEntry
		return
	}

	// put to last position
	l.last.after = newEntry
	newEntry.before = l.last
	l.last = newEntry
}

// Get gets an entry from the linked hash map
func (l *linkedHashMap) Get(key interface{}) interface{} {
	hash := l.hash(key)
	if _, ok := l.table[hash]; !ok {
		return nil
	}

	tmp := l.table[hash]
	for tmp != nil {
		if tmp.key == key {
			return tmp.value
		}
		tmp = tmp.after
	}

	return nil
}

// Remove removes an entry from the linked hash map
func (l *linkedHashMap) Remove(key interface{}) bool {
	hash := l.hash(key)
	if _, ok := l.table[hash]; !ok {
		return false
	}

	current := l.table[hash]

	switch current.key {
	// first entry
	case l.header.key:
		if l.header == l.last {
			l.header = nil
			l.last = nil
		} else {
			l.header = l.header.after
			l.header.before = nil
		}
	// last entry
	case l.last.key:
		if l.header == l.last {
			l.header = nil
			l.last = nil
		} else {
			l.last = l.last.before
			l.last.after = nil
		}
	// arbitrary entry
	default:
		var shouldNotUpdate bool
		innerCurrent := l.header
		for innerCurrent.key != current.key {
			if innerCurrent.after == nil {
				shouldNotUpdate = true
				break
			}
			innerCurrent = innerCurrent.after
		}

		if !shouldNotUpdate {
			innerCurrent.before.after = innerCurrent.after
			innerCurrent.after.before = innerCurrent.before
		}
	}

	delete(l.table, hash)
	l.currentLength--

	return true
}

// Iter iterates over each entry of the linked hash map
func (l *linkedHashMap) Iter() <-chan *entry {
	ch := make(chan *entry, l.currentLength)
	go func() {
		current := l.header
		for current != nil {
			ch <- current
			current = current.after
		}
		close(ch)
	}()
	return ch
}

func (l *linkedHashMap) hash(key interface{}) uint64 {
	h := fnv.New64()
	_, err := h.Write([]byte(fmt.Sprintf("%#v", key)))
	if err != nil {
		// noop
	}
	return h.Sum64()
}

func newLinkedHashMap() *linkedHashMap {
	return &linkedHashMap{
		table: make(map[uint64]*entry),
	}
}
