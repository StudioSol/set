package set

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const giantSliceLength = 100000

var giantINT64Slice = make([]int64, giantSliceLength)

func init() {
	for i := 0; i < giantSliceLength; i++ {
		giantINT64Slice[i] = int64(i + 1)
	}
}

func TestLinkedHashSetINT64Add(t *testing.T) {
	Convey("Given LinkedHashSetINT64.Add", t, func() {
		Convey("It should not store elements that are already on the Set", func() {
			set := NewLinkedHashSetINT64()
			set.Add(0, 0)
			set.Add(0)
			So(set.Length(), ShouldEqual, 1)
		})
		Convey("It should store elements with the correct constraints", func() {
			set := NewLinkedHashSetINT64()
			set.Add(0, 1, 2, 99, 93, 32, 00, 01, 2)
			So(set.Length(), ShouldEqual, 6)
		})
	})
}

func TestLinkedHashSetINT64Remove(t *testing.T) {
	Convey("Given LinkedHashSetINT64.Remove", t, func() {
		Convey("When a big list is given", func() {
			set := NewLinkedHashSetINT64()
			set.Add(giantINT64Slice...)
			Convey("It should remove elements from a Set", func() {
				// first element
				set.Remove(giantINT64Slice[0])
				set.Remove(giantINT64Slice[0])
				set.Remove(giantINT64Slice[0])
				set.Remove(giantINT64Slice[0])
				// last element
				set.Remove(giantINT64Slice[len(giantINT64Slice)-1])
				// arbitrary elements
				set.Remove(giantINT64Slice[1000], giantINT64Slice[2000], giantINT64Slice[3000])
				So(set.Length(), ShouldEqual, len(giantINT64Slice)-5)
			})
		})
		Convey("When list with one item is given", func() {
			set := NewLinkedHashSetINT64()
			set.Add(1)
			Convey("It should remove the element from the set", func() {
				set.Remove(1)
				So(set.Length(), ShouldEqual, 0)
			})
		})
	})
}

func TestLinkedHashSetINT64Iter(t *testing.T) {
	Convey("Given LinkedHashSetINT64.Iter", t, func() {
		Convey("It should iterate over all elements of the set respecting the insertion order", func() {
			set := NewLinkedHashSetINT64()
			set.Add(giantINT64Slice...)
			var (
				i                  int
				somethingWentWrong bool
			)
			for value := range set.Iter() {
				if value != giantINT64Slice[i] {
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

func TestLinkedHashSetINT64Length(t *testing.T) {
	Convey("Given LinkedHashSetINT64.Length", t, func() {
		Convey("It should return the correct length of the Set", func() {
			set := NewLinkedHashSetINT64()
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
			set := NewLinkedHashSetINT64()
			set.Add(giantINT64Slice...)
			So(set.Length(), ShouldEqual, len(giantINT64Slice))
		})
	})
}

func TestInArray(t *testing.T) {
	Convey("Given LinkedHashSetINT64.InArray", t, func() {
		Convey("When the element is in the list", func() {
			set := NewLinkedHashSetINT64(2, 4, 6, 8)
			So(set.InArray(2), ShouldBeTrue)
			So(set.InArray(4), ShouldBeTrue)
			So(set.InArray(6), ShouldBeTrue)
			So(set.InArray(8), ShouldBeTrue)
		})
		Convey("When the element is not in the list", func() {
			set := NewLinkedHashSetINT64(2, 4, 6, 8)
			So(set.InArray(1), ShouldBeFalse)
			So(set.InArray(3), ShouldBeFalse)
			So(set.InArray(5), ShouldBeFalse)
			So(set.InArray(7), ShouldBeFalse)
		})
		Convey("When the list is empty", func() {
			set := NewLinkedHashSetINT64()
			So(set.InArray(1), ShouldBeFalse)
			So(set.InArray(3), ShouldBeFalse)
			So(set.InArray(5), ShouldBeFalse)
			So(set.InArray(7), ShouldBeFalse)
		})
	})
}

func TestAsSlice(t *testing.T) {
	Convey("Given LinkedHashSetINT64.AsSlice", t, func() {
		Convey("It should return the correct slice", func() {
			expectedArray := []int64{2, 4, 6, 8}
			set := NewLinkedHashSetINT64(expectedArray...)

			array := set.AsSlice()
			So(array, ShouldHaveLength, len(expectedArray))

			for i, value := range array {
				So(value, ShouldEqual, expectedArray[i])
			}
		})
	})
}

func TestAsInterface(t *testing.T) {
	Convey("Given LinkedHashSetINT64.AsInterface", t, func() {
		Convey("It should return the correct slice", func() {
			expectedArray := []int64{2, 4, 6, 8}
			set := NewLinkedHashSetINT64(expectedArray...)

			array := set.AsInterface()
			So(array, ShouldHaveLength, len(expectedArray))

			for i, value := range array {
				So(value.(int64), ShouldEqual, expectedArray[i])
			}
		})
	})
}
