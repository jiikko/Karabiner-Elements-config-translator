package transformer

import "github.com/jiikko/Karabiner-Elements-config-yaml/internal/util"

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
			for _, item := range v {
				if key, ok := item.(string); ok {
					keyCode, err := util.ConvertToKeyCode(key)
					if err != nil {
						return nil, err
					}
					transforms = append(transforms, map[string]interface{}{
						"key_code": keyCode,
					})
				}
			}
		}
	}

	return transforms, nil
}
