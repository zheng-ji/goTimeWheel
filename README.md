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

```go
import (
    "fmt"
    "github.com/zheng-ji/goTimeWheel"
)

func main() {
    // timer ticker
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
    select{}

}
```

License
-------

Copyright (c) 2019 by [zheng-ji](http://zheng-ji.info) released under MIT License.

