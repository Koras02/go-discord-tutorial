// 파일명과 동일한 패키지를 만들고, 나중에 go가 인식하고 임포하게 도와주는 packace 이다.
package main

// 이 둘중 fmt라는 모듈을 임포트 해준다.
// 여기서 fmt 라는 모듈은 다른언어 (파이썬, 자바 스크립트) 등등 내장 모듈이라는 특성도 있어 따로 fmt모듈은 설치
// 하지 않아도 된다.
// import "fmt";
import (
	"fmt"
)

// 함수 -> func main() 함수는 Go를  실행할 때 원하는 값이 출력되고 싶을때
// 이 함수안에 집어넣어야 한다.
func main() {

	// fmt 라는 모듈안에 Println이라는 함수 및 속성이 있어
	// fmt.println으로 말할 내용을 콘솔에 출력 가능
	// 이것 말고도 fmt.Sprinf 등등 여러가지 일들이 있으나, 가장 기초적인 것은 fmt.Println
	// fmt.Println("Hello, world!")
	fmt.Println("Ksk 최고")
}
