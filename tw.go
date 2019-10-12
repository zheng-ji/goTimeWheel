package goTimeWheel

import (
	"container/list"
	"time"
)

// TimeWheel Struct
type TimeWheel struct {
	interval time.Duration // ticker run interval

	ticker *time.Ticker

	slots []*list.List

	keyPosMap map[interface{}]int // keep each timer's postion

	slotNum int
	currPos int // timewheel current postion

	addChannel    chan Task        // channel to  add Task
	removeChannel chan interface{} // channel to remove Task
	stopChannel   chan bool        // stop signal
}

// Task Struct
type Task struct {
	key interface{} // Timer Task ID

	delay  time.Duration // Run after delay
	circle int           // when circle equal 0 will trigger

	fn     func(interface{}) // custom function
	params interface{}       // custom parms
}

// New Func: Generate TimeWheel with ticker and slotNum
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

// Start Func: start ticker and monitor channel
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

// handle Func: Do currPosition slots Task
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

	tw.currPos = (tw.currPos + 1) % tw.slotNum
}

// getPosAndCircle Func: parse duration by interval to get circle and position
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
