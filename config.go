package config

import (
	"fmt"
	"hash/maphash"
	"os"
	"sort"
	"strings"
	"sync"
)

type Config interface {
	Get(key string) string
	Set(key, value string)
	Save(fileName string) error
}

type config struct {
	files  []string          // files to look for configuration, only the first one that exists is taken into account.
	keys   []string          // parsed keys - keys are case sensitive by default. This is required to save ...
	values map[uint64]string // prsed values, in a hash map
	once   sync.Once         // used for lazy initialisation, upon first get of a key value (setting will not trigger a parse)
	h      maphash.Hash      // used to map string keys to their hashed value
}

// defaults serached files added to the ones specified
var DefaultFiles = []string{"config.conf", "~/config.conf", "../config.conf", "/config.conf"}

// New creates a new Config.
// Provided file names will be searched, in order, then the DefaultFiles, until a valid file is found.
// If none is found, all keys have an empty string value.
func New(files ...string) Config {
	return newConfig(files...)
}

// newConfig returns the actual config object for testing purposes.
func newConfig(files ...string) *config {
	c := new(config)
	c.files = files
	c.keys = make([]string, 0, 10)
	c.values = make(map[uint64]string)
	return c
}

func (c *config) Get(k string) (v string) {
	c.once.Do(c.parse)
	c.h.Reset()
	c.h.WriteString(k)
	return c.values[c.h.Sum64()]
}

func (c *config) Set(k, v string) {
	set, h := c.isSet(k)
	if !set {
		// if key does not exists yet, take note of it !
		c.keys = append(c.keys, k)
	}
	c.values[h] = v
}

func (c *config) isSet(key string) (bool, uint64) {
	c.h.Reset()
	c.h.WriteString(key)
	h := c.h.Sum64()
	_, ok := c.values[h]
	return ok, h
}

func (c *config) Save(fname string) error {
	c.once.Do(c.parse)
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	sort.Strings(c.keys)

	var p, pp, kk string
	for _, k := range c.keys {
		pp, kk = getPrefix(k)
		if pp != p {
			fmt.Fprintf(f, "[%s]\n", pp)
			p = pp
		}
		fmt.Fprintf(f, "%s=%s\n", kk, c.Get(k))
	}
	return err
}

func (c *config) openConfFile() *os.File {
	var f *os.File
	var err error
	for _, fn := range append(c.files, DefaultFiles...) {
		f, err = os.Open(fn)
		if err == nil {
			return f
		}
	}
	return nil
}

// Extract the prefix from a key
// prefix may be the empty string, and may contain inside '.', but not the final one.
func getPrefix(key string) (prefix, shortkey string) {
	kk := strings.Split(key, ".")
	if len(kk) <= 1 {
		return "", key
	}

	return strings.Join(kk[0:len(kk)-1], "."), kk[len(kk)-1]
}
