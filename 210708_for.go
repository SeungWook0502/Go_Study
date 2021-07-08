package main
import "fmt"

func main(){

	sum := 0
    for i := 1; i <= 100; i++ { //삼항연산자
        sum += i
    }
    fmt.Println(sum)
    
    n := 1
    for n < 100 { //go isn't have "while"
        n *= 2      
        if n > 90 {
          break 
        }     
    }
    fmt.Println(n)

    names := []string{"홍길동", "이순신", "강감찬"}
 
	for index, name := range names { //for-range
	    fmt.Println(index, name)
	}

	var a = 1
    for a < 15 {
        if a == 5 {
            a += a
            continue // for루프 시작으로 // 하위무시 바로 다름 loop
        }
        a++
        if a > 10 {
            break  //루프 빠져나옴
        }
    }
    if a == 11 {
        goto END //goto 사용예
    }
    fmt.Println(a)
 
	END: //label
	    fmt.Println("End")

	i := 0
 
		L1: //label
		    for {
		     
		        if i == 3 {
		            break L1 //label break
		        }
		        i++
		    }
	 
    fmt.Println("OK")

}