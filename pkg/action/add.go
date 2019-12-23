// Copyright 2109 Mario Amato. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package action

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Internal type that represents an action
type action struct {
	Action string  `json:",omitempty"`
	Time   float32 `json:",omitempty"`
}

// AddAction takes a json string representing an "action".
// A valid action looks like {"action": "myaction", "time": 100}
// If the action is new, it is parsed and tracked.
// If the action has been added before, the average time for that action is updated.
// It returns an error if the json is invalid; otherwise nil.
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
