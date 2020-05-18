# Config



## ConsulKV
```golang
package main

import (
    "github.com/yautz/go-tools/config/consul"
)

func main() {
    v , err := config.ConsulKV("consulhost" , "key" , "ftype{json/yaml}")
    if err != nil {
        log.Fatal(err)
    }

    v.GetString("xxx.xxx")


    x := config.Database{}

    // maping to struct
    err := viper.Unmarshal(&x)

}

```
