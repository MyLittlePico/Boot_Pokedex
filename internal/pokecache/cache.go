package pokecache

import (
	"time"
	"sync"
)


type Cache struct{
	cmap	map[string]CacheEntry
	mu     *sync.Mutex
}
type CacheEntry struct{
	createdAt time.Time
	val []byte
}


func NewCache(interval time.Duration) Cache{
	c := Cache{
		cmap : make(map[string]CacheEntry, 0),
		mu : &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cmap [key] = CacheEntry{
		createdAt : time.Now(),
		val : val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()

	val , ok := c.cmap [key]
	if !ok{
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for{
		<-ticker.C
		for key, val := range c.cmap{
			
			
			if time.Since(val.createdAt) > interval{
				c.mu.Lock()
				delete (c.cmap, key)
				c.mu.Unlock()
			}
		} 
	}
}