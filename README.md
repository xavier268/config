# config

[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/xavier268/config)

Extremely simple, yet efficient, file based configuration librairy


## Features

* lazily load configuration files for quick startup upon first read request (Get).
  * Set will alway overwrite file configuration, even if lazily loaded after Set was called.
* get/set configuration keys
* can save modified config on disk 

## Configuration file syntax

* leading spaces and tabs ignored
* blank lines are ignored
* comment line start with # or //

* key = value
  * key has no whitespace nor tabs, it is case sensitive
  * value starts at the first character following "=" until end of line. No quotes needed. Spaces surrounding the = sign are removed.

* [ mainkey ] defines a section, where all following keys, until another section is defined, are prefixed with *mainkey.*
* [] defines a section witjhout prefix. This is the default.

## Example configuration
        // this is a comment
        version = 3.2.1
        # this is also a comment

        [ date ]
        created = 1/8/2019
        modified = 3/8/2019

        // above will give : config.Get("date.created") --> "1/8/2019"

See [examples](./examples/)

## Typical use

        import ".../config"

        conf := config.New("file1.conf", "file2.conf","/usr/bin/file3.conf", "~/bin/file4.conf")

        // save config data, including modified values.
        defer conf.Save("myconf.conf")

        // access data 
        if conf.Get() == "" { .... }

        // set data
        conf.Set("name","John Doe")


        
