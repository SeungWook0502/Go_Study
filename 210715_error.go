package main
import(
	"log"
	"os"
	"fmt"
)

func main(){
	r, err := os.Open("C://temp//1.txt")
	if err != nil{
		log.Fatal(err.Error())
	}
	fmt.Println(r.Name())

	_,err2 := aaa(3)
	switch err2.(type){
	default:
		fmt.Println("ok")
	case MyError:
		log.Print("Log my error")
	case error:
		log.Fatal(err2.Error())
	}
}

func aaa(a int) string{
	return a.(string)
}