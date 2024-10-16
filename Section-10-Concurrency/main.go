// package main

// import (
// 	"fmt"
// 	"time"
// )

// func greet(phrase string) {
// 	fmt.Println("Hello!", phrase)
// }

// func slowGreet(phrase string) {
// 	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
// 	fmt.Println("Hello!", phrase)
// }

//	func main() {
//		greet("Nice to meet you!")
//		greet("How are you?")
//		slowGreet("How ... are ... you ...?")
//		greet("I hope you're liking the course!")
//	}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func sendMsg(msg string) {
	pause()
	fmt.Println(msg)
}

func main() {
	sendMsg("hello") // sync

	go sendMsg("test1") // async
	go sendMsg("test2") // async
	go sendMsg("test3") // async

	sendMsg("main") // sync

	time.Sleep(2 * time.Second)
}
