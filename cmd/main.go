package main

import (
	"fmt"

	"github.com/amazing0x41/jc-project/pkg/stats"
)

func main() {

	err := stats.AddAction("{}")
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: %v", err))
	}
	stats.AddAction(`{"action": "test", "time": 10}`)
	stats.AddAction(`{"action": "test", "time": 8}`)
	s := stats.GetStats()
	fmt.Println(s)
	fmt.Println("Done.")
}