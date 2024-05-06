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

func ConvertToKeyCode(value string) (string, error) {
	keyCodeMap := map[string]string{
		" ":  "spacebar",
		":":  "semicolon",
		";":  "semicolon",
		"'":  "quote",
		"\"": "quote",
		"\\": "backslash",
		"|":  "backslash",
		",":  "comma",
		"<":  "comma",
		".":  "period",
		">":  "period",
		"/":  "slash",
		"?":  "slash",
		"=":  "equal_sign",
		"+":  "equal_sign",
		"-":  "hyphen",
		"_":  "hyphen",
		"*":  "asterisk",
	}

	customKeyCodeMap := map[string]string{
		"none": "vk_none",
	}
	returnValue, exists := keyCodeMap[value]
	if exists {
		return returnValue, nil
	} else {
		returnValue, exists = customKeyCodeMap[value]
		if exists {
			return returnValue, nil
		} else {
			return value, nil
		}
	}
}
