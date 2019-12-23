package action

import "testing"

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
