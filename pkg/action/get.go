// Copyright 2109 Mario Amato. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

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
