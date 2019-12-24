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

// GetStats returns a JSON array of all stats along with the average time for each stat.
// Concurrent calls are supported. Creates a copy of the underlying data structure to
// minimize the time that the data structure is locked.
// Each object in the array looks like {"action": "myaction", "avg": 200}.
// Returns a string that represents the JSON array.
func GetStats() string {
	go getStatsInternal()
	return <-statsInfoChan
}

// *** IMPORTANT: calls copyActions which locks the underlying map and copies it.
// The actions map must always use the common mutex for access.
// Copy should make the time that the underlying map is locked consistent.
// Downside is increased memory use. Upside is consistent add performance.
func getStatsInternal() {
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
