package main
import "fmt"

//bool
//string -> 생성시 수정될 수 없다.
//int int8 int16 int32 int64
//uint uint8 uint16 uint32 uint64 uintptr
//float32 float64 complex64 complex128
//byte == unit8 -> byte-code
//rune == int32 -> uni-code

func main(){
	rawLiteral := `아리랑\n아리랑` //raw string literal
	interLiteral := "아리랑\n아라리요" //interpreted string literal

	fmt.Println(rawLiteral+"\n")
	fmt.Println(interLiteral)

//Type Conversion
	i := 100
	u := uint64(i)
	f := float32(u)
	fmt.Println(f,u)

	s := "ABC"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(b,s2)
	s = s2
}