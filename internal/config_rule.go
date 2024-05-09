package internal

import (
	"github.com/jiikko/Karabiner-Elements-config-yaml/internal/serializer"
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

type KeyCodeStruct struct {
	KeyCode string `json:"key_code"`
}

type ConfigRuleFrom struct {
	KeyCode string `json:"key_code,omitempty"`
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

func (m ConfigRuleManipulator) serialize() (JSONRuleManipulator, error) {
	serializerFrom := serializer.ConfigRuleManipulatorFrom{
		From:         m.From,
		FromOptional: m.FromOptional,
	}

	serializedFrom, err := serializerFrom.Serialize()
	if err != nil {
		return JSONRuleManipulator{}, err
	}

	return JSONRuleManipulator{
		From: serializedFrom,
		To:   m.toSerialize(),
		Type: "basic",
	}, nil
}

func (m ConfigRuleManipulator) toSerialize() []KeyCodeStruct {
	var to []KeyCodeStruct
	for _, value := range m.To {
		switch v := value.(type) {
		case string:
			keyCode, err := util.ConvertToKeyCode(v)
			if err != nil {
				return nil
			}
			to = append(to, KeyCodeStruct{
				KeyCode: keyCode,
			})
		case []interface{}:
			for _, item := range v {
				if key, ok := item.(string); ok {
					keyCode, err := util.ConvertToKeyCode(key)
					if err != nil {
						return nil
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
