package slugger

// Language defines the language used for slugification.
type Language int

// Supported languages for slugification.
const (
	English Language = iota
	German
	Spanish
)

// Suffix defines the type of suffix to append to slugs.
type Suffix int

// Different suffix strategies.
const (
	None Suffix = iota
	Numbered
	TimestampBased
	HashBased
)

// Config holds configuration options for the slugger.
type Config struct {
	Lowercase bool
	Separator string
	MaxLength int
	Language   Language
	SuffixStrategy Suffix
	SuffixLength int
	TimestampTimezone string
	CounterProvider CounterProviderFunction
}

// ConfigOption defines a function type for configuring the Slugger.
type ConfigOption func(*Config)

// WithoutLowercase sets whether slugs should be lowercase.
func WithoutLowercase() ConfigOption {
	return func(c *Config) {
		c.Lowercase = false
	}
}

// WithSeparator sets the separator character for slugs.
func WithSeparator(separator string) ConfigOption {
	return func(c *Config) {
		c.Separator = separator
	}
}

// WithMaxLength sets the maximum length for slugs.
func WithMaxLength(maxLength int) ConfigOption {
	return func(c *Config) {
		c.MaxLength = maxLength
	}
}


// WithLanguage sets the language for slugification.
func WithLanguage(language Language) ConfigOption {
	return func(c *Config) {
		c.Language = language
	}
}

// WithSuffixStrategy sets the suffix strategy for slugs.
func WithSuffixStrategy(suffix Suffix) ConfigOption {
	return func(c *Config) {
		c.SuffixStrategy = suffix
	}
}

// WithHashSuffixLength sets the length of the hash-based suffix.
func WithHashSuffixLength(length int) ConfigOption {
	return func(c *Config) {
		c.SuffixStrategy = HashBased
		c.SuffixLength = length
	}
}

// WithTimestampTimezone sets the timezone for timestamp-based suffixes.
func WithTimestampTimezone(timezone string) ConfigOption {
	return func(c *Config) {
		c.SuffixStrategy = TimestampBased
		c.TimestampTimezone = timezone
	}
}

// WithNumberedCounterProvider sets the counter provider function for numbered suffixes.
func WithNumberedCounterProvider(counterProvider CounterProviderFunction) ConfigOption {
	return func(c *Config) {
		c.SuffixStrategy = Numbered
		c.CounterProvider = counterProvider
	}
}

func WithNoSuffix() ConfigOption {
	return func(c *Config) {
		c.SuffixStrategy = None
	}
}

// NewConfig creates a new Config with the provided options.
func NewConfig(options ...ConfigOption) *Config {
	config := &Config{
		Lowercase:     true,
		Separator:     "-",
		MaxLength:     240,
		Language:      English,
		SuffixStrategy: None,
	}

	for _, option := range options {
		option(config)
	}

	return config
}
