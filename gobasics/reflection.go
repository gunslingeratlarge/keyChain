package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

//insert into order values(456, 56)
//insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")
//unsupported type

func createQuery(q interface{}) {
	fmt.Printf("%T\n", q)
	var v reflect.Value = reflect.ValueOf(q)
	t := reflect.TypeOf(q)
	fmt.Printf("%T\n", t)
	for i := 0; i < v.NumField(); i++ {

		switch v.Field(i).Kind() {
		case reflect.String:
			fmt.Println(v.Field(i).String())
		case reflect.Int:
			fmt.Println(v.Field(i).Int())

		}

	}
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	//i := 90
	//createQuery(i)

}
