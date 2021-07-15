package main

import(
	"container/list"
	"fmt"
)

func main(){

	mylist := list.New()

	mylist.PushBack("A")
	mylist.PushBack(100)
	mylist.PushBack(true) //append back
	mylist.PushFront("A") //append front

	for e := mylist.Front(); e!= nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}