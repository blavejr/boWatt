package utils

import (
	"os"
	"sync"
)

var store = make(map[string]string)
var lock = sync.RWMutex{}

func ReadFileContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SaveToMemory(filename, content string) {
	lock.Lock()
	defer lock.Unlock()
	store[filename] = content
}

func GetFromMemory(filename string) string {
	lock.RLock()
	defer lock.RUnlock()
	return store[filename]
}
