package config

import (
	"bufio"
	"fmt"
	"strings"
)

// parse is called only once, to parse the configuration from disk.
// parse can happen after a first set, and should not overwite existing keys.
func (c *config) parse() {

	f := c.openConfFile()
	if f == nil {
		// no file found, we're done !
		fmt.Println("File not found ! ")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	prefix := ""
	for scanner.Scan() {
		l := strings.TrimLeft(scanner.Text(), " \t\n\r")
		switch {
		case strings.HasPrefix(l, "#"):
			fmt.Println("ignoring : ", l)
		case strings.HasPrefix(l, "//"):
			fmt.Println("ignoring : ", l)
		case strings.HasPrefix(l, "["):
			fmt.Println("prefix found : ", l)
			pp := strings.SplitN(l[1:], "]", 1)
			if len(pp) == 2 {
				prefix = strings.Trim(pp[0], " \t\n\r")
			} else {
				prefix = ""
			}
		default:
			kk := strings.SplitAfterN(l, "=", 1)
			if len(kk) == 2 {
				key := strings.Trim(kk[0], " \t\r\n")
				if len(prefix) != 0 {
					key = prefix + "." + key
				}
				value := kk[1]
				// TODO - store key,value
				fmt.Println(key, "-->", value)

			}
		}
	}
	fmt.Println(c.values)
}
