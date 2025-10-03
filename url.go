package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func Parse(rawURL string) (*URL, error) {
	scheme, rest, ok := strings.Cut(rawURL, ":")

	if !ok || scheme == "" {
		return nil, errors.New("missing scheme")
	}

	if !strings.HasPrefix(rest, "//") {
		return &URL{Scheme: scheme}, nil
	}

	host, path, _ := strings.Cut(rest[2:], "/")
	return &URL{Scheme: scheme, Host: host, Path: path}, nil
}

func (u *URL) String() string {
	if u == nil {
		return ""
	}

	var sb strings.Builder
	sb.Grow(len(u.Scheme) + 3 + len(u.Host) + 1 + len(u.Path))

	if s := u.Scheme; s != "" {
		sb.WriteString(s)
		sb.WriteString("://")
	}
	if s := u.Host; s != "" {
		sb.WriteString(s)
	}
	if s := u.Path; s != "" {
		sb.WriteString("/")
		sb.WriteString(s)
	}

	return sb.String()
}
