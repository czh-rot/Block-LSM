package main

import (
	"bytes"
	"fmt"
)

func main() {
	// l == 108
	for i:=0;i<=21;i++{
		fmt.Println(i/10)
	}
	str := []byte("l")
	fmt.Println(str)
	key := append(str,[]byte{1,2}...)
	fmt.Println(key)
	fmt.Println(key[:1])
	fmt.Println(bytes.Compare(key[:1],[]byte("l")) == 0)
}
