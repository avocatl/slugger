package slugger_test

import (
	"testing"

	"github.com/avocatl/slugger"
)

func TestReplacer(t *testing.T) {
	cases := []struct {
		name    string
		language slugger.Language
		input    string
		expected string
	}{
		{
			name:     "German Replacer",
			language: slugger.German,
			input:    "Fähigkeit Straße über",
			expected: "Faehigkeit Strasse ueber",
		},
		{
			name:     "English Replacer",
			language: slugger.English,
			input:    "Café & Drinks",
			expected: "Café and Drinks",
		},
		{
			name:     "Spanish Replacer",
			language: slugger.Spanish,
			input:    "niño corazón jalapeño",
			expected: "nino corazon jalapeno",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			replacer := slugger.NewReplacer(tc.language)
			result := replacer.Replace(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}
