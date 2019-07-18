package set

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var giantINTSlice = make([]int, giantSliceLength)

func init() {
	for i := 0; i < giantSliceLength; i++ {
		giantINTSlice[i] = int(i + 1)
	}
}

func TestLinkedHashSetINTAdd(t *testing.T) {
	Convey("Given LinkedHashSetINT.Add", t, func() {
		Convey("It should not store elements that are already on the Set", func() {
			set := NewLinkedHashSetINT()
			set.Add(0, 0)
			set.Add(0)
			So(set.Length(), ShouldEqual, 1)
		})
		Convey("It should store elements with the correct constraints", func() {
			set := NewLinkedHashSetINT()
			set.Add(0, 1, 2, 99, 93, 32, 00, 01, 2)
			So(set.Length(), ShouldEqual, 6)
		})
	})
}

func TestLinkedHashSetINTRemove(t *testing.T) {
	Convey("Given LinkedHashSetINT.Remove", t, func() {
		Convey("When a big list is given", func() {
			set := NewLinkedHashSetINT()
			set.Add(giantINTSlice...)
			Convey("It should remove elements from a Set", func() {
				// first element
				set.Remove(giantINTSlice[0])
				set.Remove(giantINTSlice[0])
				set.Remove(giantINTSlice[0])
				set.Remove(giantINTSlice[0])
				// last element
				set.Remove(giantINTSlice[len(giantINTSlice)-1])
				// arbitrary elements
				set.Remove(giantINTSlice[1000], giantINTSlice[2000], giantINTSlice[3000])
				So(set.Length(), ShouldEqual, len(giantINTSlice)-5)
			})
		})
		Convey("When list with one item is given", func() {
			set := NewLinkedHashSetINT()
			set.Add(1)
			Convey("It should remove the element from the set", func() {
				set.Remove(1)
				So(set.Length(), ShouldEqual, 0)
			})
		})
	})
}

func TestLinkedHashSetINTIter(t *testing.T) {
	Convey("Given LinkedHashSetINT.Iter", t, func() {
		Convey("It should iterate over all elements of the set respecting the insertion order", func() {
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
			So(somethingWentWrong, ShouldBeFalse)
			So(i, ShouldEqual, giantSliceLength)
		})
	})
}

func TestLinkedHashSetINTLength(t *testing.T) {
	Convey("Given LinkedHashSetINT.Length", t, func() {
		Convey("It should return the correct length of the Set", func() {
			set := NewLinkedHashSetINT()
			set.Add(0, 1, 2, 99, 93, 32, 00, 01, 2)
			So(set.Length(), ShouldEqual, 6)
			set.Remove(1)
			So(set.Length(), ShouldEqual, 5)
			set.Remove(2, 99, 94)
			So(set.Length(), ShouldEqual, 3)
			set.Add(94)
			So(set.Length(), ShouldEqual, 4)
		})

		Convey("It should return the correct length of the Set no matter the length of the Set", func() {
			set := NewLinkedHashSetINT()
			set.Add(giantINTSlice...)
			So(set.Length(), ShouldEqual, len(giantINTSlice))
		})
	})
}

func TestIntInArray(t *testing.T) {
	Convey("Given LinkedHashSetINT.InArray", t, func() {
		Convey("When the element is in the list", func() {
			set := NewLinkedHashSetINT(2, 4, 6, 8)
			So(set.InArray(2), ShouldBeTrue)
			So(set.InArray(4), ShouldBeTrue)
			So(set.InArray(6), ShouldBeTrue)
			So(set.InArray(8), ShouldBeTrue)
		})
		Convey("When the element is not in the list", func() {
			set := NewLinkedHashSetINT(2, 4, 6, 8)
			So(set.InArray(1), ShouldBeFalse)
			So(set.InArray(3), ShouldBeFalse)
			So(set.InArray(5), ShouldBeFalse)
			So(set.InArray(7), ShouldBeFalse)
		})
		Convey("When the list is empty", func() {
			set := NewLinkedHashSetINT()
			So(set.InArray(1), ShouldBeFalse)
			So(set.InArray(3), ShouldBeFalse)
			So(set.InArray(5), ShouldBeFalse)
			So(set.InArray(7), ShouldBeFalse)
		})
	})
}

func TestIntAsSlice(t *testing.T) {
	Convey("Given LinkedHashSetINT.AsSlice", t, func() {
		Convey("It should return the correct slice", func() {
			expectedArray := []int{2, 4, 6, 8}
			set := NewLinkedHashSetINT(expectedArray...)

			array := set.AsSlice()
			So(array, ShouldHaveLength, len(expectedArray))

			for i, value := range array {
				So(value, ShouldEqual, expectedArray[i])
			}
		})
	})
}

func TestIntAsInterface(t *testing.T) {
	Convey("Given LinkedHashSetINT.AsInterface", t, func() {
		Convey("It should return the correct slice", func() {
			expectedArray := []int{2, 4, 6, 8}
			set := NewLinkedHashSetINT(expectedArray...)

			array := set.AsInterface()
			So(array, ShouldHaveLength, len(expectedArray))

			for i, value := range array {
				So(value.(int), ShouldEqual, expectedArray[i])
			}
		})
	})
}
