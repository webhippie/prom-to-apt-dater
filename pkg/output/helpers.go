package output

import (
	"github.com/Knetic/govaluate"
)

// Helpers defines a list of helpers for Govaluate.
func Helpers() map[string]govaluate.ExpressionFunction {
	return map[string]govaluate.ExpressionFunction{
		"static": func(args ...interface{}) (interface{}, error) {
			result := ""

			for _, arg := range args {
				result = result + arg.(string)
			}

			return result, nil
		},
	}
}
