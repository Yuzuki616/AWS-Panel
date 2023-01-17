package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var client = cache.New(6*time.Hour, 2*time.Hour)

func Get(key string) (interface{}, bool) {
	return client.Get(key)
}

func Set(key string, value interface{}, expiration time.Duration) {
	client.Set(key, value, expiration)
}

func Delete(key string) {
	client.Delete(key)
}

func DeleteExpired() {
	client.DeleteExpired()
}

func DeleteAll() {
	client.Flush()
}
