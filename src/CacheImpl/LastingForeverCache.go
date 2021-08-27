package src

type CacheForever struct {
	cache map[interface{}]interface{}
}

func NewCacheForever() *CacheForever{
	var cacheForever CacheForever
	cacheForever.cache = make(map[interface{}]interface{})
	return &cacheForever
}

func (c *CacheForever) Size() int {
	return len(c.cache)
}
func (c *CacheForever) Set (key interface{}, value interface{})  {
	c.cache[key] = value
}
func (c *CacheForever) Get (key interface{}) interface{}  {
	return c.cache[key]
}

func (c *CacheForever) Remove (key interface{}){
	delete(c.cache, key)
}
func (c *CacheForever) Clear (){
	c.cache = make(map[interface{}]interface{})
}



