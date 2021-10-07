# config
Extremely simple, yet efficient, file based configuration librairy


## Features

* lazily load configuration files for quick startup
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
* [] defines an empty section.

## Example configuration
        // this is a comment
        version = 3.2.1
        # this is also a comment

        [ date ]
        created = 1/8/2019
        modified = 3/8/2019

        // above will give : config.Get("date.created") --> "1/8/2019"


## Typical use

        import ".../config"

        type struct conf {
            config.Config
            other dynamic parameters ...
            ...
        }

        // save config data
        defer conf.Save("myconf.conf")

        // access data 
        if conf.Get() == "" { .... }

        // set data
        conf.Set("name","John Doe")
