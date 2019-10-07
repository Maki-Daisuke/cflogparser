package cflogparser

import (
	"net"
	"reflect"
	"testing"
	"time"
)

func TestParseRTMPLogSamples(t *testing.T) {
	tests := []struct {
		in  string
		out *RTMPLog
	}{
		// Log samples is borrowed from here: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html
		{
			`2010-03-12	23:51:20	SEA4	192.0.2.147	connect	2014	OK	bfd8a98bee0840d9b871b7f6ade9908f	rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st​	key=value	http://player.longtailvideo.com/player.swf	http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204	LNX%2010,0,32,18	-	-	-	-`,
			&RTMPLog{
				Time:          time.Date(2010, 3, 12, 23, 51, 20, 0, time.UTC),
				Location:      "SEA4",
				RequestIP:     net.ParseIP("192.0.2.147"),
				EventType:     "connect",
				Bytes:         2014,
				Status:        "OK",
				ClientID:      "bfd8a98bee0840d9b871b7f6ade9908f",
				URI:           "rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st​",
				QueryString:   "key=value",
				Referrer:      "http://player.longtailvideo.com/player.swf",
				PageURL:       "http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204",
				UserAgent:     "LNX 10,0,32,18",
				StreamName:    "",
				StreamQuery:   "",
				StreamFileExt: "",
				StreamID:      0,
			},
		},
		{
			`2010-03-12	23:51:21	SEA4	192.0.2.222	play	3914	OK	bfd8a98bee0840d9b871b7f6ade9908f	rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st	key=value	http://player.longtailvideo.com/player.swf	http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204	LNX%2010,0,32,18	myvideo	p=2&q=4	flv	1`,
			&RTMPLog{
				Time:          time.Date(2010, 3, 12, 23, 51, 21, 0, time.UTC),
				Location:      "SEA4",
				RequestIP:     net.ParseIP("192.0.2.222"),
				EventType:     "play",
				Bytes:         3914,
				Status:        "OK",
				ClientID:      "bfd8a98bee0840d9b871b7f6ade9908f",
				URI:           "rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st",
				QueryString:   "key=value",
				Referrer:      "http://player.longtailvideo.com/player.swf",
				PageURL:       "http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204",
				UserAgent:     "LNX 10,0,32,18",
				StreamName:    "myvideo",
				StreamQuery:   "p=2&q=4",
				StreamFileExt: "flv",
				StreamID:      1,
			},
		},
		{
			`2010-03-12	23:53:44	SEA4	192.0.2.4	stop	323914	OK	bfd8a98bee0840d9b871b7f6ade9908f	rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st	key=value	http://player.longtailvideo.com/player.swf	http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204	LNX%2010,0,32,18	dir/other/myvideo	p=2&q=4	flv	1`,
			&RTMPLog{
				Time:          time.Date(2010, 3, 12, 23, 53, 44, 0, time.UTC),
				Location:      "SEA4",
				RequestIP:     net.ParseIP("192.0.2.4"),
				EventType:     "stop",
				Bytes:         323914,
				Status:        "OK",
				ClientID:      "bfd8a98bee0840d9b871b7f6ade9908f",
				URI:           "rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st",
				QueryString:   "key=value",
				Referrer:      "http://player.longtailvideo.com/player.swf",
				PageURL:       "http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204",
				UserAgent:     "LNX 10,0,32,18",
				StreamName:    "dir/other/myvideo",
				StreamQuery:   "p=2&q=4",
				StreamFileExt: "flv",
				StreamID:      1,
			},
		},
		{
			`2010-03-12	23:53:44	SEA4	192.0.2.103	play	8783724	OK	bfd8a98bee0840d9b871b7f6ade9908f	rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st	key=value	http://player.longtailvideo.com/player.swf	http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204	LNX%2010,0,32,18	dir/favs/myothervideo	p=42&q=14	mp4	2`,
			&RTMPLog{
				Time:          time.Date(2010, 3, 12, 23, 53, 44, 0, time.UTC),
				Location:      "SEA4",
				RequestIP:     net.ParseIP("192.0.2.103"),
				EventType:     "play",
				Bytes:         8783724,
				Status:        "OK",
				ClientID:      "bfd8a98bee0840d9b871b7f6ade9908f",
				URI:           "rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st",
				QueryString:   "key=value",
				Referrer:      "http://player.longtailvideo.com/player.swf",
				PageURL:       "http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204",
				UserAgent:     "LNX 10,0,32,18",
				StreamName:    "dir/favs/myothervideo",
				StreamQuery:   "p=42&q=14",
				StreamFileExt: "mp4",
				StreamID:      2,
			},
		},
		{
			`2010-03-12	23:56:21	SEA4	192.0.2.199	stop	429822014	OK	bfd8a98bee0840d9b871b7f6ade9908f	rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st	key=value	http://player.longtailvideo.com/player.swf	http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204	LNX%2010,0,32,18	dir/favs/myothervideo	p=42&q=14	mp4	2`,
			&RTMPLog{
				Time:          time.Date(2010, 3, 12, 23, 56, 21, 0, time.UTC),
				Location:      "SEA4",
				RequestIP:     net.ParseIP("192.0.2.199"),
				EventType:     "stop",
				Bytes:         429822014,
				Status:        "OK",
				ClientID:      "bfd8a98bee0840d9b871b7f6ade9908f",
				URI:           "rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st",
				QueryString:   "key=value",
				Referrer:      "http://player.longtailvideo.com/player.swf",
				PageURL:       "http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204",
				UserAgent:     "LNX 10,0,32,18",
				StreamName:    "dir/favs/myothervideo",
				StreamQuery:   "p=42&q=14",
				StreamFileExt: "mp4",
				StreamID:      2,
			},
		},
		{
			`2010-03-12	23:59:44	SEA4	192.0.2.14	disconnect	429824092	OK	bfd8a98bee0840d9b871b7f6ade9908f	rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st	key=value	http://player.longtailvideo.com/player.swf	http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204	LNX%2010,0,32,18	-	-	-	-`,
			&RTMPLog{
				Time:          time.Date(2010, 3, 12, 23, 59, 44, 0, time.UTC),
				Location:      "SEA4",
				RequestIP:     net.ParseIP("192.0.2.14"),
				EventType:     "disconnect",
				Bytes:         429824092,
				Status:        "OK",
				ClientID:      "bfd8a98bee0840d9b871b7f6ade9908f",
				URI:           "rtmp://shqshne4jdp4b6.cloudfront.net/cfx/st",
				QueryString:   "key=value",
				Referrer:      "http://player.longtailvideo.com/player.swf",
				PageURL:       "http://www.longtailvideo.com/support/jw-player-setup-wizard?example=204",
				UserAgent:     "LNX 10,0,32,18",
				StreamName:    "",
				StreamQuery:   "",
				StreamFileExt: "",
				StreamID:      0,
			},
		},
	}
	for _, test := range tests {
		log, err := ParseLineRTMP(test.in)
		if err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(log, test.out) {
			t.Errorf(`got %v, want %v`, log, test.out)
		}
	}
}
