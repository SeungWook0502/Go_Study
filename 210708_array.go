package main
import "fmt"

func main(){

	var a [3]int
	a[0] = 1	//start index 0
	a[1] = 2
	a[2] = 4
	fmt.Println(a[2])

	var b = [3]int{1,3,6}
	var c = [...]int{3,5,6} //non-declare size
	var d [3][2]int //n dimension declare
	var e = [3][2]int{
		{1,2},
		{3,4},
		{5,6}, //마지막에도 ,(comma)
	}
	fmt.Println(b,c,d,e)

	//slice
	var f []int //non-declare size
	f = []int{1,2,3}
	fmt.Println(f)

	s:=make([]int,5,10) //data type; lengtg; capacity
	fmt.Println(len(s),cap(s),s) // length,capacity

	var g [][]int
	if g == nil{
		fmt.Println(len(g),cap(g))
		g = [][]int{{100,20,43},{43,22,15}}
	}
	fmt.Println(len(g),cap(g),len(g[1]),cap(g[1]))

	//sub slice
	h := []int{32, 12, 45, 65, 34, 75}
    i := h[2:5] //[start index : end index + 1]
    j := h[2:] // start index ~ 
    k := h[:5] // ~ end index + 1
    
    fmt.Println(i) // index 2,3,4 
    fmt.Println(j) // index 2,3,4,5(end)
	fmt.Println(k) // index 1(start),2,3,4

	//append
	k = append(k,75)
	//j = append(32,12,j) error - first argument must be slice

	fmt.Println(k) // == h
	fmt.Println(append(k,53,261,64)) //multi append

	l := make([]int, 0, 3)
    // 계속 한 요소씩 추가
    for i := 1; i <= 15; i++ {
        l = append(l, i)
        // 슬라이스 길이와 용량 확인
        fmt.Println(len(l), cap(l))
    }
    fmt.Println(l) // 1 부터 15 까지 숫자 출력 

    //append slice to slice
    k = append(k,j...) //after second slice must be add ...
    fmt.Println(k)

    //copy
    source := []int{0, 1, 2}
    target := make([]int, len(source), cap(source)*2)
    fmt.Println(target) // 0, 0, 0
    fmt.Println(len(target), cap(target)) // 3, 6
    copy(target, source)
    fmt.Println(target)  // 0, 1, 2
    fmt.Println(len(target), cap(target)) // 3, 6

    //slice inter Structure
    m := []int{0,1,2,3,4,5}
    fmt.Println(len(m),cap(m)) // pointer > [{*0 1 2 3 4 5}] len:{6} cap:[6]
    m = m[2:5]
    fmt.Println(len(m),cap(m)) // pointer > 0 1 [{*2 3 4} 5] len:{3} cap:[4]

    n := []int{0,1,2,3,4,5,6,7}
    fmt.Println(len(n),cap(n)) // pointer > [{*0 1 2 3 4 5 6 7}] len:{8} cap:[8]
    n = n[2:5]
    fmt.Println(len(n),cap(n)) // pointer > 0 1 [{*2 3 4} 5 6 7] len:{3} cap:[6]
    //capacity == pointer ~ final array point

    
}