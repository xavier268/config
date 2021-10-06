# config
Extremely simple, yet efficient, file based configuration librairy


## Features

* lazily load files for quick startup
* get/set configuration keys
* can save modified config on disk 

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
