package set

// LinkedHashSetINT64 linked hash set implementation using linkedHashMap as its
// underlying data structure.
//
// - Does not allow storing duplicated values
// - Does not allow storing nil values
// - Maintains insertion order over iteration
type LinkedHashSetINT64 struct {
	linkedHashMap *linkedHashMap
}

// Add adds elements to the linked hash set
func (l *LinkedHashSetINT64) Add(elements ...int64) {
	for _, element := range elements {
		l.linkedHashMap.Put(element, nil)
	}
}

// Remove removes elements from the linked hash set
func (l *LinkedHashSetINT64) Remove(elements ...int64) {
	for _, element := range elements {
		l.linkedHashMap.Remove(element)
	}
}

// Iter iterates over each element of the linked hash set
func (l *LinkedHashSetINT64) Iter() <-chan int64 {
	ch := make(chan int64, l.Length())
	go func() {
		for element := range l.linkedHashMap.Iter() {
			ch <- element.key.(int64)
		}
		close(ch)
	}()
	return ch
}

// Length returns the length of the linked hash set
func (l *LinkedHashSetINT64) Length() int {
	return l.linkedHashMap.Length()
}

// AsSlice returns a slice of all values of the linked hash set
func (l *LinkedHashSetINT64) AsSlice() []int64 {
	values := make([]int64, 0, l.Length())
	for value := range l.Iter() {
		values = append(values, value)
	}
	return values
}

// AsInterface returns a slice of all values of the linked hash set
// as interface{}
func (l *LinkedHashSetINT64) AsInterface() []interface{} {
	values := make([]interface{}, 0, l.Length())
	for value := range l.Iter() {
		values = append(values, value)
	}
	return values
}

// InArray returns whether the given item is in array or not
func (l *LinkedHashSetINT64) InArray(search int64) bool {
	for item := range l.Iter() {
		if item == search {
			return true
		}
	}
	return false
}

// NewLinkedHashSetINT64 returns a new LinkedHashSetINT64 with the provided items
func NewLinkedHashSetINT64(ints ...int64) *LinkedHashSetINT64 {
	lhm := &LinkedHashSetINT64{
		linkedHashMap: newLinkedHashMap(),
	}
	if len(ints) > 0 {
		lhm.Add(ints...)
	}
	return lhm
}
