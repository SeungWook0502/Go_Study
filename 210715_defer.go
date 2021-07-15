package main

import (
	"os"
	"fmt"
)

func main(){
	f, err := os.Open("C://Users//wjstm//eclipse-workspace//Today_News//Naver_News_URL_Analysis.txt")
	if err != nil{
		panic(err) // 실행시 다른 panic 함수들을 모두 실행한 후 즉시 리턴한다.
	}

	defer f.Close() //가장 마지막에 무조건 실행 == finally

	bytes := make([]byte, 1024)
	fmt.Println(f.Read(bytes))
	fmt.Println(len(bytes))

	openFile("C://Users//wjstm//eclipse-workspace//Today_News//Naver_News_URL_Analyis.txt")
	openFile("C://Users//wjstm//eclipse-workspace//Today_News//Naver_News_URL_Analysis.txt")
	fmt.Println("Done")	
}

func openFile(fn string) {
    // defer 함수. panic 호출시 실행됨
    defer func() {
        if r := recover(); r != nil { //panic상태를 다시 정상으로 되돌리는 method 때문에 이후 command-line이 실행된다.
            fmt.Println("OPEN ERROR", r)
        }
    }()
 
    f, err := os.Open(fn)
    if err != nil {
        panic(err)
    }
 
    defer f.Close()
}