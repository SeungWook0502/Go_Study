package main
import (
	"log"
	"os"
	"io"
)

var myLogger *log.Logger

func main(){

	go_log, err := os.OpenFile("logfile.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil{
		panic(err)
	}
	defer go_log.Close()

	// myLogger = log.New(go_log,  "INFO: ", log.LstdFlags | log.Ltime | log.Lshortfile) //log object생성으로 해당 log 객체의 출력을 지정
	//log.SetOutput(go_log) //표준로거를 실행창에서 파일로 변경

	multiWriter := io.MultiWriter(go_log, os.Stdout) //파일과 실행창을 설정
	log.SetOutput(multiWriter) //파일과 실행창 둘다 출력되도록 설정

	run()

	log.Println("End of Program")
}

func run(){
	log.Print("Start of Program")
}