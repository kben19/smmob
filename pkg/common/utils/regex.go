package utils

import (
	"regexp"
)

func RegexFindMultipleString(pattern []string, target string) []string {
	var results []string
	for _, val := range pattern {
		reg := regexp.MustCompile(val)
		result := reg.FindString(target)
		results = append(results, result)
	}
	return results
}
