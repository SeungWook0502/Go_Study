package main

import(
	"io"
	"os"
	"fmt"
)

func main(){

	fi, err := os.Open("C://Users//wjstm//OneDrive//바탕 화면//Go_Study//1.txt") //Open
	if err != nil {
        panic(err)
    }
    defer fi.Close()

    fo, err := os.Create("C://Users//wjstm//OneDrive//바탕 화면//Go_Study//2.txt") //Create
    if err != nil {
        panic(err)
    }
    defer fo.Close()

    buff := make([]byte, 1024) //file content variable

    fa, err := os.OpenFile("C://Users//wjstm//OneDrive//바탕 화면//Go_Study//3.txt", os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil{
    	panic(err)
    }
    defer fa.Close()


    for {
        // 읽기
        cnt, err := fi.Read(buff) //
        if err != nil && err != io.EOF {
            panic(err)
        }
        fmt.Println(cnt)
 		
        // 끝이면 루프 종료
        if cnt == 0 {
    		fmt.Fprintf(fa,"%s\n",buff)
            break
        }
 
        // 쓰기
        _, err = fo.Write(buff[:cnt])
        if err != nil {
            panic(err)
        }
    }
}