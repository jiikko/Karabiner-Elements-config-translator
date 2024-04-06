package internal

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title       string
	Maintainers []string
	Rules       []Rule
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
	// p.P(config)
	return &config, nil
}

func (c Config) ToJSON(content string) (string, error) {
	jsonConfig := map[string]interface{}{
		"title":       c.Title,
		"maintainers": c.Maintainers,
		"rules": func() []JSONRule {
			var rules []JSONRule
			for _, rule := range c.Rules {
				jsonRule := JSONRule{
					Description: rule.Description,
					From:        rule.FromSerialize(),
					To:          rule.ToSerialize(),
					Type:        "basic",
				}
				rules = append(rules, jsonRule)
			}
			return rules
		}(),
	}

	jsonData, err := json.Marshal(jsonConfig)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
