package goTimeWheel

import (
	"fmt"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {

	tw := New(1*time.Second, 3600)

	tw.Start()

	name := "ID1"
	params := map[string]int{"age": 1}
	fn := func(data interface{}) {
		fmt.Printf("hello, %v\n", data)
	}
	tw.AddTimer(3*time.Second, name, fn, params)

	time.Sleep(time.Duration(5) * time.Second)
}

func TestRmTask(t *testing.T) {

	tw := New(1*time.Second, 3600)

	tw.Start()

	tw.AddTimer(3*time.Second, "key1", func(data interface{}) {
		fmt.Printf("hello, %v\n", data)
	}, map[string]int{"age": 1})

	tw.RemoveTimer("key1")

	time.Sleep(time.Duration(5) * time.Second)
}

func TestStopTimeWheel(t *testing.T) {

	tw := New(1*time.Second, 3600)

	tw.Start()

	tw.AddTimer(3*time.Second, "key1", func(data interface{}) {
		fmt.Printf("hello, %v\n", data)
	}, map[string]int{"age": 1})

	tw.Stop()

	time.Sleep(time.Duration(5) * time.Second)
}
