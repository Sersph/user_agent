package main

import "fmt"

func main() {
	//A1计算过程1111 1111 + 0000 0001 = 1 0000 0000，因为uint8 只保留8位进制，“1” 就被舍弃，结果为0000 0000，即0.
	//A2计算过程1111 1111 + 1000 0000 = 1 0111 1111，同样保留8位进制，结果为0111 1111，即127。
	var a1 uint8 = 255
	var a2 uint8 = 255
	a1 = a1 + 1
	a2 = a2 + 128
	fmt.Println(" = a1:", a1)
	fmt.Println(" = a2:", a2)

	//0000 0000 - 0000 0001 = 1111 1111;即255
	//0000 0000 - 1111 1111 = 0000 0001;即1
	var b1 uint8 = 0
	var b2 uint8 = 0
	b1 = b1 - 1
	b2 = b2 - b1
	fmt.Println(" = b1:", b1)
	fmt.Println(" = b2:", b2)
}
