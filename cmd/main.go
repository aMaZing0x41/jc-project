// Copyright 2109 Mario Amato. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/amazing0x41/jc-project/pkg/action"
)

var keys = []string{"test1", "test2", "test3", "test4"}

const NUM_ITERS = 1000000

// This is a simple test harness for the action package NUM_ITERS results in many go routines.
// Keep the number reasonable. The harness will wait for a bit every few iterations in order
// to keep the number of active routines manageable. This is most useful to use the -race flag
// on to see if go tools can detect any race conditions.
//
// FUTURE: it may be useful to have similar logic baked into the normal unit/benchmark tests
// so that we can ensure code performance and more regularly test for races using this more complex logic.
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
			time.Sleep(time.Millisecond * 1)
		}
	}
	fmt.Println(action.GetStats())
	fmt.Println("Done.")
}
