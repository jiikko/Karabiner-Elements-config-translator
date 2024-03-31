package internal

import (
	"fmt"

	"github.com/Code-Hex/dd/p"
	"gopkg.in/yaml.v2"
)

type Rule struct {
	Description string      `yaml:"description"`
	From        []string    `yaml:"from"`
	To          interface{} `yaml:"to"` // TODO: interface{}を具体的な型に変更する
}

type Config struct {
	Title       []string `yaml:"title"`
	Maintainers []string `yaml:"maintainers"`
	Rules       []Rule   `yaml:"rules"`
	Format      string
}

func (c Config) ToJSON(content string) string {
	if true { // FIXME: configの中身を判断する
		c.Format = "simple"
	} else {
		c.Format = "full"
	}

	return content
}

func NewConfig(content string) (*Config, error) {
	var config Config
	err := yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		return nil, err
	}
	p.P(config)
	fmt.Println("--")
	return &config, nil
}
