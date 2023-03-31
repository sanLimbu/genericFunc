package main

import (
	"fmt"
	"genericexample/genericlist"
)

func main() {

	genList := genericlist.New[string]()
	genList.Insert("san") //0
	genList.Insert("dev") //1
	genList.Insert("mum") //2
	fmt.Printf("%+v\n", genList)

	genList.Remove(1)
	fmt.Printf("%+v\n", genList)
	genList.RemoveByValue("san")
	fmt.Printf("%+v\n", genList)
}
