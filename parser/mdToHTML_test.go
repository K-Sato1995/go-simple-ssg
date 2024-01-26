package parser

import (
	"strings"
	"testing"
)

func normalizeWhitespace(input string) string {
	return strings.Join(strings.Fields(input), " ")
}

func TestMdToHTML(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "# Heading 1\n\nThis is a paragraph.",
			expected: `<h1 id="heading-1">Heading 1</h1> <p>This is a paragraph.</p>`,
		},
		{
			input:    "## Heading 2\n\n- List item 1\n- List item 2",
			expected: `<h2 id="heading-2">Heading 2</h2> <ul> <li>List item 1</li> <li>List item 2</li> </ul>`,
		},
		{
			input:    "*Italic* and **Bold** text",
			expected: `<p><em>Italic</em> and <strong>Bold</strong> text</p>`,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := MdToHTML([]byte(test.input))
			expected := normalizeWhitespace(test.expected)
			actual := normalizeWhitespace(string(result))
			if actual != expected {
				t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, actual)
			}
		})
	}
}
