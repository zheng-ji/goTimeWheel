## goTimeWheel

[![Build Status](https://travis-ci.org/zheng-ji/goTimeWheel.svg)](https://travis-ci.org/zheng-ji/goTimeWheel)
[![codecov](https://codecov.io/gh/zheng-ji/goTimeWheel/branch/master/graph/badge.svg)](https://codecov.io/gh/zheng-ji/goTimeWheel)
[![GoDoc](https://godoc.org/github.com/zheng-ji/goTimeWheele?status.svg)](https://godoc.org/github.com/zheng-ji/goTimeWheel)

TimeWheel Implemented By Go.
Go 实现的时间轮，俗称定时器

![goTimeWheel](https://github.com/zheng-ji/goTimeWheel/blob/master/goTimeWheel.png)

Feature
--------

* Effective at Space Usage
* Each Timer Can Custom Its Task


Installation
-------------

```
go get github.com/zheng-ji/goTimeWheel
```

Example
-------

```
package main

import (
	"github.com/euclidr/bloomf"
    "github.com/go-redis/redis"
)


func main() {
	client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    bl, err := bloomf.New(client, "bf" 1000000, 0.001)
    if err == bloomf.ErrDuplicated {
        bl, _ := bloomf.GetByName(client, "bf")
    }

    bl.Add([]bytes("awesome key"))

    exists, _ := bl.Exists([]bytes("awesome key"))
    if exists {
        ...
    }

    ...

    // bl.Clear() // you can clean up all datas in redis
}
```

```go
import (
    "fmt"
    "github.com/zheng-ji/goTimeWheele"
)

func main() {
    tw := goTimeWheel.New(1*time.Second, 3600)
    tw.Start()
    // "ID1" means the timer's name
    // Specify a function and params, it will run after 3s later
    name := "ID1"
    params := map[string]int{"age": 1}
    fn := func(data interface{}) {
        fmt.Printf("hello, %v\n", data)
    }
    tw.AddTimer(3*time.Second, name, fn, params)
    // Your Logic Code
}
```

License
-------

Copyright (c) 2019 by [zheng-ji](http://zheng-ji.info) released under MIT License.

