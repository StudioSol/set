## Set

[![Build Status](https://travis-ci.org/StudioSol/set.svg?branch=master)](https://travis-ci.org/StudioSol/set)
[![codecov](https://codecov.io/gh/StudioSol/set/branch/master/graph/badge.svg)](https://codecov.io/gh/StudioSol/set)
[![Go Report Card](https://goreportcard.com/badge/github.com/StudioSol/set)](https://goreportcard.com/report/github.com/StudioSol/set)
[![GoDoc](https://godoc.org/github.com/StudioSol/set?status.svg)](https://godoc.org/github.com/StudioSol/set)

Set is a simple Set data structure implementation in Go (Golang) using LinkedHashMap.

This library allow you to get a set of `int64` or `string` without duplicated items.

### Usage

```go
package main

import (
	"fmt"

	"github.com/StudioSol/set"
)

func main() {
	duplicatedInt64 := []int64{1, 1, 2, 2, 3, 3}

	unduplicatedInt64 := set.NewLinkedHashSetINT64(duplicatedInt64...)

	// Get a []int64 from set
	unduplicatedArray := unduplicatedInt64.AsSlice()
	fmt.Println(unduplicatedArray) // will print [1 2 3]

	// Get the Length from set
	fmt.Println(unduplicatedInt64.Length()) // will print 3

	// Add new items in set
	unduplicatedInt64.Add(1, 2, 3, 4)
	fmt.Println(unduplicatedInt64.AsSlice()) // will print [1 2 3 4]

	// Check if item is in set
	fmt.Println(unduplicatedInt64.InArray(1)) // will print true
	fmt.Println(unduplicatedInt64.InArray(5)) // will print false

	// Get a []interface{} from set
	interfaceList := unduplicatedInt64.AsInterface()
	fmt.Println(interfaceList) // will print [1 2 3 4]

	// Allow to iter over set
	for i := range unduplicatedInt64.Iter() {
		fmt.Println(i)
	}

	// Remove items from set
	unduplicatedInt64.Remove(0, 1, 2, 3)
	fmt.Println(unduplicatedInt64.AsSlice()) // will print [4]
}

// You have same methods to LinkedHashSetString
```
