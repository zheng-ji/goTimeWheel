## goTimeWheel

[![Build Status](https://travis-ci.org/zheng-ji/goTimeWheel.svg)](https://travis-ci.org/zheng-ji/goTimeWheel)
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
	"github.com/zheng-ji/goTimeWheele"
)

func main() {
	tw := goTimeWheel.New(1*time.Second, 3600)

    // 3s Later the function will run
	tw.AddTimer(3 *time.Second, "ID1", func(data interface{}) {
		fmt.Printf("hello, %v\n", data)
	}, map[string]int{"age": 1})
}
```

License
-------

Copyright (c) 2019 by [zheng-ji](http://zheng-ji.info) released under MIT License.

