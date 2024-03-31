package internal

type ConfigRuleFrom struct {
	value []string
}

func (crf *ConfigRuleFrom) serialize() map[string]interface{} {
	from := make(map[string]interface{})
	for _, value := range crf.value {
		if value == "command" {
			from["modifiers"] = map[string][]string{
				"mandatory": []string{"command"},
			}
		} else {
			from["key_code"] = value
		}
	}
	return from
}

type ConfigRuleTo struct {
	value []interface{}
}

func (crt *ConfigRuleTo) serialize() []struct {
	KeyCode string `json:"key_code"`
} {
	var to []struct {
		KeyCode string `json:"key_code"`
	}
	for _, value := range crt.value {
		switch v := value.(type) {
		case string:
			keyCode := v
			if v == "none" {
				keyCode = "vk_none"
			}
			to = append(to, struct {
				KeyCode string `json:"key_code"`
			}{
				KeyCode: keyCode,
			})
		case []interface{}:
			for _, item := range v {
				if key, ok := item.(string); ok {
					keyCode := key
					if key == "none" {
						keyCode = "vk_none"
					}
					to = append(to, struct {
						KeyCode string `json:"key_code"`
					}{
						KeyCode: keyCode,
					})
				}
			}
		}
	}
	return to
}

type Rule struct {
	Description string        `json:"description" yaml:"description"`
	From        []string      `yaml:"from"`
	To          []interface{} `yaml:"to"` // TODO: interface{}を具体的な型に変更する
}

type JSONRuleKeyCode struct {
	KeyCode string `json:"key_code"`
}

type JSONRuleFrom struct {
	KeyCode   string   `json:"key_code"`
	Modifiers []string `json:"modifiers"`
}

type JSONRule struct {
	Description string `json:"description"`
	From        map[string]interface{}
	To          []struct {
		KeyCode string `json:"key_code"`
	} `json:"to"`
	Type string `json:"type"`
}
