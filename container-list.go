package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Dio")
	data.PushBack("Surandito")
	data.PushBack("Tampan dan Berani")
	// Dari depan ke belakang
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//  Dari belakang ke depan
	// for e := data.Back(); e != nil; e = e.Prev() {
	// 	fmt.Println(e.Value)
	// }
}
