package fast

type fastCache struct {
	shards    []*cacheShard
	sharkMask uint64
	hash      fnv64
}

func NewFastCache(maxEntries, shardsNum int, onEvicted func(key string, value interface{})) *fastCache {
	fastCache := &fastCache{
		hash:      newDefaultHasher(),
		shards:    make([]*cacheShard, shardsNum),
		sharkMask: uint64(shardsNum - 1),
	}
	for i := 0; i < shardsNum; i++ {
		fastCache.shards[i] = newCacheShard(maxEntries, onEvicted)
	}
	return fastCache
}

func (c *fastCache) getShard(key string) *cacheShard {
	hashKey := c.hash.Sum64(key)
	return c.shards[hashKey&c.sharkMask]
}

func (c *fastCache) Set(key string, value interface{}) {
	c.getShard(key).set(key, value)
}

func (c *fastCache) Get(key string) interface{} {
	return c.getShard(key).get(key)
}

func (c *fastCache) Del(key string) {
	c.getShard(key).del(key)
}

func (c *fastCache) Len() int {
	length := 0
	for _, shard := range c.shards {
		length += shard.len()
	}
	return length
}
