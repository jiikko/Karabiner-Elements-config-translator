package transformer

import (
	"github.com/jiikko/Karabiner-Elements-config-yaml/internal/util"
)

type ManipulatorTo struct {
	To []interface{}
}

func (m ManipulatorTo) Transform() ([]map[string]interface{}, error) {
	var transforms []map[string]interface{}
	for _, value := range m.To {
		switch v := value.(type) {
		case string:
			keyCode, err := util.ConvertToKeyCode(v)
			if err != nil {
				return nil, err
			}
			transforms = append(transforms, map[string]interface{}{
				"key_code": keyCode,
			})
		case []interface{}:
			item_in_to := make(map[string]interface{})

			hasModifierKey := false
			for _, vv := range v {
				if s, ok := vv.(string); ok {
					if util.IsModifierKey(s) {
						hasModifierKey = true
						break
					}
				}
			}

			if hasModifierKey {
				item_in_to["modifiers"] = map[string]interface{}{
					"mandatory": []string{},
				}
			} else {
				delete(item_in_to, "modifiers")
			}

			for _, vv := range v {
				if s, ok := vv.(string); ok {
					keyCode, err := util.ConvertToKeyCode(s)
					if err != nil {
						return nil, err
					}

					if util.IsModifierKey(s) {
						// item_in_to["modifiers"] = append(item_in_to["modifiers"].([]interface{}), s)
						item_in_to["modifiers"].(map[string]interface{})["mandatory"] = append(item_in_to["modifiers"].(map[string]interface{})["mandatory"].([]string), s)
					} else {
						item_in_to["key_code"] = keyCode
					}
				}
			}
			transforms = append(transforms, item_in_to)
		}
	}

	return transforms, nil
}
