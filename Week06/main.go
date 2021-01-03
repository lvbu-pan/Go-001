package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SlidingWindow struct {
	mutex        sync.RWMutex
	window       [6]int //定义了一个子窗口为6个的数组
	cursor       int //当前窗口的游标值
	threshold    int //阈值
	cycle        int64 // 一个子窗口周期
	oldTimestamp int64
}

func NewSlidingWindow(t int) *SlidingWindow {
	return &SlidingWindow{cursor: -1, threshold: t, cycle: 10}
}

// 窗口中的计数总和
func (s *SlidingWindow) Count() int {
	var sum int
	for _, value := range s.window {
		sum = sum + value
	}
	return sum
}

// 根据时间 -> index
func (s *SlidingWindow) locationIndex() int {
	return time.Now().Second() / int(s.cycle)
}

// 更新 oldTimestamp字段
func (s *SlidingWindow) oldTimestampReset() {
	s.oldTimestamp = time.Now().Unix()
}

// 重置过期子窗口
func (s *SlidingWindow) windowSizeReset(index int) {
	s.window[index] = 0
}

// 计数器加一
func (s *SlidingWindow) increment() {
	index := s.locationIndex()
	s.mutex.Lock()
	if s.cursor != index {
		s.oldTimestampReset()
		s.windowSizeReset(index)
		s.cursor = index
	}

	// 清空过期的过期的子窗口
	if time.Now().Unix()-s.oldTimestamp >= s.cycle {
		s.oldTimestampReset()
		s.windowSizeReset(index)
	}

	s.window[index] += 1
	s.mutex.Unlock()
}

func main() {
	window := NewSlidingWindow(100)
	go func(slidingWindow *SlidingWindow) {
		rand.Seed(5)
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			slidingWindow.increment()
		}
	}(window)

	for {
		fmt.Println("Current: ", window)
		time.Sleep(time.Millisecond * 500)
		if window.Count() >= window.threshold {
			fmt.Println("溢出")
			return
		}
		fmt.Println("正常处理")
	}
}
