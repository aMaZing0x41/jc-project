package main

import (
	"fmt"

	"github.com/amazing0x41/jc-project/pkg/action/stats"
)

func main() {

	err := stats.AddAction("{}")
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: %v", err))
	}
	stats.AddAction(`{"action": "test", "time": 10}`)
	stats.AddAction(`{"action": "test", "time": 8}`)
	fmt.Println(stats.GetStats())
	fmt.Println("Done.")
}
