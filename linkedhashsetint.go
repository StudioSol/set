package set

// LinkedHashSetINT linked hash set implementation using linkedHashMap as its
// underlying data structure.
//
// - Does not allow storing duplicated values
// - Does not allow storing nil values
// - Maintains insertion order over iteration
type LinkedHashSetINT struct {
	linkedHashMap *linkedHashMap
}

// Add adds elements to the linked hash set
func (l *LinkedHashSetINT) Add(elements ...int) {
	for _, element := range elements {
		l.linkedHashMap.Put(element, nil)
	}
}

// Remove removes elements from the linked hash set
func (l *LinkedHashSetINT) Remove(elements ...int) {
	for _, element := range elements {
		l.linkedHashMap.Remove(element)
	}
}

// Iter iterates over each element of the linked hash set
func (l *LinkedHashSetINT) Iter() <-chan int {
	ch := make(chan int, l.Length())
	go func() {
		for element := range l.linkedHashMap.Iter() {
			ch <- element.key.(int)
		}
		close(ch)
	}()
	return ch
}

// Length returns the length of the linked hash set
func (l *LinkedHashSetINT) Length() int {
	return l.linkedHashMap.Length()
}

// AsSlice returns a slice of all values of the linked hash set
func (l *LinkedHashSetINT) AsSlice() []int {
	values := make([]int, 0, l.Length())
	for value := range l.Iter() {
		values = append(values, value)
	}
	return values
}

// AsInterface returns a slice of all values of the linked hash set
// as interface{}
func (l *LinkedHashSetINT) AsInterface() []interface{} {
	values := make([]interface{}, 0, l.Length())
	for value := range l.Iter() {
		values = append(values, value)
	}
	return values
}

// InArray returns whether the given item is in array or not
func (l *LinkedHashSetINT) InArray(search int) bool {
	for item := range l.Iter() {
		if item == search {
			return true
		}
	}
	return false
}

// NewLinkedHashSetINT returns a new LinkedHashSetINT with the provided items
func NewLinkedHashSetINT(ints ...int) *LinkedHashSetINT {
	lhm := &LinkedHashSetINT{
		linkedHashMap: newLinkedHashMap(),
	}
	if len(ints) > 0 {
		lhm.Add(ints...)
	}
	return lhm
}
