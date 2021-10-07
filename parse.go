package config

import (
	"bufio"
	"strings"
)

const SP = " \t\r\n"

// parse is called only once, to parse the configuration from disk.
// parse can happen after a first set, and should not overwite existing keys.
func (c *config) parse() {

	f := c.openConfFile()
	if f == nil {
		// no file found, we're done !
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	prefix := ""
	for scanner.Scan() {
		l := strings.TrimLeft(scanner.Text(), SP)
		switch {
		case strings.HasPrefix(l, "#"):
		case strings.HasPrefix(l, "//"):
		case strings.HasPrefix(l, "["):

			pp := strings.SplitN(l[1:], "]", 2)
			if len(pp) == 2 {
				prefix = strings.Trim(pp[0], SP)
			} else {
				prefix = ""
			}
		case l == "":
		default:
			kk := strings.SplitN(l, "=", 2)
			if len(kk) == 2 {
				key := strings.Trim(kk[0], SP)
				if len(prefix) != 0 {
					key = prefix + "." + key
				}
				value := strings.TrimLeft(kk[1], SP)
				c.Set(key, value)
			}
		}
	}
}
