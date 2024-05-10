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
