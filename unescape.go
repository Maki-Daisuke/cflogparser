package cflogparser

import (
	"net/url"
	"strings"
)

func unhex(c byte) (byte, error) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', nil
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, nil
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, nil
	}
	return 0, url.EscapeError("")
}

func Unescape(s string) (string, error) {
	var builder strings.Builder
	builder.Grow(len(s))

	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			if i+2 >= len(s) {
				s = s[i:]
				return "", url.EscapeError(s)
			}
			// Some characters in CloudFront's log are escaped twice, such as,
			// '"' => %2522, '\' => %255C, and ' ' => %2520.
			// See CloudFront doc for details:
			// https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html
			// Also, please see AWS developer forum for the background:
			// https://forums.aws.amazon.com/thread.jspa?threadID=134017
			if i+4 < len(s) && s[i+1] == '2' && s[i+2] == '5' {
				if s[i+3] == '2' {
					if s[i+4] == '0' {
						builder.WriteByte(' ')
						i += 5
						continue
					}
					if s[i+4] == '2' {
						builder.WriteByte('"')
						i += 5
						continue
					}
				} else if s[i+3] == '5' && s[i+4] == 'C' {
					builder.WriteByte('\\')
					i += 5
					continue
				}
			}
			// Decode "%XX"
			l, e1 := unhex(s[i+1])
			r, e2 := unhex(s[i+2])
			if e1 != nil || e2 != nil {
				return "", url.EscapeError(s[i : i+3])
			}
			builder.WriteByte(l<<4 | r)
			i += 3
		default:
			builder.WriteByte(s[i])
			i++
		}
	}

	if builder.Len() == 0 {
		return s, nil
	}

	return builder.String(), nil
}

func MustUnescape(s string) string {
	r, err := Unescape(s)
	if err != nil {
		panic(err)
	}
	return r
}
