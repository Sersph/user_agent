package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//u1 := reflect.ValueOf(new(int)).Pointer()
	//p1 := (*int)(unsafe.Pointer(u1))

	a := int(111)
	fmt.Printf("%x \n", &a)

	ss := []int{111, 111, 111}
	addr := (*reflect.SliceHeader)(unsafe.Pointer(&ss)).Data
	n2 := (*int)(unsafe.Pointer(addr + unsafe.Alignof(int(0))*1))
	n3 := (*int)(unsafe.Pointer(addr + unsafe.Alignof(int(0))*2))
	println(n2, n3)

}
