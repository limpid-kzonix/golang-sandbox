package main

import (
	"fmt"

	katas "github.com/limpid-kzonix/golang-sandbox/main/katas/sub"
)

func main() {
	arr1 := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	res := katas.CountSheeps(arr1)
	s := Test()
	str := fmt.Sprintf("Count = %d, %d", res, s)
	fmt.Println(str)
}
