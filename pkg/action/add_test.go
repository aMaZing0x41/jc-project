// Copyright 2109 Mario Amato. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//NOTE: the underlying actions map is not cleared by default
// call the clearActions() function to clear it.

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
		clearActions()
		t.Run(tt.name, func(t *testing.T) {
			if err := AddAction(tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("AddAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddActionMany(t *testing.T) {
	clearActions()
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

func benchmarkAddAction(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for k := 0; k < i; k++ {
			AddAction(`{"action": "test", "time": 100}`)
		}
	}
}

func BenchmarkAddAction1(b *testing.B)       { benchmarkAddAction(1, b) }
func BenchmarkAddAction10(b *testing.B)      { benchmarkAddAction(10, b) }
func BenchmarkAddAction100(b *testing.B)     { benchmarkAddAction(100, b) }
func BenchmarkAddAction1000(b *testing.B)    { benchmarkAddAction(1000, b) }
func BenchmarkAddAction100000(b *testing.B)  { benchmarkAddAction(100000, b) }
func BenchmarkAddAction1000000(b *testing.B) { benchmarkAddAction(1000000, b) }
