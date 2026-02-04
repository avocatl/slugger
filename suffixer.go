package slugger

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// Suffixer defines an interface for generating suffixes for slugs.
type Suffixer interface {
	GenerateSuffix(string) string
}

type hashSuffixer struct{
	length int
}

// GenerateSuffix generates a hash-based suffix of the specified length.
func (h *hashSuffixer) GenerateSuffix(input string) string {
	var hash uint32 = 5381 // Standard seed for DJB2 hash
	for _, char := range input {
		// hash * 33 + char
		hash = ((hash << 5) + hash) + uint32(char)
	}

	// Use a builder for memory efficiency
	var sb strings.Builder
	tempHash := int(math.Abs(float64(hash)))

	for i := 0; i < h.length; i++ {
		sb.WriteByte('a' + byte(tempHash%26))
		tempHash /= 26
	}
	return sb.String()
}

type timestampSuffixer struct{
	location *time.Location
}

// GenerateSuffix generates a timestamp-based suffix using the current time in the specified timezone.
func (ts *timestampSuffixer) GenerateSuffix(input string) string {
	// Use the pre-loaded location
	now := time.Now().In(ts.location)
	
	// Convert Unix timestamp to Base36 for a shorter, alphanumeric suffix
	return strconv.FormatInt(now.Unix(), 36)
}

// CounterProviderFunction defines a function type that provides the next counter value.
type CounterProviderFunction func() int

type numberedSuffixer struct {
	counter int
	counterProvider CounterProviderFunction
}

// GenerateSuffix generates a numbered suffix by incrementing a counter.
func (ns *numberedSuffixer) GenerateSuffix(input string) string {
	if ns.counterProvider != nil {
		ns.counter = ns.counterProvider()
	} else {
		ns.counter++
	}

	return strconv.Itoa(ns.counter)
}

type noSuffixer struct{}

// GenerateSuffix generates no suffix.
func (ns *noSuffixer) GenerateSuffix(input string) string {
	return ""
}

