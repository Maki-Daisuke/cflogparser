package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Maki-Daisuke/cflogparser"
	"github.com/Maki-Daisuke/go-argvreader"
	"github.com/mattn/go-forlines"
)

func main() {
	var optRTMP bool
	flag.BoolVar(&optRTMP, "rtmp", false, "Parse input as RTMP distribution log")
	flag.Parse()

	rd := argvreader.NewReader(flag.Args())
	for {
		err := forlines.Do(rd, func(line string) error {
			if strings.HasPrefix(line, "#") {
				// Ignore leading comment lines for meta-information
				return nil
			}

			var l interface{}
			var err error
			if !optRTMP {
				l, err = cflogparser.ParseLineWeb(line)
			} else {
				l, err = cflogparser.ParseLineRTMP(line)
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}
			b, err := json.Marshal(l)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}
			_, err = os.Stdout.Write(b)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}
			os.Stdout.Write([]byte{'\n'})
			return nil
		})
		if err == nil {
			break
		}
		// Here, err is not EOF. So, report error to STDERR, and then, continue reading.
		fmt.Fprintln(os.Stderr, err)
	}
}
