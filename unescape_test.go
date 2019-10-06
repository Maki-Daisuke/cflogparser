package cflogparser

import (
	"testing"
	"net/url"
)

func TestUserAgents(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"Mozilla/5.0%2520(compatible;%2520bingbot/2.0;%2520+http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"},
		{"Mozilla/5.0%2520(iPhone;%2520CPU%2520iPhone%2520OS%252011_0%2520like%2520Mac%2520OS%2520X)%2520AppleWebKit/604.1.38%2520(KHTML,%2520like%2520Gecko)%2520Version/11.0%2520Mobile/15A372%2520Safari/604.1", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"},
		{"Mozilla/5.0%2520(Windows%2520NT%252010.0;%2520Win64;%2520x64;%2520rv:61.0)%2520Gecko/20100101%2520Firefox/61.0", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:61.0) Gecko/20100101 Firefox/61.0"},
		{"Mozilla/5.0%2520(Windows%2520NT%25206.1;%2520WOW64;%2520Trident/7.0;%2520rv:11.0)%2520like%2520Gecko", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko"},
	}
	for _, test := range tests {
		r := MustUnescape(test.in)
		if r != test.out {
			t.Errorf("got %q, want %q", r, test.out)
		}
	}
}

func TestInvalid(t *testing.T) {
	_, err := Unescape("aaaa%XXaaa")
	if ee, ok := err.(url.EscapeError); !ok || string(ee) != "%XX" {
		t.Error(err)
	}
}
