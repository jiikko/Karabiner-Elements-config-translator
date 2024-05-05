package internal

import (
	"errors"
)

type ConfigRule struct {
	Description  string        `yaml:"description"`
	From         []string      `yaml:"from"`
	FromOptional []string      `yaml:"from_optional"`
	To           []interface{} `yaml:"to"` // TODO: interface{}を具体的な型に変更する
}

func (r ConfigRule) Serialize() (JSONRule, error) {
	serializedFrom, err := r.FromSerialize()
	if err != nil {
		return JSONRule{}, err
	}

	return JSONRule{
		Description: r.Description,
		From:        serializedFrom,
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

func isModifierKey(value string) bool {
	modifierKeys := map[string]bool{
		"command":       true,
		"shift":         true,
		"control":       true,
		"option":        true,
		"caps_lock":     true,
		"fn":            true,
		"left_command":  true,
		"right_command": true,
		"left_shift":    true,
		"right_shift":   true,
		"left_control":  true,
		"right_control": true,
		"left_option":   true,
		"right_option":  true,
	}

	return modifierKeys[value]
}

func (r ConfigRule) FromSerialize() (map[string]interface{}, error) {
	from := make(map[string]interface{})

	hasModifierKey := false
	for _, value := range r.From {
		if isModifierKey(value) {
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

	for _, value := range r.From {
		if isModifierKey(value) {
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

	if r.FromOptional != nil {
		from["optional"] = []string{}
	}

	for _, value := range r.FromOptional {
		if isModifierKey(value) || value == "any" {
			from["optional"] = append(from["optional"].([]string), value)
		}
	}

	return from, nil
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
