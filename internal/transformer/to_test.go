package transformer

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromTransoforToOnlyKeyCode(t *testing.T) {
	mt := ManipulatorTo{
		To: []interface{}{"a", "b"},
	}
	to, _ := mt.Transform()
	j, _ := json.Marshal(to)
	expected := `[
		{ "key_code": "a" },
		{ "key_code": "b" }
	]`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}

func TestFromTransoforTo(t *testing.T) {
	mt := ManipulatorTo{
		To: []interface{}{"a", "b", "none"},
	}
	to, _ := mt.Transform()
	j, _ := json.Marshal(to)
	expected := `[
		{ "key_code": "a" },
		{ "key_code": "b" },
		{ "key_code": "vk_none" }
	]`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}

func TestFromTransoforTo2(t *testing.T) {
	mt := ManipulatorTo{
		To: []interface{}{":", " ", "none"},
	}
	to, _ := mt.Transform()
	j, _ := json.Marshal(to)
	expected := `[
		{ "key_code": "semicolon" },
		{ "key_code": "spacebar" },
		{ "key_code": "vk_none" }
	]`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}

func TestFromTransoforToWithModifier(t *testing.T) {
	mt := ManipulatorTo{
		To: []interface{}{
			"a",
			[]interface{}{":", "shift"},
		},
	}
	to, _ := mt.Transform()
	j, _ := json.Marshal(to)
	// p.P(string(j))
	expected := `[
		{ "key_code": "a" },
		{ "key_code": "semicolon", "modifiers": { "mandatory": ["shift"] } }
	]`
	assert.JSONEq(
		t,
		expected,
		string(j),
	)
}
