package html

import (
	"fmt"
	"github.com/state-alchemists/ayanami/service"
)

// pre declaration
func pre(input string) string {
	return fmt.Sprintf("<pre>%s</pre>", input)
}

// Wrappedpre wrapper for pre
func Wrappedpre(inputs service.Dictionary) (service.Dictionary, error) {
	Result := make(service.Dictionary)
	input := fmt.Sprintf("%s", inputs.Get("input"))
	output := pre(input)
	Result.Set("output", output)
	return Result, nil
}
