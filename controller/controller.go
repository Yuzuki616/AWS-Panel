package controller

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var cacheClient = cache.New(6*time.Hour, 6*time.Hour)
