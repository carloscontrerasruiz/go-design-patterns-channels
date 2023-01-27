package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonnacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonnacci(n-1) + Fibonnacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock  *sync.Mutex
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exists := m.cache[key]
	m.lock.Unlock()

	if !exists {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
	}

	return result.value, result.err
}

type Function func(key int) (interface{}, error)
type FunctionResult struct {
	value interface{}
	err   error
}

var lock sync.Mutex

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
		lock:  &lock,
	}
}

func GetFibonnacci(n int) (interface{}, error) {
	return Fibonnacci(n), nil
}

func main() {
	cache := NewCache(GetFibonnacci)
	fibo := []int{42, 40, 41, 42, 38, 40, 41}
	var wg sync.WaitGroup

	for _, v := range fibo {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, _ := cache.Get(index)
			fmt.Printf("%d, %s, %d\n", index, time.Since(start), value)
		}(v)

	}

	wg.Wait()
}
