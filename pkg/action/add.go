package action

import (
	"encoding/json"
	"errors"
	"fmt"
)

type action struct {
	Action string  `json:",omitempty"`
	Time   float32 `json:",omitempty"`
}

func AddAction(a string) error {
	go addActionInternal(a)
	return <-errChan
}

func addActionInternal(a string) {
	if a == "" {
		msg := "invalid action: no data"
		errChan <- errors.New(msg)
		fmt.Println(msg)
		return
	}
	var localAction action
	err := json.Unmarshal([]byte(a), &localAction)
	if err != nil {
		errChan <- fmt.Errorf("invalid action: bad json - %w", err)
		fmt.Println(fmt.Sprintf("invalid action: bad json - %v", err))
		return
	}
	if localAction.Action == "" {
		msg := "invalid action: empty action"
		errChan <- errors.New(msg)
		fmt.Println(msg)
		return
	}

	updateAction(localAction.Action, localAction.Time)
	errChan <- nil
}

// *** IMPORTANT ***
// This is the main critical section of the actions map update.
// Acquires a full lock on the actions map and either:
// 1. upserts an exsting key with a new average
// 2. inserts a new action with inital time
func updateAction(action string, time float32) {
	lock.Lock()
	defer lock.Unlock()
	if val, ok := actions[action]; ok {
		actions[action] = averageOfTwo(val, time)
	} else {
		actions[action] = time
	}
}
