package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		mu:       sync.RWMutex{},
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.items[key]

	if !ok {
		if c.queue.Len() >= c.capacity {
			back := c.queue.Back()
			c.queue.Remove(back)
			delete(c.items, back.Key)
		}

		item := c.queue.PushFront(value)
		item.Key = key

		c.items[key] = item

		return false
	}

	item.Value = value
	c.queue.PushFront(item)

	return true
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.RLock()
	item, ok := c.items[key]
	c.mu.RUnlock()

	if !ok {
		return nil, false
	}

	c.mu.Lock()
	c.queue.PushFront(item)
	c.mu.Unlock()

	return item.Value, true
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[Key]*ListItem)
	c.queue.Clear()
}
