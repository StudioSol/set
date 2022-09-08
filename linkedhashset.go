package set

// LinkedHashSet linked hash set implementation using linkedHashMap as its
// underlying data structure.
//
// - Does not allow storing duplicated values
// - Does not allow storing nil values
// - Maintains insertion order over iteration
type LinkedHashSet[T comparable] struct {
	linkedHashMap *linkedHashMap
}

// Add adds elements to the linked hash set
func (l *LinkedHashSet[T]) Add(elements ...T) {
	for _, element := range elements {
		l.linkedHashMap.Put(element, nil)
	}
}

// Remove removes elements from the linked hash set
func (l *LinkedHashSet[T]) Remove(elements ...T) {
	for _, element := range elements {
		l.linkedHashMap.Remove(element)
	}
}

// Iter iterates over each element of the linked hash set
func (l *LinkedHashSet[T]) Iter() <-chan T {
	ch := make(chan T, l.Length())
	go func() {
		for element := range l.linkedHashMap.Iter() {
			ch <- element.key.(T)
		}
		close(ch)
	}()
	return ch
}

// Length returns the length of the linked hash set
func (l *LinkedHashSet[T]) Length() int {
	return l.linkedHashMap.Length()
}

// AsSlice returns a slice of all values of the linked hash set
func (l *LinkedHashSet[T]) AsSlice() []T {
	values := make([]T, 0, l.Length())
	for value := range l.Iter() {
		values = append(values, value)
	}
	return values
}

// AsInterface returns a slice of all values of the linked hash set
// as interface{}
func (l *LinkedHashSet[T]) AsInterface() []interface{} {
	values := make([]interface{}, 0, l.Length())
	for value := range l.Iter() {
		values = append(values, value)
	}
	return values
}

// InArray returns whether the given item is in array or not
func (l *LinkedHashSet[T]) InArray(search T) bool {
	for item := range l.Iter() {
		if item == search {
			return true
		}
	}
	return false
}

// NewLinkedHashSet returns a new LinkedHashSet with the provided items
func NewLinkedHashSet[T comparable](items ...T) *LinkedHashSet[T] {
	lhm := &LinkedHashSet[T]{
		linkedHashMap: newLinkedHashMap(),
	}
	if len(items) > 0 {
		lhm.Add(items...)
	}
	return lhm
}
