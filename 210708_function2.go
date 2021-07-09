package main
import "fmt"

type calculator func(int, int)(int) //define origin func/struc/interface/custom type

func main(){

	//anonymous function
	sum := func(n ...int) int {
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    } //assignment function for value
 
    result := sum(1, 2, 3, 4, 5)
    fmt.Println(result)

    //function as parameter
    add := func(i int, j int) int {
        return i + j
    }
 
    // add 함수 전달
    r1 := calc(add, 10, 20)
    fmt.Println(r1)
 
 	// 일급함수
    // 직접 첫번째 파라미터에 익명함수를 정의함
    r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
    fmt.Println(r2)


    r3 := calc2(func(x int, y int) int { return x - y }, 10, 20)
    fmt.Println("r3",r3)

    //Closure
    next := nextValue()
 
    fmt.Println(next())  // 1
    fmt.Println(next())  // 2
    fmt.Println(next())  // 3

    anotherNext := nextValue()
    fmt.Println(anotherNext()) // 1 다시 시작
    fmt.Println(anotherNext()) // 2
    //assignment 될 때 마다 값이 0이된다 (initialize)

}
func calc(f func(int, int) int, a int, b int) int {
    result := f(a, b)
    return result
}
func calc2(f calculator, a int, b int)(int){
    result := f(a, b)
    return result

}
func nextValue() func() int {
    i := 0
    return func() int {
        i++ //외부 변수 사용
        return i
    }
}