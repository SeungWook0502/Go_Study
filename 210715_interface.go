package main
import( 
	"math"
	"fmt"
)

type Shape interface{
	area() float64
	perimeter() float64
}

type Rect struct {
    width, height float64
}
 
type Circle struct {
	radius float64
}
 
//Rect 타입에 대한 Shape 인터페이스 구현 
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
     return 2 * (r.width + r.height)
}

//Circle 타입에 대한 Shape 인터페이스 구현 
func (c Circle) area() float64 { 
    return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 { 
    return 2 * math.Pi * c.radius
}
func showArea(shapes ...Shape) {
    for _, s := range shapes {
        a := s.area() //인터페이스 메서드 호출
        fmt.Println(a)
    }
}

func a (b interface{}){ //object type
	fmt.Println(b)
}

func main(){

	r := Rect{10.0, 20.0}
	c := Circle{10}

	showArea(r,c)
	a(r)

	var cc interface{} = 1
	d := cc
	e := cc.(int)
	fmt.Println(d) //dynamic type
	fmt.Println(e) //int type
}