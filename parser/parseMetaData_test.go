package parser

import (
	"errors"
	"strings"
	"testing"
)

func normalizeMarkdown(input string) string {
	// Remove leading and trailing whitespace
	trimmed := strings.TrimSpace(input)

	// Replace multiple consecutive newlines with a single newline
	normalized := strings.Replace(trimmed, "\n\n", "\n", -1)

	return normalized
}

func TestParseMetadata(t *testing.T) {
	tests := []struct {
		input    string
		expected MetaData
		md       string
		err      error
	}{
		{
			input: `----
Title: Test Title
Description: Test Description
PublishedDate: 2022-01-01
----
This is the content.`,
			expected: MetaData{
				Title:         "Test Title",
				Description:   "Test Description",
				PublishedDate: "2022-01-01",
			},
			md:  "This is the content.",
			err: nil,
		},
		{
			// Missing metadata section
			input:    `This is the content.`,
			expected: MetaData{},
			md:       "",
			err:      errors.New("invalid format: metadata not found"),
		},
		{
			// Missing some metadata fields
			input: `----
Title: Test Title
Description: Test Description
----
This is the content.`,
			expected: MetaData{
				Title:       "Test Title",
				Description: "Test Description",
			},
			md:  "This is the content.",
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			meta, mdContent, err := ParseMetadata([]byte(test.input))
			if err != nil && test.err == nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
			if err == nil && test.err != nil {
				t.Errorf("Expected error: %v, but got no error", test.err)
			}
			if meta != test.expected {
				t.Errorf("Expected metadata:\n%v\n\nGot:\n%v", test.expected, meta)
			}
			if normalizeMarkdown(string(mdContent)) != normalizeMarkdown(test.md) {
				t.Errorf("Expected Markdown content:\n%s\n\nGot:\n%s", test.md, string(mdContent))
			}
		})
	}
}
