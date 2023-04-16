package main

import (
	"sync"
	"time"
)

type Cache struct {
	storage map[string]int
	mu      sync.RWMutex
}

func (c *Cache) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
}

const (
	k1   = "key1"
	step = 7
)

type Semaphore struct {
	sem chan struct{}
}

// Создание нового семафора
func newSemaphore(n int) *Semaphore {
	return &Semaphore{make(chan struct{}, n)}
}

// Блокировка, если количество одновременно работающих горутин уже достигло максимально заданного значения.
func (s *Semaphore) Acquire() {
	s.sem <- struct{}{}
}

// Освобождение ресурса, удаляя его из канала sem и позволяя заблокированным горутинам продолжить работу.
func (s *Semaphore) Release() {
	<-s.sem
}

// Подсчет количества элементов в буфере
func (s *Semaphore) Len() int {
	return len(s.sem)
}

func main() {
	cache := Cache{storage: make(map[string]int)}

	semaphore1 := newSemaphore(3)
	semaphore2 := newSemaphore(3)

	for i := 0; i < 10; i++ {
		go func() {
			semaphore1.Acquire()
			defer semaphore1.Release()
			cache.Increase(k1, step)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			semaphore2.Acquire()
			defer semaphore2.Release()
			cache.Set(k1, step*i)
			time.Sleep(time.Millisecond * 100)
		}(i)
	}
	if semaphore1.Len() > 0 || semaphore2.Len() > 0 {
		time.Sleep(time.Second * 10)
	}
}
