package goose

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