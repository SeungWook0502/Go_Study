package main
import "fmt"

type a struct{
	b int
	c int
}

func main(){

	e := a{2,6}
	f := a{2,6}

	g := e.spell()
	h := f.spell2()

	fmt.Println(g,e.b)
	fmt.Println(h,f.b)

}

func (d a) spell() int{ //value receiver
	d.b++
	return d.b + d.c
}
func (d *a) spell2() int{ //pointer receiver
	d.b++
	return d.b + d.c
}