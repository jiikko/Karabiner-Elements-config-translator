package transformer

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromTransoformWithCommand(t *testing.T) {
	mf := ManipulatorFrom{
		From: []string{"command", "a"},
	}
	from, _ := mf.Transform()
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

func TestFromTransoformWithShift(t *testing.T) {
	mf := ManipulatorFrom{
		From: []string{"shift", "a"},
	}
	from, _ := mf.Transform()
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

func TestFromTransoformWithShiftAndHasOptional(t *testing.T) {
	mf := ManipulatorFrom{
		From:         []string{"shift", "a", "control"},
		FromOptional: []string{"option", "command"},
	}
	from, _ := mf.Transform()
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

func TestFromTransoformWithShiftAndHasAnyOptional(t *testing.T) {
	mf := ManipulatorFrom{
		From:         []string{"shift", "a"},
		FromOptional: []string{"any"},
	}
	from, _ := mf.Transform()
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

func TestFromTransoformOnlyKeyCode(t *testing.T) {
	mf := ManipulatorFrom{
		From: []string{"a"},
	}
	from, _ := mf.Transform()
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

func TestFromTransoformWithMultipleKeyCodesError(t *testing.T) {
	mf := ManipulatorFrom{
		From: []string{"a", "b"},
	}
	from, err := mf.Transform()
	assert.Nil(t, from)
	assert.EqualError(t, err, "multiple key_code values are not allowed")
}
