package internal

import "os"

type Parser struct {
	filepath string
	content  string
	config   *Config
}

func NewParser(filepath string) (*Parser, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	parser := &Parser{
		filepath: filepath,
		content:  string(content),
	}
	parser.config, err = NewConfig(parser.content)
	if err != nil {
		return nil, err
	}

	return parser, nil
}

func (p *Parser) ToJSON() (string, error) {
	json, err := p.config.ToJSON(p.content)
	if err != nil {
		return "", err
	}
	return json, nil
}
