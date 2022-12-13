package rules

import "github.com/nikunjy/rules/parser"

func Evaluate(rule string, items map[string]interface{}) (bool, error) {
	ev, err := parser.NewEvaluator(rule)
	if err != nil {
		return false, err
	}
	return ev.Process(items)
}

func EvaluateWithErrInfo(rule string, items map[string]interface{}) (bool, error) {
	ev, err := parser.NewEvaluator(rule)
	if err != nil {
		return false, err
	}
	result, err := ev.Process(items)
	if err != nil {
		return result, err
	} else {
		return result, ev.LastDebugErr()
	}
}
