package main
import(
	"fmt"
	"time"
)

func main(){

	startTime := time.Now() //start

	for i := 0; i<1000; i++{
		fmt.Println("Hello")
	}
	elapsedTime := time.Since(startTime) //time stamp from start to now

	fmt.Println("실행시간 :",elapsedTime)
}