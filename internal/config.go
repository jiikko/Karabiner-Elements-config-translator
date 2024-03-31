package internal

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

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
type JSNORuleTo struct {
	KeyCode string `json:"key_code"`
}

type JSONRule struct {
	Description string `json:"description"`
	From        map[string]interface{}
	To          []struct {
		KeyCode string `json:"key_code"`
	} `json:"to"`
	Type string `json:"type"`
}

type Config struct {
	Title       string   `yaml:"title"`
	Maintainers []string `yaml:"maintainers"`
	Rules       []Rule   `yaml:"rules"`
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
					From: func() map[string]interface{} {
						from := make(map[string]interface{})
						for _, value := range rule.From {
							if value == "command" {
								from["modifiers"] = map[string][]string{
									"mandatory": []string{"command"},
								}
							} else {
								from["key_code"] = value
							}
						}
						return from
					}(),
					To: func() []struct {
						KeyCode string `json:"key_code"`
					} {
						var to []struct {
							KeyCode string `json:"key_code"`
						}
						for _, value := range rule.To {
							switch v := value.(type) {
							case string:
								to = append(to, struct {
									KeyCode string `json:"key_code"`
								}{
									KeyCode: v,
								})
							case []string:
								for _, key := range v {
									to = append(to, struct {
										KeyCode string `json:"key_code"`
									}{
										KeyCode: key,
									})
								}
							}
						}
						return to
					}(),
					Type: "basic",
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

func NewConfig(content string) (*Config, error) {
	var config Config
	err := yaml.Unmarshal([]byte(content), &config)

	if err != nil {
		return nil, err
	}
	// p.P(config)
	return &config, nil
}
