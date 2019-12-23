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

type empty struct{}

var (
	actions       = map[string]float32{}
	addChan       = make(chan string)
	statsChan     = make(chan empty)
	statsInfoChan = make(chan string)
)

func init() {
	go func() {
		for {
			select {
			case a := <-addChan:
				addActionInternal(a)
			case <-statsChan:
				getStatsInternal()
			}
		}
	}()
}

func AddAction(a string) error {
	addChan <- a
	//todo: need an error chan
	return nil
}

func addActionInternal(a string) error {
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
		actions[localAction.Action] = averageOfTwo(val, localAction.Time)
	} else {
		actions[localAction.Action] = localAction.Time
	}

	return nil
}

func GetStats() string {
	statsChan <- empty{}
	return <-statsInfoChan
}

func getStatsInternal() {
	//TODO: needs to return different structure

	result, _ := json.Marshal(actions)
	statsInfoChan <- string(result)
}

func averageOfTwo(a, b float32) float32 {
	return (a + b) / 2.0
}
