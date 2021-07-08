package main
import "fmt"

func main(){

	msg:="hello"
	say_v(msg) //pass by value
	say_r(&msg) //pass by reference
	fmt.Println("3-",msg)
	say_vf(msg)
	say_vf("Hello","SeungWook"," ","Jeon")
	fmt.Println("5-",sum(1,4,3,25))
	sum_n, sum_c := sum_m(1,4,3,25)
	fmt.Println("6-",sum_n,sum_c)
	sum_n, sum_c = sum_dr(1,4,3,25)
	fmt.Println("7-",sum_n,sum_c)
}

func say_v(msg string) { //pass by value
    println("1-",msg)
}
func say_r(msg *string){ //pass by reference
	println("2-",*msg)
	*msg = "HELLO"
}
func say_vf(msg ...string){ //variadic function 가변인자함수 개수를 지정하지 않고 받음
	for _, s := range msg {
        println("4-",s)
    }
}
func sum(nums ...int) (int){ //return : func func_name(argument)(return type)
	s := 0
	for _,n := range nums{
		s+=n
	}
	return s
}
func sum_m(nums ...int) (int,int){ //multiple return
	s := 0
	c := 0
	for _,n := range nums{
		s+=n
		c++
	}
	return s,c
}
func sum_dr(nums ...int) (total int, count int) { //declare return values : func func_name(argument)(return_value_name, return_type)
    for _, n := range nums {
        total += n
    }
    count = len(nums)
    return
}