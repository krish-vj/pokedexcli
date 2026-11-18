package pokecache

import (
	"sync"
	"time"
)
type CacheEntry struct{
	createdAt time.Time
	val []byte

}
type Cache struct{
	mtx sync.Mutex
	mp map[string]CacheEntry
}
func (c *Cache) Add(key string, valb[] byte){
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.mp[key]=CacheEntry{createdAt: time.Now(), val: valb}
	
}
func (c *Cache) Get(key string) ([]byte, bool){
	c.mtx.Lock()
	defer c.mtx.Unlock()
	cachStr, ok:=c.mp[key]
	if (ok){
		return cachStr.val, true;
	}
	return nil, false;
	
}
func reapLoop(interval time.Duration, c *Cache){
	ticker:=time.NewTicker(interval);
	defer ticker.Stop();
	
		
			for range ticker.C{
				c.mtx.Lock()
				
				for key, entry:= range c.mp{
					if (time.Since(entry.createdAt)>interval){
						delete(c.mp, key)
					}
				}
				c.mtx.Unlock()
			}
		
	
}

func NewCache(interval time.Duration) *Cache{
	rt:= &Cache{
		mp: make(map[string]CacheEntry),
	};
	go reapLoop(interval, rt);
	return rt
}

