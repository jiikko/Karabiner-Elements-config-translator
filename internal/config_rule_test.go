package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestfromSerializeWithCommand(t *testing.T) {
	rule := ConfigRuleManipulator{
		From: []string{"command", "a"},
	}
	from, _ := rule.fromSerialize()
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

func TestfromSerializeWithShift(t *testing.T) {
	rule := ConfigRuleManipulator{
		From: []string{"shift", "a"},
	}
	from, _ := rule.fromSerialize()
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

func TestfromSerializeWithShiftAndHasOptional(t *testing.T) {
	rule := ConfigRuleManipulator{
		From:         []string{"shift", "a", "control"},
		FromOptional: []string{"option", "command"},
	}
	from, _ := rule.fromSerialize()
	j, _ := json.Marshal(from)

	expected := `{
		"key_code": "a",
		"modifiers": { "mandatory": ["shift", "control"] },
	  "optional": ["option", "command"]
	}`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}

func TestfromSerializeWithShiftAndHasAnyOptional(t *testing.T) {
	rule := ConfigRuleManipulator{
		From:         []string{"shift", "a"},
		FromOptional: []string{"any"},
	}
	from, _ := rule.fromSerialize()
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

func TestfromSerializeOnlyKeyCode(t *testing.T) {
	rule := ConfigRuleManipulator{
		From: []string{"a"},
	}
	from, _ := rule.fromSerialize()
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

func TestfromSerializeWithMultipleKeyCodesError(t *testing.T) {
	rule := ConfigRuleManipulator{
		From: []string{"a", "b"},
	}
	from, err := rule.fromSerialize()
	assert.Nil(t, from)
	assert.EqualError(t, err, "multiple key_code values are not allowed")
}
