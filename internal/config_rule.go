package internal

type Rule struct {
	Description string        `yaml:"description"`
	From        []string      `yaml:"from"`
	To          []interface{} `yaml:"to"` // TODO: interface{}を具体的な型に変更する
}

func (r Rule) FromSerialize() map[string]interface{} {
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

type KeyCodeStruct struct {
	KeyCode string `json:"key_code"`
}

func (r Rule) ToSerialize() []KeyCodeStruct {
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
