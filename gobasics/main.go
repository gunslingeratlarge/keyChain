package main

import (
	"fmt"
)

type student struct {
	name   string
	number int
}

func main() {
	var objects = make([]interface{}, 2)
	objects[0] = 1
	objects[1] = "hello world"
	fmt.Println(len(objects), cap(objects))

	objects = append(objects, student{"mike", 123})
	fmt.Println(len(objects), cap(objects))

	for _, data := range objects {

		if value, ok := data.(int); ok == true {
			fmt.Println("该值是int类型", value)
		} else if ok == false {
			fmt.Println("该值不是int类型")
		}
	}
}
