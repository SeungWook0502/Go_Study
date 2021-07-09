package main
 
import "fmt"
 
// struct define
type person struct {
    name string
    age  int
}
type human struct{
	e int
	f string
}
type animal struct{
	h map[string]string
}
 
func main() {
    // person 객체 생성
    a := person{}
     
    // 필드값 설정
    a.name = "Lee" // .
    a.age = 10
     
    fmt.Println(a)

    var b person 
	b = person{"Bob", 20}
	b2 := person{name: "Sean", age: 50}
	fmt.Println(b.name,b2.age)

	c := new(person) //instance object
	c.name = "Lee"

	g := new_human()
	fmt.Println(g.f)

	j := new_animal()
	j.h["cat"] = "dog"
	fmt.Println(j.h["cat"])
}

//constructor 생성자
func new_human() *human{
	d := human{}
	d.e = 0
	d.f = "a"
	return &d
}
func new_animal() *animal{
	i := animal{} //assignment object
	i.h = map[string]string{}
	return &i
}