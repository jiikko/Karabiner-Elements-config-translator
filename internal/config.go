package internal

import "gopkg.in/yaml.v2"

type Rule struct {
	Description string     `yaml:"description"`
	From        []string   `yaml:"from"`
	To          [][]string `yaml:"to"`
}

type Config struct {
	Title       []string `yaml:"title"`
	Maintainers []string `yaml:"maintainers"`
	Rules       []Rule   `yaml:"rules"`
	Format      string   `yaml:"format"`
}

func (c Config) ToJSON(content string) string {
	if true { // FIXME: configの中身を判断する
		c.Format = "simple"
	} else {
		c.Format = "complex"
	}

	return content
}

func NewConfig(content string) Config {
	var config Config
	err := yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		return Config{}
	}
	return config
}
