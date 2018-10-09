package set

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	Convey("Given LinkedHashSet.Get", t, func() {
		Convey("When the key exists", func() {
			set := newLinkedHashMap()
			value := rand.Int()
			set.Put("test", value)
			Convey("It should return the expected value", func() {
				So(set.Get("test"), ShouldEqual, value)
			})
		})

		Convey("When the key not exists", func() {
			set := newLinkedHashMap()
			Convey("It should return an empty value", func() {
				result := set.Get("bla")
				So(result, ShouldEqual, nil)
			})
		})
	})
}


func TestPut(t *testing.T) {
	Convey("Given LinkedHashSet.Put", t, func() {
		Convey("When an invalid key is given", func() {
			set := newLinkedHashMap()

			value := rand.Int()
			set.Put(nil, value)

			So(set.Length(), ShouldEqual, 0)
		})
	})
}

func TestRemove(t *testing.T) {
	Convey("Given a valid list of numbers", t, func() {
		testNumbers := []int{1,2,3}
		set := newLinkedHashMap()
		for _, number := range testNumbers {
			set.Put(number, number)
		}
		Convey("When first value is removed", func() {
			setCopy := *set
			setCopy.Remove(1)
			Convey("It should remove the correct value", func() {
				So(setCopy.Length(), ShouldEqual, 2)
				So(setCopy.Get(1), ShouldBeNil)
			})
		})
		Convey("When last value is removed", func() {
			setCopy := *set
			setCopy.Remove(3)
			Convey("It should remove the correct value", func() {
				So(setCopy.Length(), ShouldEqual, 2)
				So(setCopy.Get(3), ShouldBeNil)
			})
		})
		Convey("When a middle value is removed", func() {
			setCopy := *set
			setCopy.Remove(2)
			Convey("It should remove the correct value", func() {
				So(setCopy.Length(), ShouldEqual, 2)
				So(setCopy.Get(2), ShouldBeNil)
			})
		})
	})
	Convey("Given a list with single value", t, func() {
		set := newLinkedHashMap()
		set.Put(1, 1)
		Convey("When the values is removed", func() {
			setCopy := *set
			setCopy.Remove(1)
			Convey("It should remove the correct value", func() {
				So(setCopy.Length(), ShouldEqual, 0)
				So(setCopy.Get(1), ShouldBeNil)
			})
		})
	})
}
