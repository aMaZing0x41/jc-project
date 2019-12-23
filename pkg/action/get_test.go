// Copyright 2109 Mario Amato. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package action

import (
	"encoding/json"
	"reflect"
	"testing"
)

var (
	noActions = func() {}
	oneAction = func() {
		AddAction(`{"action": "test", "time": 100}`)
	}
	fiveActons = func() {
		clearActions()
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
		want     []string
	}{
		{"empty", noActions, []string{}},
		{"one", oneAction, []string{`{"action":"test","avg":100}`}},
		//{"five", fiveActons, `[{"action":"test","avg":100}]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.populate()
			var gotObj []string
			got := GetStats()
			err := json.Unmarshal([]byte(got), &gotObj)
			if err != nil {
				t.Errorf("Could not unmarshal response: %v", got)
			}
			if !reflect.DeepEqual(gotObj, tt.want) {
				t.Errorf("GetStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
