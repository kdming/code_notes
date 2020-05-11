package main

import (
	cache "codes/cache/go-cache"
	"fmt"
	"time"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	// 默认5分钟过期，10分钟清理一次
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", time.Second*2)

	c.Set("baz", 42, cache.NoExpiration) // 其实就是设置超时为-1

	c.Delete("foo")

	time.Sleep(time.Second * 3)

	v, t := c.Get("foo")
	fmt.Println(v, "v", t, "t")

	select {} // 阻塞
}
