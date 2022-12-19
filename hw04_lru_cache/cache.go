package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mx       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	// Добавляет значение в кэш по ключу
	cache.mx.Lock()
	defer cache.mx.Unlock()
	item := cacheItem{key, value}

	if i, ok := cache.items[key]; ok {
		cache.items[key].Value = item
		cache.queue.MoveToFront(i)
		return true
	}
	if cache.queue.Len() >= cache.capacity {
		lastItem := cache.queue.Back()
		cache.queue.Remove(lastItem)
		delete(cache.items, lastItem.Value.(cacheItem).key)
	}
	cache.queue.PushFront(item)
	cache.items[key] = cache.queue.Front()
	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	// Возвращает значение из кэша по ключу
	cache.mx.Lock()
	defer cache.mx.Unlock()
	if item, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(item)
		cacheValue := cache.items[key].Value
		return cacheValue.(cacheItem).value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	// Очищает кэш
	cache.mx.Lock()
	defer cache.mx.Unlock()
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}
