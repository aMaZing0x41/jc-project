// Copyright 2109 Mario Amato. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TODO: tests should start with a clean actions map every time
func TestAddAction(t *testing.T) {
	type args struct {
		action string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"empty", args{""}, true},
		{"empty json", args{"{}"}, true},
		{"bad json", args{`{action: ", "time": 300}`}, true},
		{"action", args{`{"action": "test", "time": 300}`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddAction(tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("AddAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddActionMany(t *testing.T) {
	err := AddAction(`{"action":"test_many", "time": 75}`)
	if err != nil {
		t.Errorf("AddAction() returned error: %v", err)
	}
	err = AddAction(`{"action":"test_many", "time": 125}`)
	if err != nil {
		t.Errorf("AddAction() returned error: %v", err)
	}
	lock.RLock()
	defer lock.RUnlock()
	assert.Equal(t, float32(100), actions["test_many"])
}
