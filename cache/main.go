package main

import "time"

type Cache struct {
	storage map[string]int
}

func (c *Cache) Increase(key string, value int) {
	if _, ok := c.storage[key]; !ok {
		c.storage[key] = value
	} else {
		c.storage[key] += value
	}
}

func (c *Cache) Set(key string, value int) {
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	delete(c.storage, key)
}
