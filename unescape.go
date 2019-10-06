package cflogparser

import (
	"net/url"
	"strings"
)

func ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func Unescape(s string) (string, error) {
	var builder strings.Builder
	builder.Grow(len(s))
	
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			if i+2 >= len(s) || !ishex(s[i+1]) || !ishex(s[i+2]) {
				s = s[i:]
				if len(s) > 3 {
					s = s[:3]
				}
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
			builder.WriteByte(unhex(s[i+1])<<4 | unhex(s[i+2]))
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
