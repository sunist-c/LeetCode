package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var input string
	fmt.Println("input array: ")
	fmt.Scanln(&input)

	var res interface{}
	err := json.Unmarshal([]byte(input), &res)
	if err != nil {
		fmt.Println(err)
		return
	}

	printArray(res)
	fmt.Print("\n")
}

func printArray(arr interface{}) {
	switch arr := arr.(type) {
	case []interface{}:
		fmt.Print("{")
		for i, v := range arr {
			if i > 0 {
				fmt.Print(",")
			}
			printArray(v)
		}
		fmt.Print("}")
	case float64:
		fmt.Print(int(arr))
	default:
		fmt.Printf("unexpected type: %T", arr)
	}
}
