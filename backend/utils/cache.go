package utils

import "sync"

var (
	queryCache = make(map[string][]string)
	cacheMutex sync.RWMutex
)

func GetCachedQuery(key string) ([]string, bool) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	val, found := queryCache[key]
	return val, found
}

func SetCachedQuery(key string, value []string) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	queryCache[key] = value
}
