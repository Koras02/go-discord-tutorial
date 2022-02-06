// 6강 defer문법
package main

import (
	"fmt"
)

func main() {
	// defer 문법을 보자
	// 이 문법은 b출력후 c 출력후 a가 출려되는 코드다.
	// defer문법은 defer 단어 이후 코드 한줄을 맨 마지막에 실행시킨다.

	//defer문법은 이게 끝이다
	defer fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")
}
