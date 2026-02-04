package slugger_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/avocatl/slugger"
)

func TestCofigOptions(t *testing.T) {
	cases := []struct {
		name string
		opts []slugger.ConfigOption
		expected *slugger.Config
	}{
		{
			name: "Default configuration",
			opts: []slugger.ConfigOption{},
			expected: &slugger.Config{
				Lowercase: true,
				Separator: "-",
				MaxLength: 240,
			},
		},
		{
			name: "Custom separator and max length",
			opts: []slugger.ConfigOption{
				slugger.WithSeparator("_"),
				slugger.WithMaxLength(10),
			},
			expected: &slugger.Config{
				Lowercase: true,
				Separator: "_",
				MaxLength: 10,
			},
		},
		{
			name: "Disable lowercase",
			opts: []slugger.ConfigOption{
				slugger.WithoutLowercase(),
			},
			expected: &slugger.Config{
				Lowercase: false,
				Separator: "-",
				MaxLength: 240,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := slugger.NewConfig(tc.opts...)

			log.Println(cfg == tc.expected)

			if !reflect.DeepEqual(cfg, tc.expected) {
				t.Errorf("Expected config %+v, got %+v", tc.expected, cfg)
			}
		})
	}
}