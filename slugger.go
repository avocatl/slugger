package slugger

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Config holds configuration options for the Slugger service.
type Slugger interface {
	Slugify(input string) string
}

type srv struct {
	c *Config
	alphaRegex *regexp.Regexp
	hyphenRegex *regexp.Regexp
	replacer LanguageReplacer
	suffixer Suffixer
	transformer transform.Transformer
}

// Slugify converts the input string into a slug format.
func (b *srv) Slugify(input string) string {
    slug := input

    // 1. Lowercase early so following steps are simpler
    if b.c.Lowercase {
        slug = strings.ToLower(slug)
    }

    // 2. Language-specific replacements (e.g., @ -> at, ß -> ss)
    slug = b.replacer.Replace(slug)

    // 3. Normalize to NFD & Strip marks (handles ç -> c, ñ -> n)
    slug = b.normalizeToNFD(slug)

    // 4. Clean up characters
    slug = b.alphaRegex.ReplaceAllString(slug, "")
    slug = b.hyphenRegex.ReplaceAllString(slug, b.c.Separator)
    slug = strings.Trim(slug, b.c.Separator)

    // 5. Truncate safely (check bounds)
    if b.c.MaxLength > 0 && len(slug) > b.c.MaxLength {
        slug = slug[:b.c.MaxLength]
        slug = strings.Trim(slug, b.c.Separator) // Trim again in case we cut into a separator
    }

    // 6. Suffixing
    if b.c.SuffixStrategy != None {
        suffix := b.suffixer.GenerateSuffix(slug)
        slug = slug + b.c.Separator + suffix
    }

    return slug
}

func (b *srv) normalizeToNFD(input string) string {	
	output, _, err := transform.String(b.transformer, input)
	if err != nil {
		return input
	}

	return output
}

// New creates a new instance of Slugger with the provided configuration.
func New(c *Config) Slugger {
    // 1. Build Regex based on case sensitivity
    charRange := "a-z0-9"
    if !c.Lowercase {
        charRange = "a-zA-Z0-9"
    }
    // We include the separator in the 'allowed' list so we don't strip it
    alphaRe := regexp.MustCompile(fmt.Sprintf(`[^%s\s%s]`, charRange, regexp.QuoteMeta(c.Separator)))
    hyphenRe := regexp.MustCompile(fmt.Sprintf(`[\s%s]+`, regexp.QuoteMeta(c.Separator)))

    // 2. Pre-configure the Transformer
    transformer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

    // 3. Strategy Switch
    var suffixer Suffixer
    switch c.SuffixStrategy {
    case HashBased:
        suffixer = &hashSuffixer{length: c.SuffixLength}
    case TimestampBased:
        // Load location ONCE here to avoid I/O in the Slugify loop
        loc, err := time.LoadLocation(c.TimestampTimezone)
        if err != nil {
            loc = time.UTC
        }
        suffixer = &timestampSuffixer{location: loc}
    case Numbered:
        suffixer = &numberedSuffixer{counterProvider: c.CounterProvider}
    default:
        suffixer = &noSuffixer{}
    }

    return &srv{
        c:           c,
        alphaRegex:  alphaRe,
        hyphenRegex: hyphenRe,
        replacer:    NewReplacer(c.Language),
        suffixer:    suffixer,
        transformer: transformer, // No need for pointer if it's an interface or specific type
    }
}