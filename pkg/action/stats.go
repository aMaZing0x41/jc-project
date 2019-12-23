package action

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type action struct {
	Action string  `json:",omitempty"`
	Time   float32 `json:",omitempty"`
}

//TODO: rearrange this. move to common area; split add and get stats
type empty struct{}

var (
	actions       = map[string]float32{}
	lock          = sync.RWMutex{}
	errChan       = make(chan error)
	statsInfoChan = make(chan string)
)

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

	//TODO: move this into separate func and defer unlock
	lock.Lock()
	if val, ok := actions[localAction.Action]; ok {
		actions[localAction.Action] = averageOfTwo(val, localAction.Time)
	} else {
		actions[localAction.Action] = localAction.Time
	}
	lock.Unlock()

	errChan <- nil
}

func GetStats() string {
	go getStatsInternal()
	return <-statsInfoChan
}

func getStatsInternal() {
	//TODO: needs to return different structure

	lock.RLock()
	result, _ := json.Marshal(actions)
	lock.RUnlock()
	statsInfoChan <- string(result)
}

func averageOfTwo(a, b float32) float32 {
	return (a + b) / 2.0
}
