// 함수의 모든것
package main

import (
	"fmt"
)

// 함수의 모든것
func main() {
	// 3가지 문자열이 들어갈때 3글자 또는 다른 것들을
	fmt.Print(echo("hello", "world", "going!!"))
}

// echo 하고 싶다면,그것이 1000개나 1만개가 된다면 곤란한데
// 이런식으로 ...string으로 몇개가 오든 상관 없다.
// 1000 개도 가능하고 100개도 가능하다.
// 실행 결과는 [hello world going!!] 이런 식으로 나온다.
// 그리고 여기서는 []string 형태로 리스트를 만드는 데 문자열만 들어가는 리스트로
// 범위를 줄인 것
func echo(strings ...string) []string {
	return strings
}
