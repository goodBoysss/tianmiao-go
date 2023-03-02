package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		HelloWold       string `json:"hello_wold"`
		LightWeightBaby string
	}
	var a = Person{HelloWold: "chenqionghe", LightWeightBaby: "muscle"}
	res, _ := json.Marshal(a)

	var dat map[string]interface{}
	json.Unmarshal(res, &dat)

	fmt.Println(dat)
}
