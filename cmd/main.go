package main

import (
	"fmt"

	"github.com/amazing0x41/jc-project/pkg/action"
)

func main() {

	err := action.AddAction("{}")
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: %v", err))
	}
	action.AddAction(`{"action": "test", "time": 10}`)
	action.AddAction(`{"action": "test", "time": 8}`)
	fmt.Println(action.GetStats())
	fmt.Println("Done.")
}
