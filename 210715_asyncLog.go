package main
 
import (
    "os"
    "strconv"
    "time"
)
 
var logChannel chan string

func logSetup(logFile string) {
 
    // 로그 파일이 없으면, 생성한다
    if _, err := os.Stat(logFile); os.IsNotExist(err) {
        f, _ := os.Create(logFile)
        f.Close()
    }
 
    // 로그 채널을 만든다
    logChannel = make(chan string, 100)
 
    // 채널을 통한 비동기 로깅
    go func() {
        // 채널이 닫힐 때까지 메시지 받으면 로깅
        for msg := range logChannel {
            f, _ := os.OpenFile(logFile, os.O_APPEND, os.ModeAppend)
            f.WriteString(time.Now().String() + " " + msg + "\n")
            f.Close()
        }
    }()
}
 
/* 테스트 코드 */
func main() {
    logSetup("./logfile.txt")
 
    go func() {
        for i := 1; i < 20; i++ {
            n := strconv.Itoa(i)
            println(n)
            logChannel <- n
        }
    }()
 
    go func() {
        for i := 100; i < 120; i++ {
            logChannel <- strconv.Itoa(i)
            println(strconv.Itoa(i))
        }
    }()
 
    time.Sleep(1 * time.Second)
    close(logChannel)

    ch1 := make(chan bool)
    ch2 := make(chan bool)
 
    go func(done chan bool) {
        time.Sleep(3 * time.Second)
        done <- true
    }(ch1)
 
    go func(done chan bool) {
        time.Sleep(1 * time.Second)
        done <- true
    }(ch2)
 
    // time.After()는 입력파라미터에 지정된 시간이
    // 지나면 Ready되는 채널을 리턴한다
    timeoutChan := time.After(2 * time.Second)
 
    select {
    case <-ch1:
        println("run1")
 
    case <-ch2:
        println("run2")
 
    // select 문 내에서 타임아웃 체크
    case <-timeoutChan:
        println("timeout")
    }
}