package slugger_test

import (
	"testing"

	"github.com/avocatl/slugger"
)

func TestSluggify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		config   *slugger.Config
		expected string
	}{
		{
			name:     "Basic English Slug",
			input:    "Hello World!",
			config:   slugger.NewConfig(),
			expected: "hello-world",
		},
		{
			name:     "German Characters",
			input:    "Fähigkeit & Übermaß",
			config:   slugger.NewConfig(slugger.WithLanguage(slugger.German)),
			expected: "faehigkeit-und-uebermass",
		},
		{
			name:     "Spanish Characters",
			input:    "Niño & Acción",
			config:   slugger.NewConfig(slugger.WithLanguage(slugger.Spanish)),
			expected: "nino-y-accion",
		},
		{
			name:     "Max Length",
			input:    "This is a very long title that should be truncated",
			config:   slugger.NewConfig(slugger.WithMaxLength(20)),
			expected: "this-is-a-very-long",
		},
		{
			name:     "No Lowercase",
			input:    "Mixed CASE Title",
			config:   slugger.NewConfig(slugger.WithoutLowercase()),
			expected: "Mixed-CASE-Title",
		},
		{
			name:     "Special Characters",
			input:    "Café @ Home #1!",
			config:   slugger.NewConfig(),
			expected: "cafe-at-home-1",
		},
		{
			name:     "Custom Separator",
			input:    "Custom Separator Test",
			config:   slugger.NewConfig(slugger.WithSeparator("_")),
			expected: "custom_separator_test",
		},
		{
			name:     "Trim Hyphens",
			input:    "---Trim This---",
			config:   slugger.NewConfig(),
			expected: "trim-this",
		},
		{
			name:     "Multiple Spaces",
			input:    "Multiple    Spaces Here",
			config:   slugger.NewConfig(),
			expected: "multiple-spaces-here",
		},
		{
			name:     "Empty String",
			input:    "",
			config:   slugger.NewConfig(),
			expected: "",
		},
		{
			name:     "Only Special Characters",
			input:    "@#$%^&*()!",
			config:   slugger.NewConfig(),
			expected: "atand",
		},
		{
			name:     "Normalize Unicode",
			input:    "Café Noël",
			config:   slugger.NewConfig(),
			expected: "cafe-noel",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sluggerService := slugger.New(tt.config)
			result := sluggerService.Slugify(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
