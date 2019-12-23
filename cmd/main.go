package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/amazing0x41/jc-project/pkg/action"
)

var keys = []string{"test1", "test2", "test3", "test4"}

//TODO: doc this
const NUM_ITERS = 100000

// This is a simple test harness for the action package
// NUM_ITERS results in many go routines. Keep the number
// reasonable. The harness will wait for a bit every few
// iterations in order to keep the number of active routines
// manageable.
func main() {
	err := action.AddAction("{}")
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: %v", err))
	}
	go action.AddAction("{")
	go action.AddAction(`{"happy":true}`)

	for i := 0; i < NUM_ITERS; i++ {
		go action.AddAction(fmt.Sprintf(`{"action": "%v", "time": %v}`, keys[rand.Intn(len(keys))], rand.Float32()))
		if i%10000 == 0 {
			go fmt.Println(action.GetStats())
			time.Sleep(time.Millisecond * 5)
		}
	}
	fmt.Println(action.GetStats())
	fmt.Println("Done.")
}
