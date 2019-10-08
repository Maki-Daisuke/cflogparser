package cflogparser

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// WebLog represents a record of CloudFront's Web distribution log.
// Naming convention is borrowed from this article:
// https://docs.aws.amazon.com/athena/latest/ug/cloudfront-logs.html
type WebLog struct {
	Time               time.Time `json:"time"`
	Location           string    `json:"location"`
	Bytes              uint64    `json:"bytes"`
	RequestIP          net.IP    `json:"request_ip"`
	Method             string    `json:"method"`
	Host               string    `json:"host"`
	URI                string    `json:"uri"`
	Status             uint16    `json:"status"`
	Referrer           string    `json:"referrer"`
	UserAgent          string    `json:"user_agent"`
	QueryString        string    `json:"query_string"`
	Cookie             string    `json:"cookie"`
	ResultType         string    `json:"result_type"`
	RequestID          string    `json:"request_id"`
	HostHeader         string    `json:"host_header"`
	RequestProtocol    string    `json:"request_protocol"`
	RequestBytes       uint64    `json:"request_bytes"`
	TimeTaken          float32   `json:"time_taken"`
	XforwardedFor      string    `json:"xforwarded_for"`
	SslProtocol        string    `json:"ssl_protocol"`
	SslCipher          string    `json:"ssl_cipher"`
	ResponseResultType string    `json:"response_result_type"`
	HTTPVersion        string    `json:"http_version"`
	FleStatus          string    `json:"fle_status"`
	FleEncryptedFields uint32    `json:"fle_encrypted_fields"`
}

// ParseLineWeb parses a line of log for a web distribution.
func ParseLineWeb(line string) (l *WebLog, err error) {
	vals := strings.Split(line, "\t")
	if len(vals) < 26 {
		return nil, fmt.Errorf("Insufficient number of fields: %s", line)
	}

	l = &WebLog{}

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
	l.Bytes = uint64(parseInt(vals[3]))

	l.RequestIP = net.ParseIP(vals[4])
	if l.RequestIP == nil {
		panic(err)
	}

	l.Method = parseString(vals[5])
	l.Host = parseString(vals[6])
	l.URI = parseString(vals[7])
	l.Status = uint16(parseInt(vals[8]))
	l.Referrer = parseString(vals[9])
	l.UserAgent = parseString(vals[10])
	l.QueryString = parseString(vals[11])
	l.Cookie = parseString(vals[12])
	l.ResultType = parseString(vals[13])
	l.RequestID = parseString(vals[14])
	l.HostHeader = parseString(vals[15])
	l.RequestProtocol = parseString(vals[16])
	l.RequestBytes = uint64(parseInt(vals[17]))
	l.TimeTaken = float32(parseFloat(vals[18]))
	l.XforwardedFor = parseString(vals[19])
	l.SslProtocol = parseString(vals[20])
	l.SslCipher = parseString(vals[21])
	l.ResponseResultType = parseString(vals[22])
	l.HTTPVersion = parseString(vals[23])
	l.FleStatus = parseString(vals[24])
	l.FleEncryptedFields = uint32(parseInt(vals[25]))

	return l, nil
}

func parseString(f string) string {
	if f == "-" {
		return ""
	}
	return MustUnescape(f)
}

func parseInt(f string) int64 {
	if f == "-" {
		return 0
	}
	n, err := strconv.ParseInt(f, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func parseFloat(f string) float64 {
	if f == "-" {
		return 0.0
	}
	n, err := strconv.ParseFloat(f, 64)
	if err != nil {
		panic(err)
	}
	return n
}
