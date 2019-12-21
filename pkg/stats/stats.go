package stats

import (
	"encoding/json"
	"errors"
	"fmt"
)

type action struct {
	Action string  `json:",omitempty"`
	Time   float32 `json:",omitempty"`
}

var actions = map[string]float32{}

func AddAction(a string) error {
	if a == "" {
		return errors.New("invalid action: no data")
	}
	var localAction action
	err := json.Unmarshal([]byte(a), &localAction)
	if err != nil {
		return fmt.Errorf("invalid action: bad json - %w", err)
	}
	if localAction.Action == "" {
		return errors.New("invalid action: empty action")
	}

	if val, ok := actions[localAction.Action]; ok {
		fmt.Println("key exists")
		val = (val + localAction.Time) / 2
		actions[localAction.Action] = val
	} else {
		actions[localAction.Action] = localAction.Time
	}

	return nil
}

func GetStats() string {
	result, _ := json.Marshal(actions)
	return string(result)
}
