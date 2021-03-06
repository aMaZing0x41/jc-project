// Copyright 2109 Mario Amato. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package action

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"testing"
)

//NOTE: the underlying actions map is not cleared by default
// call the clearActions() function to clear it.

var (
	noActions = func() {}
	oneAction = func() {
		AddAction(`{"action": "test", "time": 100}`)
	}
	fiveActions = func() {
		AddAction(`{"action": "test1", "time": 1}`)
		AddAction(`{"action": "test2", "time": 2}`)
		AddAction(`{"action": "test3", "time": 3}`)
		AddAction(`{"action": "test4", "time": 4}`)
		AddAction(`{"action": "test5", "time": 5}`)
	}
)

func TestGetStats(t *testing.T) {
	tests := []struct {
		name     string
		populate func()
		want     []actionAvg
	}{
		// Tests can assume that "want" is sorted by the action string
		{"empty", noActions, []actionAvg{}},
		{"one", oneAction, []actionAvg{actionAvg{Action: "test", Avg: 100}}},
		{
			"five", fiveActions, []actionAvg{
				actionAvg{Action: "test1", Avg: 1},
				actionAvg{Action: "test2", Avg: 2},
				actionAvg{Action: "test3", Avg: 3},
				actionAvg{Action: "test4", Avg: 4},
				actionAvg{Action: "test5", Avg: 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearActions()
			tt.populate()
			var gotObj []actionAvg
			got := GetStats()
			err := json.Unmarshal([]byte(got), &gotObj)
			if err != nil {
				t.Errorf("Could not unmarshal response: %v", got)
			}
			// Since maps are unordered the resulting slice needs to be sorted
			// so we compare apples to apples
			sort.SliceStable(gotObj, func(i, j int) bool {
				return strings.Compare(gotObj[i].Action, gotObj[j].Action) < 0
			})
			if !reflect.DeepEqual(gotObj, tt.want) {
				t.Errorf("GetStats() JSON = %v", got)
				t.Errorf("GetStats() = %v, want %v", gotObj, tt.want)
			}
		})
	}
}

// dummy string so that compiler doesn't optimize out the call below
var dummy string

func benchmarkGetStats(i int, b *testing.B) {

	// Add i entries to the map
	for k := 0; k < i; k++ {
		AddAction(fmt.Sprintf(`{"action": "%v", "time": %v}`, fmt.Sprintf("test%v", i), rand.Float32()))
	}

	var s string
	for n := 0; n < b.N; n++ {
		s = GetStats()
	}
	dummy = s
}

func BenchmarkGetStats1(b *testing.B)       { benchmarkGetStats(1, b) }
func BenchmarkGetStats10(b *testing.B)      { benchmarkGetStats(10, b) }
func BenchmarkGetStats100(b *testing.B)     { benchmarkGetStats(100, b) }
func BenchmarkGetStats1000(b *testing.B)    { benchmarkGetStats(1000, b) }
func BenchmarkGetStats100000(b *testing.B)  { benchmarkGetStats(100000, b) }
func BenchmarkGetStats1000000(b *testing.B) { benchmarkGetStats(1000000, b) }
