// 代码生成时间: 2025-08-16 07:17:50
@Author: Your Name
@Date: 2023-11-02
*/

package main

import (
    "encoding/json"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/cache"
    "time"
)

// CachePolicy defines the structure for the cache policy
type CachePolicy struct {
    cache *cache.Cache
}

// NewCachePolicy creates a new cache policy with the specified cache lifecycle
func NewCachePolicy(life time.Duration) *CachePolicy {
    // Create a new cache with the specified lifecycle
    bcache, err := cache.NewCache("memory", 
        "{\"interval\":60,\"lazy\":false,\"capacity\":10000,\"duration\": 60,\"timeout\": 1000,\"idle\": 60}