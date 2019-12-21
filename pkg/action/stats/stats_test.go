package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	actions["test"] = 75
	err := AddAction(`{"action":"test", "time": 125}`)
	if err != nil {
		t.Errorf("AddAction() returned error: %v", err)
	}
	assert.Equal(t, float32(100), actions["test"])
}

func TestGetStats(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStats(); got != tt.want {
				t.Errorf("GetStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageOfTwo(t *testing.T) {
	type args struct {
		a float32
		b float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"1", args{1, 1}, 1},
		{"2", args{1.5, 1}, 1.25},
		{"3", args{1002, 1}, 501.5},
		{"4", args{3, 4}, 3.5},
		{"5", args{0, 0}, 0},
		{"6", args{0, 1}, .5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfTwo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("averageOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
