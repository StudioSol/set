package set

// LinkedHashSetString linked hash set implementation using linkedHashMap as its
// underlying data structure.
//
// - Does not allow storing duplicated values
// - Does not allow storing nil values
// - Maintains insertion order over iteration
type LinkedHashSetString struct {
	linkedHashMap *linkedHashMap
}

// Add adds elements to the linked hash set
func (l *LinkedHashSetString) Add(elements ...string) {
	for _, element := range elements {
		l.linkedHashMap.Put(element, nil)
	}
}

// Remove removes elements from the linked hash set
func (l *LinkedHashSetString) Remove(elements ...string) {
	for _, element := range elements {
		l.linkedHashMap.Remove(element)
	}
}

// Iter iterates over each element of the linked hash set
func (l *LinkedHashSetString) Iter() <-chan string {
	ch := make(chan string, l.Length())
	go func() {
		for element := range l.linkedHashMap.Iter() {
			ch <- element.key.(string)
		}
		close(ch)
	}()
	return ch
}

// Length returns the length of the linked hash set
func (l *LinkedHashSetString) Length() int {
	return l.linkedHashMap.Length()
}

// AsSlice returns a slice of all values of the linked hash set
func (l *LinkedHashSetString) AsSlice() []string {
	values := make([]string, 0, l.Length())
	for value := range l.Iter() {
		values = append(values, value)
	}
	return values
}

// AsInterface returns a slice of all values of the linked hash set
// as interface{}
func (l *LinkedHashSetString) AsInterface() []interface{} {
	values := make([]interface{}, 0, l.Length())
	for value := range l.Iter() {
		values = append(values, value)
	}
	return values
}

// InArray returns whether the given item is in array or not
func (l *LinkedHashSetString) InArray(search string) bool {
	for item := range l.Iter() {
		if item == search {
			return true
		}
	}
	return false
}

// NewLinkedHashSetString returns a new LinkedHashSetString with the provided items
func NewLinkedHashSetString(strings ...string) *LinkedHashSetString {
	lhm := &LinkedHashSetString{
		linkedHashMap: newLinkedHashMap(),
	}
	if len(strings) > 0 {
		lhm.Add(strings...)
	}
	return lhm
}
