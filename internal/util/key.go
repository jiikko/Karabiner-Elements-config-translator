package util

func IsModifierKey(value string) bool {
	modifierKeys := map[string]bool{
		"command":       true,
		"shift":         true,
		"control":       true,
		"option":        true,
		"caps_lock":     true,
		"fn":            true,
		"left_command":  true,
		"right_command": true,
		"left_shift":    true,
		"right_shift":   true,
		"left_control":  true,
		"right_control": true,
		"left_option":   true,
		"right_option":  true,
	}

	return modifierKeys[value]
}
