package greeting

import (
	"fmt"
	"strings"
)

// Hello declaration
func Hello(input interface{}) interface{} {
	url := strings.Trim(fmt.Sprintf("%s", input), "/")
	urlParts := strings.Split(url, "/")
	name := urlParts[len(urlParts)-1]
	return fmt.Sprintf("Hello %s", name)
}
