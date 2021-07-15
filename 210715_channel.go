package main
import (
	"fmt"
	"time"
)

func main() {
  	// 정수형 채널을 생성한다 
  	ch := make(chan int) //make를 통해서 생성해야 한다.
 
	go func() {
		ch <- 123   //송신
	}()
	
	var i int
	i = <- ch  //수신
	println(i)

	done := make(chan bool)
	go func() {
	    for i := 0; i < 10; i++ {
	        fmt.Println(i)
	    }
	    done <- true
	}()

	// 위의 Go루틴이 끝날 때까지 대기	
	<-done	
	fmt.Println("asdf")

	//Unbuffered Channel
	// c := make(chan int)
	// c <- 1   //수신루틴이 없으므로 데드락 
	// fmt.Println(<-c) //코멘트해도 데드락 (별도의 Go루틴없기 때문)

	//Buffered Channel
	bch := make(chan int, 1)
    bch <- 101	//수신자가 없더라도 보낼 수 있다.
    fmt.Println(<-bch)
    fch := make(chan string, 1)
    sendChan(fch)
    receiveChan(fch)

    //check channel
    cch := make(chan int,2)
    cch <- 1
    cch <- 2
    close(cch) //close channel
    for {
    	cch_c,success := <- cch
    	if !success{
    		break
    	}
    	fmt.Println(cch_c)
    }
    //usage range for 
    cch2 := make(chan int,2)
    cch2 <- 1
    cch2 <- 2
    close(cch2) //close channel
    for i := range cch2{
    	fmt.Println(i)
    }

    //channel select
    done1 := make(chan bool)
    done2 := make(chan bool)
 
    go run1(done1)
    go run2(done2)
 
EXIT:
    for {
        select { //case == channel name
        case <-done1:
            println("run1 완료")
 
        case <-done2:
            println("run2 완료")
            break EXIT
        }
    }
}

//function channel type declare
func sendChan(ch chan<- string) {
    ch <- "Data"
    // x := <-ch // 에러발생
}
 
func receiveChan(ch <-chan string) {
    data := <-ch
    fmt.Println(data)
}

func run1(done chan bool) {
    time.Sleep(1 * time.Second)
    done <- true
}
 
func run2(done chan bool) {
    time.Sleep(2 * time.Second)
    done <- true
}