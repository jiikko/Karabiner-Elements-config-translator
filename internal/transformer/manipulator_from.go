package transformer

import (
	"errors"

	"github.com/jiikko/Karabiner-Elements-config-yaml/internal/util"
)

type ConfigRuleManipulatorFrom struct {
	From         []string
	FromOptional []string
}

func (s ConfigRuleManipulatorFrom) Transform() (map[string]interface{}, error) {
	from := make(map[string]interface{})

	hasModifierKey := false
	for _, value := range s.From {
		if util.IsModifierKey(value) {
			hasModifierKey = true
			break
		}
	}

	if hasModifierKey {
		from["modifiers"] = map[string]interface{}{
			"mandatory": []interface{}{},
		}
	} else {
		delete(from, "modifiers")
	}

	for _, value := range s.From {
		if util.IsModifierKey(value) {
			from["modifiers"].(map[string]interface{})["mandatory"] = append(
				from["modifiers"].(map[string]interface{})["mandatory"].([]interface{}),
				value,
			)
		} else {
			if _, exists := from["key_code"]; exists {
				return nil, errors.New("multiple key_code values are not allowed")
			}

			from["key_code"] = value
		}
	}

	if s.FromOptional != nil {
		from["optional"] = []string{}
	}

	for _, value := range s.FromOptional {
		if util.IsModifierKey(value) || value == "any" {
			from["optional"] = append(from["optional"].([]string), value)
		}
	}

	return from, nil
}
