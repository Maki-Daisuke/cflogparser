package cflogparser

import (
	"net"
	"reflect"
	"testing"
	"time"
)

func TestParseWebLogSamples(t *testing.T) {
	tests := []struct {
		in  string
		out *WebLog
	}{
		{
			// Log samples is borrowed from here: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html
			`2014-05-23	01:13:11	FRA2	182	192.0.2.10	GET	d111111abcdef8.cloudfront.net	/view/my/file.html	200	www.displaymyfiles.com	Mozilla/4.0%20(compatible;%20MSIE%205.0b1;%20Mac_PowerPC)	-	zip=98101	RefreshHit	MRVMF7KydIvxMWfJIglgwHQwZsbG2IhRJ07sn9AkKUFSHS9EXAMPLE==	d111111abcdef8.cloudfront.net	http	-	0.001	-	-	-	RefreshHit	HTTP/1.1	Processed	1`,
			&WebLog{
				Time:               time.Date(2014, 5, 23, 1, 13, 11, 0, time.UTC),
				Location:           "FRA2",
				Bytes:              182,
				RequestIP:          net.ParseIP("192.0.2.10"),
				Method:             "GET",
				Host:               "d111111abcdef8.cloudfront.net",
				URI:                "/view/my/file.html",
				Status:             200,
				Referrer:           "www.displaymyfiles.com",
				UserAgent:          "Mozilla/4.0 (compatible; MSIE 5.0b1; Mac_PowerPC)",
				QueryString:        "",
				Cookie:             "zip=98101",
				ResultType:         "RefreshHit",
				RequestID:          "MRVMF7KydIvxMWfJIglgwHQwZsbG2IhRJ07sn9AkKUFSHS9EXAMPLE==",
				HostHeader:         "d111111abcdef8.cloudfront.net",
				RequestProtocol:    "http",
				RequestBytes:       0,
				TimeTaken:          0.001,
				XforwardedFor:      "",
				SslProtocol:        "",
				SslCipher:          "",
				ResponseResultType: "RefreshHit",
				HTTPVersion:        "HTTP/1.1",
				FleStatus:          "Processed",
				FleEncryptedFields: 1,
			},
		},
		{
			`2014-05-23	01:13:12	LAX1	2390282	192.0.2.202	GET	d111111abcdef8.cloudfront.net	/soundtrack/happy.mp3	304	www.unknownsingers.com	Mozilla/4.0%20(compatible;%20MSIE%207.0;%20Windows%20NT%205.1)	a=b&c=d	zip=50158	Hit	xGN7KWpVEmB9Dp7ctcVFQC4E-nrcOcEKS3QyAez--06dV7TEXAMPLE==	d111111abcdef8.cloudfront.net	http	-	0.002	-	-	-	Hit	HTTP/1.1	-	-`,
			&WebLog{
				Time:               time.Date(2014, 5, 23, 1, 13, 12, 0, time.UTC),
				Location:           "LAX1",
				Bytes:              2390282,
				RequestIP:          net.ParseIP("192.0.2.202"),
				Method:             "GET",
				Host:               "d111111abcdef8.cloudfront.net",
				URI:                "/soundtrack/happy.mp3",
				Status:             304,
				Referrer:           "www.unknownsingers.com",
				UserAgent:          "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1)",
				QueryString:        "a=b&c=d",
				Cookie:             "zip=50158",
				ResultType:         "Hit",
				RequestID:          "xGN7KWpVEmB9Dp7ctcVFQC4E-nrcOcEKS3QyAez--06dV7TEXAMPLE==",
				HostHeader:         "d111111abcdef8.cloudfront.net",
				RequestProtocol:    "http",
				RequestBytes:       0,
				TimeTaken:          0.002,
				XforwardedFor:      "",
				SslProtocol:        "",
				SslCipher:          "",
				ResponseResultType: "Hit",
				HTTPVersion:        "HTTP/1.1",
				FleStatus:          "",
				FleEncryptedFields: 0,
			},
		},
	}
	for _, test := range tests {
		log, err := ParseLineWeb(test.in)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(log, test.out) {
			t.Errorf(`got %v, want %v`, log, test.out)
		}
	}
}
