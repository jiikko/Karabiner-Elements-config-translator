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

type JSONRule struct {
	Description string                 `json:"description"`
	From        map[string]interface{} `json:"from"`
	To          []KeyCodeStruct        `json:"to"`
	Type        string                 `json:"type"`
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
	jsonConfig := map[string]interface{}{
		"title":       c.Title,
		"maintainers": c.Maintainers,
		"rules": func() []JSONRule {
			var outputRules []JSONRule
			for _, rule := range c.Rules {
				jsonRule, err := rule.Serialize()
				if err != nil {
					return nil
				}

				outputRules = append(outputRules, jsonRule)
			}
			return outputRules
		}(),
	}

	jsonData, err := json.Marshal(jsonConfig)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
