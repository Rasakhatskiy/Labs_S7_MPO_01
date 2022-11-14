package cell

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Cell interface {
	GetValue() int
	SetValue(int)
	SetIndex(int)
	Increment()
	Decrement()
	String() string
	Run(seconds int64, cells []Cell, p float64, ch chan int)
}

type SafeCell struct {
	mutex sync.Mutex
	value int
	index int
}

func (c *SafeCell) SetIndex(i int) {
	c.mutex.Lock()
	c.index = i
	c.mutex.Unlock()
}

func (c *SafeCell) Run(seconds int64, cells []Cell, p float64, ch chan int) {
	start := time.Now()
	elapsed := time.Since(start)
	for elapsed.Milliseconds() < seconds*1000 {
		if flipCoin(0.5) {
			// left
			if c.index != 0 && flipCoin(p) && c.GetValue() != 0 {
				cells[c.index-1].Increment()
				c.Decrement()
				continue
			}
		} else {
			// right
			if c.index != len(cells)-1 && flipCoin(p) && c.GetValue() != 0 {
				cells[c.index+1].Increment()
				c.Decrement()
				continue
			}
		}
		elapsed = time.Since(start)
	}
	ch <- c.GetValue()
}

func (c *SafeCell) String() string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return strconv.Itoa(c.value)
}

func (c *SafeCell) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

func (c *SafeCell) Decrement() {
	c.mutex.Lock()
	c.value--
	c.mutex.Unlock()
}

func (c *SafeCell) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func (c *SafeCell) SetValue(v int) {
	c.mutex.Lock()
	c.value = v
	c.mutex.Unlock()
}

type UnsafeCell struct {
	value int
	index int
}

func (c *UnsafeCell) SetIndex(i int) {
	c.index = i
}

func (c *UnsafeCell) Run(seconds int64, cells []Cell, p float64, ch chan int) {
	start := time.Now()
	elapsed := time.Since(start)
	for elapsed.Milliseconds() < seconds*1000 {
		if flipCoin(0.5) {
			// left
			if c.index != 0 && flipCoin(p) && c.GetValue() != 0 {
				cells[c.index-1].Increment()
				c.Decrement()
				continue
			}
		} else {
			// right
			if c.index != len(cells)-1 && flipCoin(p) && c.GetValue() != 0 {
				cells[c.index+1].Increment()
				c.Decrement()
				continue
			}
		}
		elapsed = time.Since(start)
	}
	ch <- c.GetValue()
}

func (c *UnsafeCell) String() string {
	return strconv.Itoa(c.value)
}

func (c *UnsafeCell) GetValue() int {
	return c.value
}

func (c *UnsafeCell) SetValue(v int) {
	c.value = v
}

func (c *UnsafeCell) Increment() {
	c.value++
}

func (c *UnsafeCell) Decrement() {
	c.value--
}

func flipCoin(p float64) bool {
	if p == 0 {
		return false
	}
	if p == 1 {
		return true
	}
	r := rand.Float64()
	return r < p
}
