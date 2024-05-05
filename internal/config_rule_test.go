package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSerializeWithCommand(t *testing.T) {
	rule := ConfigRule{
		From: []string{"command", "a"},
	}
	from, _ := rule.FromSerialize()
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

func TestFromSerializeWithShift(t *testing.T) {
	rule := ConfigRule{
		From: []string{"shift", "a"},
	}
	from, _ := rule.FromSerialize()
	j, _ := json.Marshal(from)

	expected := `{
		"key_code": "a",
		"modifiers": { "mandatory": ["shift"] }
	}`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}

func TestFromSerializeWithShiftAndHasOptional(t *testing.T) {
	rule := ConfigRule{
		From:         []string{"shift", "a"},
		FromOptional: []string{"option", "command"},
	}
	from, _ := rule.FromSerialize()
	j, _ := json.Marshal(from)

	expected := `{
		"key_code": "a",
		"modifiers": { "mandatory": ["shift"] },
	  "optional": ["option", "command"]
	}`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}

func TestFromSerializeWithShiftAndHasAnyOptional(t *testing.T) {
	rule := ConfigRule{
		From:         []string{"shift", "a"},
		FromOptional: []string{"any"},
	}
	from, _ := rule.FromSerialize()
	j, _ := json.Marshal(from)

	expected := `{
		"key_code": "a",
		"modifiers": { "mandatory": ["shift"] },
	  "optional": ["any"]
	}`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}

func TestFromSerializeOnlyKeyCode(t *testing.T) {
	rule := ConfigRule{
		From: []string{"a"},
	}
	from, _ := rule.FromSerialize()
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

func TestFromSerializeWithMultipleKeyCodesError(t *testing.T) {
	rule := ConfigRule{
		From: []string{"a", "b"},
	}
	from, err := rule.FromSerialize()
	assert.Nil(t, from)
	assert.EqualError(t, err, "multiple key_code values are not allowed")
}
