package internal

import (
	"errors"

	"github.com/jiikko/Karabiner-Elements-config-yaml/internal/util"
)

type ConfigRuleManipulator struct {
	From         []string      `yaml:"from"`
	FromOptional []string      `yaml:"from_optional"`
	To           []interface{} `yaml:"to"` // TODO: interface{}を具体的な型に変更する
}

type ConfigRule struct {
	Description  string                  `yaml:"description"`
	Manipulators []ConfigRuleManipulator `yaml:"manipulators"`
}

func (r ConfigRule) Serialize() (JSONRule, error) {
	var jsonManipulators []JSONRuleManipulator
	for _, manipulator := range r.Manipulators {
		// p.P("manipulator", manipulator)
		jsonManipulator, err := manipulator.serialize()
		if err != nil {
			return JSONRule{}, err
		}
		jsonManipulators = append(jsonManipulators, jsonManipulator)
	}

	return JSONRule{
		Description:  r.Description,
		Manipulators: jsonManipulators,
	}, nil
}

type KeyCodeStruct struct {
	KeyCode string `json:"key_code"`
}

type ConfigRuleFrom struct {
	KeyCode string `json:"key_code,omitempty"`
}

func (m ConfigRuleManipulator) serialize() (JSONRuleManipulator, error) {
	serializedFrom, err := m.FromSerialize()
	if err != nil {
		return JSONRuleManipulator{}, err
	}

	return JSONRuleManipulator{
		From: serializedFrom,
		To:   m.ToSerialize(),
		Type: "basic",
	}, nil
}

func (r ConfigRuleManipulator) FromSerialize() (map[string]interface{}, error) {
	from := make(map[string]interface{})

	hasModifierKey := false
	for _, value := range r.From {
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

	for _, value := range r.From {
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

	if r.FromOptional != nil {
		from["optional"] = []string{}
	}

	for _, value := range r.FromOptional {
		if util.IsModifierKey(value) || value == "any" {
			from["optional"] = append(from["optional"].([]string), value)
		}
	}

	return from, nil
}

func (r ConfigRuleManipulator) ToSerialize() []KeyCodeStruct {
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
