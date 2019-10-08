cflogparser
===========

Description
-----------

A parser for Amazon Web Services (AWS) CloudFront log implemented in Go


Example
-------

```golang
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
```


Supported Formtats
------------------

- CloudFront Access log
  - Web Distribution Log File Format: Version 1.0
  - RTMP Distribution Log File Format: Version 1.0


Ratinale
--------

CloudFront log format is very simple and looks easy to parse. It is essentially 
URL-encoded values delimitted by `'\t'` (tab) character. But, there are putfalls.

For example, it escape some characters twice, such as escaping `' '` (space) to `"%2520"`. 
(See [this AWS Discussion Forums post](https://forums.aws.amazon.com/thread.jspa?threadID=134017) 
for the historical background.)

Another example is that it records `"-"` when there is no value to record even though that 
field expect an integer value.

We needed something that strictly complies with 
[the specification](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html).


License
-------

MIT License


Author
------

Daisuke (yet another) Maki
