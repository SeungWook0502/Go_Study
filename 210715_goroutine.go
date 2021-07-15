package main
 
import (
    "fmt"
    "time"
    "sync"
    "runtime"
)
 
func say(s string) {
    for i := 0; i < 40; i++ {
        fmt.Println(s, "***", i)
    }
}
 
func main() {
    // 함수를 동기적으로 실행
    say("Sync")
 
    // 함수를 비동기적으로 실행 == thread
    go say("Async1")
    go say("Async2")
    go say("Async3")
 
    // 3초 대기
    time.Sleep(time.Second * 3)

    ////anonymous function
    // WaitGroup 생성. 2개의 Go루틴을 기다림.
    var wait sync.WaitGroup
    wait.Add(2)
 
    // 익명함수를 사용한 goroutine
    go func() {
        defer wait.Done() //끝나면 .Done() 호출
        fmt.Println("Hello")
    }()
 
    // 익명함수에 파라미터 전달
    hi := "Hi"
    go func(msg string) {
        defer wait.Done() //끝나면 .Done() 호출
        fmt.Println(msg)
    }(hi)
 
    wait.Wait() //Go루틴 모두 끝날 때까지 대기 == join

    //runtime.GOMAXPROCS(4)//4개의 logical cpu사용 declare
}