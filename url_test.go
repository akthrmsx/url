package url

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		uri  string
		want *URL
	}{
		{
			name: "with_data_scheme",
			uri:  "data:text/plain;base64,SGVsbG8sIFdvcmxkIQ==",
			want: &URL{Scheme: "data"},
		},
		{
			name: "full",
			uri:  "https://github.com/akthrmsx",
			want: &URL{Scheme: "https", Host: "github.com", Path: "akthrmsx"},
		},
		{
			name: "without_path",
			uri:  "https://github.com",
			want: &URL{Scheme: "https", Host: "github.com", Path: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.uri)

			if err != nil {
				t.Fatalf("Parse(%q)\ngot = %q\nerr = <nil>", tt.uri, err)
			}

			if *got != *tt.want {
				t.Errorf("Parse(%q)\ngot  = %#v\nwant = %#v", tt.uri, got, tt.want)
			}
		})
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		name string
		uri  string
	}{
		{
			name: "without_scheme",
			uri:  "github.com",
		},
		{
			name: "empty_scheme",
			uri:  "://github.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.uri)

			if err == nil {
				t.Errorf("Parse(%q)\nerr  = <nil>\nwant = <error>", tt.uri)
			}
		})
	}
}

func TestURLString(t *testing.T) {
	tests := []struct {
		name string
		uri  *URL
		want string
	}{
		{
			name: "nil",
			uri:  nil,
			want: "",
		},
		{
			name: "empty",
			uri:  new(URL),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.uri.String()

			if got != tt.want {
				t.Errorf("String(%#v)\ngot  = %q\nwant = %q", tt.uri, got, tt.want)
			}
		})
	}
}

func BenchmarkURLString(b *testing.B) {
	u := &URL{Scheme: "https", Host: "github,com", Path: "akthrmsx"}

	for b.Loop() {
		_ = u.String()
	}
}

func BenchmarkURLStringLong(b *testing.B) {
	for _, n := range []int{10, 100, 1000} {
		u := &URL{
			Scheme: strings.Repeat("x", n),
			Host:   strings.Repeat("y", n),
			Path:   strings.Repeat("z", n),
		}

		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for b.Loop() {
				_ = u.String()
			}
		})
	}
}
