package domain

import (
	"testing"
)

func TestGenerateID(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "GitHub URL with path",
			url:      "https://github.com/coolya/agora/blob/main/docs/adr/0001-use-go.md",
			expected: "G1es6c5sPsH",
		},
		{
			name:     "Different GitHub URL",
			url:      "https://github.com/coolya/agora/blob/main/docs/adr/0002-use-cobra.md",
			expected: "KBviOqpmaKR",
		},
		{
			name:     "Empty string",
			url:      "",
			expected: "HVbkRR8oQIH",
		},
		{
			name:     "URL with special characters",
			url:      "https://example.com/path?query=value&param=123#fragment",
			expected: "Aoyymj7Hsw7",
		},
		{
			name:     "URL with unicode characters",
			url:      "https://example.com/文档/ADR-001.md",
			expected: "HFkb30EpPke",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GenerateID(tc.url)
			if result != tc.expected {
				t.Errorf("GenerateID(%q) = %q, expected %q", tc.url, result, tc.expected)
			}
		})
	}

	// Test determinism explicitly
	t.Run("Determinism check", func(t *testing.T) {
		url := "https://github.com/coolya/agora/docs/adr/test.md"
		id1 := GenerateID(url)
		id2 := GenerateID(url)
		if id1 != id2 {
			t.Errorf("GenerateID is not deterministic: first call returned %q, second call returned %q", id1, id2)
		}
	})

	// Test that different URLs produce different IDs
	t.Run("Uniqueness check", func(t *testing.T) {
		url1 := "https://github.com/coolya/agora/docs/adr/0001.md"
		url2 := "https://github.com/coolya/agora/docs/adr/0002.md"
		id1 := GenerateID(url1)
		id2 := GenerateID(url2)
		if id1 == id2 {
			t.Errorf("Different URLs produced the same ID: %q and %q both generated %q", url1, url2, id1)
		}
	})
}

func TestEncodeBase62(t *testing.T) {
	testCases := []struct {
		name     string
		input    uint64
		expected string
	}{
		{
			name:     "Zero",
			input:    0,
			expected: "0",
		},
		{
			name:     "Small number - 1",
			input:    1,
			expected: "1",
		},
		{
			name:     "Small number - 10",
			input:    10,
			expected: "A",
		},
		{
			name:     "Small number - 35",
			input:    35,
			expected: "Z",
		},
		{
			name:     "Small number - 36",
			input:    36,
			expected: "a",
		},
		{
			name:     "Small number - 61",
			input:    61,
			expected: "z",
		},
		{
			name:     "62 (base boundary)",
			input:    62,
			expected: "10",
		},
		{
			name:     "Large number - 1000",
			input:    1000,
			expected: "G8",
		},
		{
			name:     "Large number - 1000000",
			input:    1000000,
			expected: "4C92",
		},
		{
			name:     "Very large number - max uint32",
			input:    4294967295,
			expected: "4gfFC3",
		},
		{
			name:     "Very large number - max uint64",
			input:    18446744073709551615,
			expected: "LygHa16AHYF",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := encodeBase62(tc.input)
			if result != tc.expected {
				t.Errorf("encodeBase62(%d) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}

	// Test that encoding is consistent
	t.Run("Encoding consistency", func(t *testing.T) {
		num := uint64(123456789)
		enc1 := encodeBase62(num)
		enc2 := encodeBase62(num)
		if enc1 != enc2 {
			t.Errorf("encodeBase62 is not consistent: first call returned %q, second call returned %q", enc1, enc2)
		}
	})
}

func TestTableName(t *testing.T) {
	adr := ADR{}
	tableName := adr.TableName()
	expected := "adrs"
	if tableName != expected {
		t.Errorf("TableName() = %q, expected %q", tableName, expected)
	}
}
