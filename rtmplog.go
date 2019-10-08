package cflogparser

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// RTMPLog represents a record of RTMP distribution log.
type RTMPLog struct {
	Time          time.Time `json:"time"`
	Location      string    `json:"location"`
	RequestIP     net.IP    `json:"request_ip"`
	EventType     string    `json:"event_type"`
	Bytes         uint64    `json:"bytes"`
	Status        string    `json:"status"`
	ClientID      string    `json:"client_id"`
	URI           string    `json:"uri"`
	QueryString   string    `json:"query_string"`
	Referrer      string    `json:"referrer"`
	PageURL       string    `json:"page_url"`
	UserAgent     string    `json:"user_agent"`
	StreamName    string    `json:"stream_name"`
	StreamQuery   string    `json:"stream_query"`
	StreamFileExt string    `json:"stream_file_ext"`
	StreamID      uint32    `json:"stream_id"`
}

// ParseLineRTMP parses a line of log for a RTMP distribution.
func ParseLineRTMP(line string) (l *RTMPLog, err error) {
	vals := strings.Split(line, "\t")
	if len(vals) < 17 {
		return nil, fmt.Errorf("Insufficient number of fields: %s", line)
	}

	l = &RTMPLog{}

	defer func() {
		if r := recover(); r != nil {
			l = nil
			err = fmt.Errorf("Can't parse line: %w: %s", r, line)
		}
	}()

	l.Time, err = time.Parse("2006-01-02 15:04:05", vals[0]+" "+vals[1])
	if err != nil {
		panic(err)
	}

	l.Location = parseString(vals[2])

	l.RequestIP = net.ParseIP(vals[3])
	if l.RequestIP == nil {
		panic(err)
	}

	l.EventType = parseString(vals[4])
	l.Bytes = uint64(parseInt(vals[5]))
	l.Status = parseString(vals[6])
	l.ClientID = parseString(vals[7])
	l.URI = parseString(vals[8])
	l.QueryString = parseString(vals[9])
	l.Referrer = parseString(vals[10])
	l.PageURL = parseString(vals[11])
	l.UserAgent = parseString(vals[12])
	l.StreamName = parseString(vals[13])
	l.StreamQuery = parseString(vals[14])
	l.StreamFileExt = parseString(vals[15])
	l.StreamID = uint32(parseInt(vals[16]))

	return l, nil
}
