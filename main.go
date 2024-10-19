package cache

type Cache struct {
	data map[string]cacheValue
}

type cacheValue struct {
	value interface{}
}

func New() *Cache {
	return &Cache{
		data: make(map[string]cacheValue),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = cacheValue{
		value: value,
	}
}

func (c *Cache) Get(key string) interface{} {
	value := c.data[key]
	return value.value
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
