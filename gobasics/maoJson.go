package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"english", "math", "chinese"}

	m["finished"] = false

	res, err := json.Marshal(m)
	if err == nil {
		fmt.Println(string(res))
	}

	var m2 map[string]interface{}
	err2 := json.Unmarshal(res, &m2)
	fmt.Println(m2, err2)

	var s string

	s = m2["company"].(string)
	fmt.Println("s = ", s)
	for _, v := range m2 {
		switch value := v.(type) {
		case string:
			s = value
			fmt.Println("是string类型", s)
		}
	}

}
