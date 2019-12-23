package action

import "sync"

var (
	actions       = map[string]float32{}
	lock          = sync.RWMutex{}
	errChan       = make(chan error)
	statsInfoChan = make(chan string)
)

func averageOfTwo(a, b float32) float32 {
	return (a + b) / 2.0
}
