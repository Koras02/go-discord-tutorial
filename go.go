// 5강 go 라는 문법
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	// 그러면 이번에는 i < 5를 2로 바꿔주면 어떻게 될까?
	// i 가 2일때는 즉 실행이 종료될쯤. go say("1") 문법이 주석처리 처럼 무시당하고
	// 그전 까지는 잘 작동 된다는 것이다.
	for i := 0; i < 5; i++ {
		// 먼저 time.Sleep으로 100밀리 세컨드 동안 기다리다 s를 프린트하는 반복문이고
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
위 코드를 실행하면
hello
world
world
world
hello
hello
world
wrold
hello
hello
이런 식으로 나온다.
*/

// 그럼 아래에서 의문을 가질점은, 처음 결과로 봐서는 무엇을 하는 건지 모른겟는데
// 또 go say("world") 한줄 만으로도 작동하지 않는 문법이고
// 이 say("hello")라는 문법을 A, go say("wolrd")을 B라고 가정해보면

// 그럼 해당 실행과정을 나타낼때
// A 실행 => B => A => B => B => A => A => B => B => A 이렇게 작동된 다는 것을
// 우리는 알 수 있다.

func main() {
	// 그런데 여기서 못보던 go라는 것이 나온다.
	// 해당 코드의 뜻은 무엇일까
	// 이 go 문법은 혼자 작동을 못한다.
	// 만약 say("hello")가 없다면 아무것도 실행하지 못한다.
	go say("world")
	say("hello")
}
