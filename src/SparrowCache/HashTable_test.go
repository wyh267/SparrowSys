package SparrowCache

import (
	"fmt"
	"testing"
	"time"
)

func TestHashTable(t *testing.T) {

	ht := NewHashTable()

	ht.InitHashTable()

	startTime := time.Now()

	for i := 0; i < 1000; i++ {

		ht.Set(fmt.Sprintf("KEY_%v", i), fmt.Sprintf("VALUE_%v", i))

		//mm, _ := ht.Get("yyy")
		//fmt.Printf("%v\n", mm)

	}

	endTime := time.Now()

	fmt.Printf("%v", endTime.Sub(startTime))

	ht.Set("yyy", "mmmdddd")
	mm, _ := ht.Get("KEY_9843")
	fmt.Printf("%v\n", mm)
}
