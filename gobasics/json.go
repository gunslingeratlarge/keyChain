package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id       int `json:"id"`
	Subjects []string
	Gender   string
	Finished bool
}

func main() {
	s := Student{1, []string{"math", "chinese", "english"}, "male", false}
	res, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		fmt.Println(string(res))
	}
}
