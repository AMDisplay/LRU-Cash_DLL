package main

type CacheInterface interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type Key string

type Cache struct {
	cap   int
	queue *List
	dict  map[Key]*ListItem
}

// Добавить значение в кэш по ключу.
func (c *Cache) Set(key Key, value interface{}) bool {
	if _, ok := c.dict[key]; ok {
		Item := *c.dict[key]
		Item.data = value
		c.queue.MoveToFront(&Item)
		return true
	}
	// Если очередь больше кеша
	if c.cap < c.queue.len+1 {
		lastItem := c.queue.Back()
		c.queue.Remove(lastItem)
	}
	newItem := &ListItem{
		next: nil,
		prev: nil,
		data: value,
		list: c.queue,
	}
	c.dict[key] = newItem
	c.queue.PushFront(newItem)
	return false

}

// Получить значение из кэша по ключу.
func (c *Cache) Get(key Key) (interface{}, bool) {
	if _, ok := c.dict[key]; ok {
		Item := c.dict[key]
		c.queue.PushFront(Item)
		return Item, true
	}
	return nil, false

}

// Очистить кэш.  O(1)
func (c *Cache) Clear() {
	newCahce := &Cache{
		cap:   c.cap,
		queue: &List{},
		dict:  make(map[Key]*ListItem),
	}
	*c = *newCahce
}
