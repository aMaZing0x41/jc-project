package stats

import "testing"

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
