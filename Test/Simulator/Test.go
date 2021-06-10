package main

import "fmt"

//var(
//	A []byte
//	B []byte
//	C []byte
//)

func main() {
	//fmt.Println("copy的不等长切片的用法:")
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice1,slice2)
	//copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	//fmt.Println(slice1,slice2)
	//
	//fmt.Println("切片和Map的copy是浅copy:")
	//slice3 := []byte{1,2,2,2,2,2}
	//var slice4 [5]byte
	//copy(slice4[0:],slice3)
	//fmt.Println(slice3,slice4)
	//slice3 = []byte{1}
	//fmt.Println(slice3[0:],slice4)

	//浅拷贝
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a地址1:%p\n", a)
	b := a[0:3]
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)
	fmt.Printf("a地址2:%p\n", a)
	fmt.Printf("b地址:%p\n", b)
	b[0] = 100
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)
	c := make([]int, 3, 3)
	c = a[0:3]
	fmt.Println("c is:", c)
	fmt.Printf("c地址:%p\n", c)

	//深拷贝

	// 当元素数量超过容量
	// 切片会在底层申请新的数组
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1
	fmt.Printf("s1地址%p,s2地址%p\n", s1, s2)
	s1 = append(s1, 100)
	s2[0] = 200
	fmt.Printf("s1地址%p,s2地址%p\n", s1, s2)
	fmt.Println("s1 is:", s1)
	fmt.Println("s2 is:", s2)
	// copy 函数提供深拷贝功能
	// 但需要在拷贝前申请空间
	s3 := make([]int, 4, 4)
	s4 := make([]int, 5, 5)
	fmt.Println(copy(s3, s1)) //4
	fmt.Println(copy(s4, s1)) //5
	s3[0] = 300
	s4[0] = 400
	fmt.Println("s1 is:", s1)
	fmt.Println("s3 is:", s3)
	fmt.Println("s4 is:", s4)

}
