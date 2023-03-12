package route

import "strings"

func PathApi(path string) string {
	if strings.HasPrefix(path, "/") {
		return "/api" + path
	}
	return "/api/" + path
}
