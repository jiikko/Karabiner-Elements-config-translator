package internal

type ConfigRule struct {
	Description string        `yaml:"description"`
	From        []string      `yaml:"from"`
	To          []interface{} `yaml:"to"` // TODO: interface{}を具体的な型に変更する
}

func (r ConfigRule) Serialize() (JSONRule, error) {
	return JSONRule{
		Description: r.Description,
		From:        r.FromSerialize(),
		To:          r.ToSerialize(),
		Type:        "basic",
	}, nil
}

type KeyCodeStruct struct {
	KeyCode string `json:"key_code"`
}

type ConfigRuleFrom struct {
	KeyCode string `json:"key_code,omitempty"`
}

func (r ConfigRule) FromSerialize() map[string]interface{} {
	from := make(map[string]interface{})
	for _, value := range r.From {
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

func (r ConfigRule) ToSerialize() []KeyCodeStruct {
	var to []KeyCodeStruct
	for _, value := range r.To {
		switch v := value.(type) {
		case string:
			keyCode := v
			if v == "none" {
				keyCode = "vk_none"
			}
			to = append(to, KeyCodeStruct{
				KeyCode: keyCode,
			})
		case []interface{}:
			for _, item := range v {
				if key, ok := item.(string); ok {
					keyCode := key
					if key == "none" {
						keyCode = "vk_none"
					}
					to = append(to, KeyCodeStruct{
						KeyCode: keyCode,
					})
				}
			}
		}
	}
	return to
}
