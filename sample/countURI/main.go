package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Maki-Daisuke/cflogparser"
	"github.com/mattn/go-forlines"
)

func main() {
	cnt := map[string]int{} // Count URI

	// Read from Stdin, parse each line and count accesses to each URI
	forlines.Do(os.Stdin, func(line string) error {
		if strings.HasPrefix(line, "#") {
			// Ignore leading comment lines for meta-information
			return nil
		}
		log, err := cflogparser.ParseLineWeb(line)  // ParseLineWeb returns *WebLog
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		cnt[log.URI]++
		return nil
	})

	for k, v := range cnt {
		fmt.Printf("%d\t%s\n", v, k)
	}
}
