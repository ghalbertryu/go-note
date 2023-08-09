package concurrent

import "sync"

var (
	normalMap = make(map[int]bool)
	syncMap   = new(sync.Map)
)

func Set() {
	normalMap[1] = true
}

func Get() bool {
	return normalMap[1]
}

func SetSyncMap() {
	syncMap.Store(1, true)
}

func GetSyncMap() bool {
	value, ok := syncMap.Load(1)
	if ok {
		return value.(bool)
	}

	return false
}
