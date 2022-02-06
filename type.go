// 2강. 간단한 타입과 함수 정의

/*

string: 스트링(문자열)
int: 정수
byte: 바이트
error: 에러
float32, float64: 실수형
bool: true of false를 나타낸다.

이렇게 대표적인 것만 알아도 이미 반은 끝났다.

여기서 특별한 규칙이 존재하는데
Type(value): 해당 값을 해당 타입으로 바꿔준다.
ex) string(1) // 1이라는 number 값을 String을 바꿔준다.
float64(3) // 3을 실수로 바꿔준다.
*/

package main

import (
	"fmt"
)

// 함수를 알아보자.
func main() {
	// 우리는 이곳에서 var을 써서 var num int = sum(1,3) 이런 형식으로 자바처럼
	// 타입 뒤에 붙여야 한다.
	// 타입을 모르거나 귀찮다면? 그래서 우리는 :=를 사용한다.
	// num := sum(1,3) 이렇게 하면 알아서 타입 해석한다.
	fmt.Println(sum(1, 3))
}

// sumb이라는 함수 옆에 num int와 num2 int라고 나와 있다.
// 여기서 num, num2는 당연히 정수여야 한다.
// 그리고 int {} 부분은 끝에 무슨 타입을 변환할지 알려준다.

// 예를 들어서 func sum(num int) string {} 이렇게 되면
// return num을 하면 string 문자열을 리턴해야 하는데
// int를 반환해주니까 오류가 발생한다.
// 여기서 의아한 사실 함수를 어떻게 가져오는지도 알려줘야 하는데,

// 근데 return을 int형으로 두개를 하는데
// 그럼 당연히 num1,num2 := sum(1,3) 이렇게 하는데
// 근데 go에서는 변수를 선언하고 안쓰면 안된다.
// 그럼 나는 num2만 가져오고 싶으면 어떻게 해야 할까
// 그럼 _로 첫번째 값을 무시했으니 우리는 오류가 안날 것이다.
func sum(num int, num2 int) int {

	return num + num2;
	// 하나만 가져오려면 _ 사용후 return
	// return num2
}
