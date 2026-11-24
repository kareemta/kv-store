package main

import (
	"fmt"
)

func main() {

	mapper := make(map[string]int)
	mapper["KTA"] = 27
	get("KTA", mapper)
	fmt.Println("Hello, World!")
}

func get(userKey string, mapper map[string]int) (int, bool) {
	_, ok := mapper[userKey]
	if ok {
		return mapper[userKey], true
	}
	return 0, false
}

func put(userKey string, userVal int, mapper map[string]int) {
	mapper[userKey] = userVal
}

func del(userKey string, mapper map[string]int) bool {
	_, found := mapper[userKey]
	if found {
		delete(mapper, userKey)
		return true
	}
	return false

}
