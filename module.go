// 4강. 모듈

// 이번에는 모듈에 대해 알아보자.
// 먼저 go라는 아이는 다시 말하자면, 매우 작은 언어여서
// js가 for 문만 10가지가 넘을때
// go는 for문이 단 하나밖에 없다.

// 1. node.js를 배웠다면 모듈을 불러욜때 require("")이 익숙할 테지만
// go 언어에서는 그렇게 쉽게 작동하지 않는다.
// 모둘을 가져올려면 go get링크로 가져와야 한다.

/* 2. 모듈은 import (
 "링크"
)
라는 방식으로 임포트 해야한다.
모듈을 불러와놓고 안쓰면 오류가 난다.
go에서는 변수를 정의하고 안쓰면 오류가 난다
해당 구조처럼, 모듈을 임포트해놓고 안쓰면 안됨
*/
package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// 해당 코드릴 실행할 시
// 여기서 결과는 뭔가 알아들을 수 없는 문자들이 있을텐데, 에러가 아니면 정상

// 여기서 따로 go get을 안해도 되는 내장 모듈이 있는데
// 따로 내장 모듈에 대해 정확하게 알고 싶을때 구글검색을 해도 되지만 그닥
// 오늘 내용을 정리하자면

/*
  1. 대부분의 모듈들은 go get으로 해야 하지만 내장 모듈들은 따로 go get으로
  안해도 된다.
  2. 임포트 할때는
  import (
	    "링크" // 내장 모듈이 아닐때
		"모듈명" // 내장 모듈일 때
  )
  3. 이 모든 것을 하기전 go mod init 모듈명, go mod tidy로 먼저해주고 하면 된다.
*/
func main() {
	fmt.Println(echo.New())
}