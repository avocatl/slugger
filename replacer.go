package slugger

import "strings"

// LanguageReplacer defines an interface for replacing special characters based on language.
type LanguageReplacer interface {
	Replace(string) string
}

type germanReplacer struct {
	rep *strings.Replacer
}

// Replace replaces special German characters in the input string.
func (r *germanReplacer) Replace(input string) string {
	return r.rep.Replace(input)
}

type englishReplacer struct {
	rep *strings.Replacer
}

// Replace replaces special English characters in the input string.
func (r *englishReplacer) Replace(input string) string {
	return r.rep.Replace(input)
}

type spanishReplacer struct {
	rep *strings.Replacer
}

// Replace replaces special Spanish characters in the input string.
func (r *spanishReplacer) Replace(input string) string {
	return r.rep.Replace(input)
}

// NewReplacer creates a LanguageReplacer based on the specified language.
func NewReplacer(language Language) LanguageReplacer {
	switch language {
	case German:
		var oldnew []string

		for old, new := range GermanCharacterMapping {
			oldnew = append(oldnew, old, new)
		}

		rep := strings.NewReplacer(oldnew...)
		
		return &germanReplacer{rep: rep}

	case English:
		var oldnew []string

		for old, new := range EnglishCharacterMapping {
			oldnew = append(oldnew, old, new)
		}

		rep := strings.NewReplacer(oldnew...)

		return &englishReplacer{rep: rep}

	case Spanish:
		var oldnew []string

		for old, new := range SpanishCharacterMapping {
			oldnew = append(oldnew, old, new)
		}

		rep := strings.NewReplacer(oldnew...)

		return &spanishReplacer{rep: rep}

	default:
		// For unsupported languages, return a no-op replacer.
		return &noopReplacer{}
	}
}

type noopReplacer struct{}

// Replace returns the input string unchanged.
func (r *noopReplacer) Replace(input string) string {
	return input
}
