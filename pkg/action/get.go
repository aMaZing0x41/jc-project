package action

import "encoding/json"

//TODO: write tests
func GetStats() string {
	go getStatsInternal()
	return <-statsInfoChan
}

func getStatsInternal() {
	//TODO: needs to return different structure

	lock.RLock()
	result, _ := json.Marshal(actions)
	lock.RUnlock()
	statsInfoChan <- string(result)
}
