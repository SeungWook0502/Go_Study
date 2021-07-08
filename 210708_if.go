package main
import "fmt"

func main(){

	//if
	k := "aaa"
	i := "aaa"
	if k==i{
		fmt.Println("s")
	}else if k!=i{
		fmt.Println("n")
	}

	if a := 2*5; a<10{ //optional statement
		fmt.Println("s2")
	}else {
		fmt.Println("n2")
	}

	//switch
	b := 3
	switch b { //non break
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	grade(86)
	check(1)
	do(21)
	do("aa")

}

func grade(score int){ //non expression == if-else
	switch{
	case score >= 90:
		fmt.Println("A2")
	case score >= 80:
		fmt.Println("B2")
	case score >= 70:
		fmt.Println("C2")
	default:
		fmt.Println("F2")
	}
}

func check(val int) {
	switch val {
	case 1:
	    fmt.Println("1 이하")
	    fallthrough //==unbreak
	case 2:
	    fmt.Println("2 이하")
	    fallthrough
	case 3:
	    fmt.Println("3 이하")
	    fallthrough
	default:
	    fmt.Println("default 도달")
	}
}
func do(i interface{}) { //reference https://iamjjanga.tistory.com/46
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}