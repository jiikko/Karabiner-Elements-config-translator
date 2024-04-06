package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSerialize(t *testing.T) {
	rule := ConfigRule{
		From: []string{"command", "a"},
	}
	from := rule.FromSerialize()
	j, _ := json.Marshal(from)

	expected := `{
		"key_code": "a",
		"modifiers": { "mandatory": ["command"] }
	}`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}

func TestFromSerialize2(t *testing.T) {
	rule := ConfigRule{
		From: []string{"a"},
	}
	from := rule.FromSerialize()
	j, _ := json.Marshal(from)

	expected := `{
		"key_code": "a"
	}`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}

func TestFromSerialize3(t *testing.T) {
	rule := ConfigRule{
		From: []string{"a", "b"},
	}
	from := rule.FromSerialize()
	j, _ := json.Marshal(from)

	// TODO: エラーにする
	expected := `{
		"key_code": "b"
	}`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}
