package url

import "testing"

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

func TestURLString(t *testing.T) {
	u := &URL{Scheme: "https", Host: "github.com", Path: "akthrmsx"}
	got := u.String()
	want := "https://github.com/akthrmsx"

	if got != want {
		t.Errorf("String(%#v)\ngot  = %q\nwant = %q", u, got, want)
	}
}
