// internal/eval/engine.go
package eval

import (
	"featurn/internal/flags"
)

func Evaluate(flag flags.Flag, context map[string]interface{}) interface{} {
	for _, rule := range flag.Rules {
		if val, ok := context[rule.Attribute]; ok {
			if rule.Operator == "eq" && val == rule.MatchValue {
				return rule.ReturnValue
			}
		}
	}
	return flag.DefaultValue
}
