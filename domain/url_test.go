package domain

import "testing"

func TestShouldParseAndValidateURLFormat(t *testing.T) {
	tests := map[string]struct {
		url    string
		sucess bool
	}{
		"valid-url-https": {
			"https://test.com.br",
			true,
		},
		"valid-url-http": {
			"http://test.com.br",
			true,
		},
		"invalid-url-scheme": {
			"test.com.br",
			false,
		},
		"invalid-url-host": {
			"https://",
			false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewURL(tc.url)
			if err != nil && tc.sucess {
				t.Errorf("Unexpected error when parsing url %v", err)
			}
		})
	}
}

func TestShouldConvertFromDecimalToBase62(t *testing.T) {
	tests := map[string]struct {
		input    int64
		expected string
		desc     string
	}{
		"valid_alphanumeric": {
			input:    25670,
			expected: "6g2",
			desc:     "valid decimal number",
		},
		"valid_alphanumeric_2": {
			input:    4567894,
			expected: "JAJi",
			desc:     "valid decimal number",
		},
		"valid_zero_value": {
			input:    0,
			expected: "0",
			desc:     "valid decimal number",
		},
		"invalid_alphanumeric": {
			input:    -1,
			expected: "",
			desc:     "invalid decimal number",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := convertToBase62(tc.input)
			if err != nil && tc.expected != "" {
				t.Errorf("%s: isValidBase62(%q) = %v; want %v (%s)",
					name, tc.input, got, tc.expected, tc.desc)
			}
		})
	}
}

func TestShouldConvertFromBase62ToDecimal(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int64
		desc     string
	}{
		"valid_base62_string": {
			input:    "6g2",
			expected: 25670,
			desc:     "valid base62 string",
		},
		"valid_base62_string_2": {
			input:    "JAJi",
			expected: 4567894,
			desc:     "valid base62 string",
		},
		"invalid_base62": {
			input:    "JA^^Ji",
			expected: -1,
			desc:     "invalid base62 string",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := convertToDecimal(tc.input)
			if err != nil && tc.expected != -1 {
				t.Errorf("%s: isValidBase62(%q) = %v; want %v (%s)",
					name, tc.input, got, tc.expected, tc.desc)
			}
		})
	}
}

func TestIsValidBase62(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected bool
		desc     string
	}{
		"valid_alphanumeric": {
			input:    "6g2",
			expected: true,
			desc:     "valid alphanumeric string",
		},
		"valid_mixed_case": {
			input:    "JAJi",
			expected: true,
			desc:     "valid mixed case string",
		},
		"invalid_special_chars": {
			input:    "JA^^Ji",
			expected: false,
			desc:     "string with special characters",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isValidBase62(tc.input)
			if got != tc.expected {
				t.Errorf("%s: isValidBase62(%q) = %v; want %v (%s)",
					name, tc.input, got, tc.expected, tc.desc)
			}
		})
	}
}
