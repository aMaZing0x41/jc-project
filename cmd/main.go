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
	for i := 0; i < 10; i++ {
		go action.AddAction(`{"action": "test", "time": 10}`)
		go action.AddAction(`{"action": "test", "time": 8}`)
		go action.AddAction(`{"action": "test1", "time": 8}`)
		go fmt.Println(action.GetStats())
	}
	fmt.Println("Done.")
}
