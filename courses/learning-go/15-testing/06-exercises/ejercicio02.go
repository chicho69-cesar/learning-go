package main

import "strings"

func Count(str string) int {
	strArray := strings.Fields(str)
	return len(strArray)
}

func UseCount(str string) map[string]int {
	strArray := strings.Fields(str)
	dictionary := make(map[string]int)
	for _, value := range strArray {
		dictionary[value]++
	}
	return dictionary
}
