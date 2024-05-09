package internal

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title       string
	Maintainers []string
	Rules       []ConfigRule
}

type JSONRuleManipulator struct {
	From map[string]interface{}   `json:"from"`
	To   []map[string]interface{} `json:"to"`
	Type string                   `json:"type"`
}

type JSONRule struct {
	Description  string                `json:"description"`
	Manipulators []JSONRuleManipulator `json:"manipulators"`
}

func NewConfig(content string) (*Config, error) {
	var config Config
	err := yaml.Unmarshal([]byte(content), &config)

	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c Config) ToJSON(content string) (string, error) {
	var err error
	jsonConfig := map[string]interface{}{
		"title":       c.Title,
		"maintainers": c.Maintainers,
		"rules": func() []JSONRule {
			var outputRules []JSONRule
			for _, rule := range c.Rules {
				jsonRule, e := rule.Transform()

				if e != nil {
					err = e
					return nil
				}

				outputRules = append(outputRules, jsonRule)
			}
			return outputRules
		}(),
	}

	if err != nil {
		return "", err
	}

	jsonData, e := json.Marshal(jsonConfig)
	if e != nil {
		return "", e
	}
	return string(jsonData), nil
}
