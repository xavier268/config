package config

import (
	"hash/maphash"
	"os"
	"sync"
)

type Config interface {
	Get(key string) string
	Set(key, value string)
	Save(fileName string)
}

type config struct {
	files  []string          // files to look for configuration, only the first one that exists is taken into account.
	keys   []string          // parsed keys - keys are case sensitive by default. This is required to save ...
	values map[uint64]string // prsed values, in a hash map
	once   sync.Once         // used for lazy initialisation, upon first get of a key value (setting will not trigger a parse)
	h      maphash.Hash      // used to map string keys to their hashed value
}

// defaults serached files added to the ones specified
var defaultFiles = []string{"config.conf", "default.conf"}

// New creates a new Config
// file names will be searched, in the provided order, then the default files, until a valid file is fouend.
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
	c.h.Reset()
	c.h.WriteString(k)
	h := c.h.Sum64()
	if _, ok := c.values[h]; !ok {
		// if key does not exists yet, take note of it !
		c.keys = append(c.keys, k)
	}
	c.values[h] = v
}

func (c *config) Save(fname string) {
	// todo
}

func (c *config) openConfFile() *os.File {
	var f *os.File
	var err error
	for _, fn := range append(c.files, defaultFiles...) {
		f, err = os.Open(fn)
		if err == nil {
			return f
		}
	}
	return nil
}
