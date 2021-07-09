package main
import "fmt"

func main(){

	//pointer value
	var a map[int]string //map[key type]value type
	a = make(map[int]string)

	b := map[string]string{ //initialize
		"a":"A",
		"b":"B",
		"c":"C",
	}

	a[32] = "1" //index X | key O
	a[6242] = "2"

	c := a[6242]
	fmt.Println(c)

	d := a[4321] //missing key is return nil or zero value
	fmt.Println(d)

	delete(a, 32) //remove

	//check map value
	e, f := b["a"] //e - value, f - bool

	if !f {
		fmt.Println("No",e)
	}else{
		fmt.Println("Yes",e)
	}

	for g, h := range b{
		fmt.Println(g,h)
	}

}