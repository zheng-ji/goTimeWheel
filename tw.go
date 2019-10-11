package goTimeWheel

import (
	"container/list"
	"time"
)

type TimeWheel struct {
	interval time.Duration

	ticker *time.Ticker

	slots []*list.List

	keyPosMap map[interface{}]int

	slotNum int
	currPos int

	addChannel    chan Task
	removeChannel chan interface{}
	stopChannel   chan bool
}

type Task struct {
	delay time.Duration

	circle int
	key    interface{}

	fn     func(interface{})
	params interface{}
}

func New(interval time.Duration, slotNum int) *TimeWheel {

	if interval <= 0 || slotNum <= 0 {
		return nil
	}

	tw := &TimeWheel{
		interval:      interval,
		slots:         make([]*list.List, slotNum),
		keyPosMap:     make(map[interface{}]int),
		currPos:       0,
		slotNum:       slotNum,
		addChannel:    make(chan Task),
		removeChannel: make(chan interface{}),
		stopChannel:   make(chan bool),
	}

	for i := 0; i < slotNum; i++ {
		tw.slots[i] = list.New()
	}

	return tw
}

func (tw *TimeWheel) Start() {
	tw.ticker = time.NewTicker(tw.interval)
	go tw.start()
}

func (tw *TimeWheel) start() {
	for {
		select {
		case <-tw.ticker.C:
			tw.handle()
		case task := <-tw.addChannel:
			tw.addTask(&task)
		case key := <-tw.removeChannel:
			tw.removeTask(key)
		case <-tw.stopChannel:
			tw.ticker.Stop()
			return
		}
	}
}

func (tw *TimeWheel) Stop() {
	tw.stopChannel <- true
}

func (tw *TimeWheel) AddTimer(delay time.Duration, key interface{}, fn func(interface{}), params interface{}) {
	if delay < 0 {
		return
	}
	tw.addChannel <- Task{delay: delay, key: key, fn: fn, params: params}
}

func (tw *TimeWheel) RemoveTimer(key interface{}) {
	if key == nil {
		return
	}
	tw.removeChannel <- key
}

func (tw *TimeWheel) handle() {
	l := tw.slots[tw.currPos]

	for e := l.Front(); e != nil; {
		task := e.Value.(*Task)

		if task.circle > 0 {
			task.circle--
			continue
		}

		go task.fn(task.params)

		next := e.Next()
		l.Remove(e)
		if task.key != nil {
			delete(tw.keyPosMap, task.key)
		}
		e = next
	}

	if tw.currPos == tw.slotNum-1 {
		tw.currPos = 0
	} else {
		tw.currPos++
	}
}

func (tw *TimeWheel) getPosAndCircle(d time.Duration) (pos int, circle int) {
	circle = int(d.Seconds()) / int(tw.interval.Seconds()) / tw.slotNum
	pos = tw.currPos + int(tw.interval.Seconds())/int(tw.interval.Seconds())%tw.slotNum
	return
}

func (tw *TimeWheel) addTask(task *Task) {
	pos, circle := tw.getPosAndCircle(task.delay)
	task.circle = circle

	tw.slots[pos].PushBack(task)
	if task.key != nil {
		tw.keyPosMap[task.key] = pos
	}
}

func (tw *TimeWheel) removeTask(key interface{}) {
	pos, ok := tw.keyPosMap[key]
	if !ok {
		return
	}

	l := tw.slots[pos]

	for e := l.Front(); e != nil; {
		task := e.Value.(*Task)
		if task.key == key {
			delete(tw.keyPosMap, task.key)
			l.Remove(e)
		}

		e = e.Next()
	}
}
