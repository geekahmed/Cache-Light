package main

import (
	src "cache-light/src/CacheImpl"
	"fmt"
)

func main()  {
	cacheForever := src.NewCacheForever()
	cacheForever.Set(1, "Ahmed")
	cacheForever.Set(2, "Moustafa")

	fmt.Println(cacheForever.Get(1))
	fmt.Println(cacheForever.Get(2))

	fmt.Println(cacheForever.Size())

	cacheForever.Remove(1)
	fmt.Println(cacheForever.Get(1))
	cacheForever.Clear()
	fmt.Println(cacheForever.Size())
}
