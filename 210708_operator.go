package main
import "fmt"

func main(){
	c := (4+3) / 5;
	i := 1
	i++; //increasing operator
	a:=true
	b:=false
	fmt.Println(a == b,a!=b,c>=i,a&&b,a||b) //arithmetic logic operator
	fmt.Println((4&3)<<1) //Bitwise operator
	
	//assignment operator
	f := 100
	fmt.Println(f)
	f *= 10
	fmt.Println(f)
	f >>= 2
	fmt.Println(f)
	f |= 1
	fmt.Println(f)

	p := &f //assignment address
	fmt.Println(*p) //pointer operator
}