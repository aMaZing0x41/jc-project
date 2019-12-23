// Copyright 2109 Mario Amato. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package action

import "encoding/json"

// Internal type that represents an action with avg time
type actionAvg struct {
	Action string  `json:"action,omitempty"`
	Avg    float32 `json:"avg,omitempty"`
}

//TODO: write tests
func GetStats() string {
	go getStatsInternal()
	return <-statsInfoChan
}

func getStatsInternal() {
	//TODO: needs to return different structure
	actionsCopy := copyActions()
	stats := []actionAvg{}
	for k, v := range actionsCopy {
		stats = append(stats, actionAvg{Action: k, Avg: v})
	}

	result, _ := json.Marshal(stats)
	statsInfoChan <- string(result)
}

func copyActions() map[string]float32 {
	lock.RLock()
	defer lock.RUnlock()
	copy := make(map[string]float32, len(actions))
	for k, v := range actions {
		copy[k] = v
	}
	return copy
}
