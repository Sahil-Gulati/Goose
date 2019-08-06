package goose

import (
	"fmt"
	"regexp"
	"strings"
)

func Isset(data interface{}, key interface{}) bool {
	var isset bool
	switch data.(type) {
	case map[string]string:
		_, isset = data.(map[string]string)[key.(string)]
		break
	case map[string]interface{}:
		_, isset = data.(map[string]interface{})[key.(string)]
		break
	case map[int]string:
		_, isset = data.(map[int]string)[key.(int)]
		break
	case map[int]interface{}:
		_, isset = data.(map[int]interface{})[key.(int)]
		break
	case []string:
		if key.(int) < len(data.([]string)) {
			isset = true
		} else {
			isset = false
		}
	default:
		_, isset = data.(map[int]interface{})[key.(int)]
		break
	}
	return isset
}
func contains(strings []string, piece string) bool {
	for _, str := range strings {
		if str == piece {
			return true
		}
	}
	return false
}
func getDynamics(url string) []string {
	dynamics := []string{}
	regex := regexp.MustCompile(`\{[a-zA-Z_]+\}`)
	matches := regex.FindAllString(url, -1)
	for _, match := range matches {
		regex = regexp.MustCompile(`([a-zA-Z_]+)`)
		matches = regex.FindAllString(match, -1)
		dynamics = append(dynamics, matches[0])
	}
	return dynamics
}
func convertDyanmicURLToRegex(url string) string {
	regex := regexp.MustCompile(`\{[a-zA-Z_]+\}`)
	matches := regex.FindAllString(url, -1)
	for _, match := range matches {
		url = strings.Replace(url, match, "([a-zA-Z0-9_.]+)", -1)
	}
	url = strings.Replace(url, "/", "\\/", -1)
	return fmt.Sprintf("^%s$", url)
}
