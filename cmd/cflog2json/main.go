package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Maki-Daisuke/cflogparser"
	"github.com/Maki-Daisuke/go-argvreader"
	"github.com/mattn/go-forlines"
)

func main() {
	rd := argvreader.New()

	for {
		err := forlines.Do(rd, func(line string) error {
			if strings.HasPrefix(line, "#") {
				// Ignore leading comment lines for meta-information
				return nil
			}

			l, err := cflogparser.ParseLineWeb(line)
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
