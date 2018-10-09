package set

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var giantStringSlice = make([]string, giantSliceLength)

func init() {
	for i := 0; i < giantSliceLength; i++ {
		giantStringSlice[i] = strconv.Itoa(i + 1)
	}
}

func TestLinkedHashSetStringAdd(t *testing.T) {
	Convey("Given LinkedHashSetString.Add", t, func() {
		Convey("It should not store elements that are already on the Set", func() {
			set := NewLinkedHashSetString()
			set.Add("0", "0")
			set.Add("0")
			So(set.Length(), ShouldEqual, 1)
		})
		Convey("It should store elements with the correct constraints", func() {
			set := NewLinkedHashSetString()
			set.Add("0", "1", "2", "99", "93", "32", "00", "01", "2")
			So(set.Length(), ShouldEqual, 8)
		})
	})
}

func TestLinkedHashSetStringRemove(t *testing.T) {
	Convey("Given LinkedHashSetString.Remove", t, func() {
		Convey("It should remove elements from a Set", func() {
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
			So(set.Length(), ShouldEqual, len(giantStringSlice)-5)
		})
	})
}

func TestLinkedHashSetStringIter(t *testing.T) {
	Convey("Given LinkedHashSetString.Iter", t, func() {
		Convey("It should iterate over all elements of the set respecting the insertion order", func() {
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
			So(somethingWentWrong, ShouldBeFalse)
			So(i, ShouldEqual, giantSliceLength)
		})
	})
}

func TestLinkedHashSetStringLength(t *testing.T) {
	Convey("Given LinkedHashSetString.Length", t, func() {
		Convey("It should return the correct length of the Set", func() {
			set := NewLinkedHashSetString()
			set.Add("0", "1", "2", "99", "93", "32", "00", "01", "2")
			So(set.Length(), ShouldEqual, 8)
			set.Remove("1")
			So(set.Length(), ShouldEqual, 7)
			set.Remove("2", "99", "94")
			So(set.Length(), ShouldEqual, 5)
			set.Add("94")
			So(set.Length(), ShouldEqual, 6)
		})

		Convey("It should return the correct length of the Set no matter the length of the Set", func() {
			set := NewLinkedHashSetString()
			set.Add(giantStringSlice...)
			So(set.Length(), ShouldEqual, len(giantStringSlice))
		})
	})
}

func TestLinkedHashSetStringInArray(t *testing.T) {
	Convey("Given LinkedHashSetString.InArray", t, func() {
		Convey("When the element is in the list", func() {
			set := NewLinkedHashSetString("02", "04", "06", "08")
			So(set.InArray("02"), ShouldBeTrue)
			So(set.InArray("04"), ShouldBeTrue)
			So(set.InArray("06"), ShouldBeTrue)
			So(set.InArray("08"), ShouldBeTrue)
		})
		Convey("When the element is not in the list", func() {
			set := NewLinkedHashSetString("02", "04", "06", "08")
			So(set.InArray("01"), ShouldBeFalse)
			So(set.InArray("03"), ShouldBeFalse)
			So(set.InArray("05"), ShouldBeFalse)
			So(set.InArray("07"), ShouldBeFalse)
		})
		Convey("When the list is empty", func() {
			set := NewLinkedHashSetString()
			So(set.InArray("01"), ShouldBeFalse)
			So(set.InArray("03"), ShouldBeFalse)
			So(set.InArray("05"), ShouldBeFalse)
			So(set.InArray("07"), ShouldBeFalse)
		})
	})
}

func TestLinkedHashSetStringAsSlice(t *testing.T) {
	Convey("Given LinkedHashSetString.AsSlice", t, func() {
		Convey("It should return the correct slice", func() {
			expectedArray := []string{"0", "1", "2", "99", "93", "32", "00", "01"}
			set := NewLinkedHashSetString(expectedArray...)

			array := set.AsSlice()
			So(array, ShouldHaveLength, len(expectedArray))

			for i, value := range array {
				So(value, ShouldEqual, expectedArray[i])
			}
		})
	})
}

func TestLinkedHashSetStringAsInterface(t *testing.T) {
	Convey("Given LinkedHashSetString.AsInterface", t, func() {
		Convey("It should return the correct slice", func() {
			expectedArray := []string{"0", "1", "2", "99", "93", "32", "00", "01"}
			set := NewLinkedHashSetString(expectedArray...)

			array := set.AsInterface()
			So(array, ShouldHaveLength, len(expectedArray))

			for i, value := range array {
				So(value.(string), ShouldEqual, expectedArray[i])
			}
		})
	})
}
