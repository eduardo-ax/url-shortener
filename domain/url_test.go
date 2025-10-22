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

	for n, v := range tests {
		t.Run(n, func(t *testing.T) {
			_, err := NewUrl(v.url)
			if err != nil && v.sucess {
				t.Errorf("Unexpected error when parsing url %v", err)
			}
		})
	}
}
