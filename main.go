package main

import (
	"fmt"

	katas "github.com/limpid-kzonix/golang-sandbox/main/katas/sub"
	server "github.com/limpid-kzonix/golang-sandbox/modules/server-app"
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

	server.Server()
	res := katas.CountSheeps(arr1)
	s := Test()
	str := fmt.Sprintf("Count = %d, %d", res, s)
	fmt.Println(str)
}
