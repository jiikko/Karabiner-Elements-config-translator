package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ToJSON(t *testing.T) {
	sampleFilePath := "../testdata/sample.yml"

	expectedJSON := `{
		"maintainers": [
		  "foo"
		],
		"rules": [
		  {
				"description": "disable command + m(最小化)",
				"manipulators": [
					{
					"from": {
						"key_code": "m",
						"modifiers": { "mandatory": [ "command" ] }
					},
					"to": [ { "key_code": "vk_none" } ],
					"type": "basic"
					}
				]
		  }
		],
		"title": "my config"
	}`

	parser, err := NewParser(sampleFilePath)
	assert.NoError(t, err)

	jsonStr, err := parser.ToJSON()
	assert.NoError(t, err)

	assert.JSONEq(t, expectedJSON, jsonStr)
}

func TestParser_ToJSON_WithOptional(t *testing.T) {
	sampleFilePath := "../testdata/config_with_optional.yml"

	expectedJSON := `{
		"maintainers": [
		  "foo"
		],
		"rules": [
		  {
				"description": "disable command + m(最小化)",
				"manipulators": [
					{
					"from": {
						"key_code": "m",
						"modifiers": { "mandatory": [ "command" ] },
	            "optional": [
	              "shift",
	              "control"
	            ]
					},
					"to": [ { "key_code": "vk_none" } ],
					"type": "basic"
					}
				]
		  }
		],
		"title": "my config"
	}`

	parser, err := NewParser(sampleFilePath)
	assert.NoError(t, err)

	jsonStr, err := parser.ToJSON()
	assert.NoError(t, err)

	assert.JSONEq(t, expectedJSON, jsonStr)
}
