package core

import (
	"github.com/coocood/freecache"
)

// ProtonLOGO console output when starting
const ProtonLOGO = `
		     _____
_______________________  /_____________
___  __ \_  ___/  __ \  __/  __ \_  __ \
__  /_/ /  /   / /_/ / /_ / /_/ /  / / /
_  .___//_/    \____/\__/ \____//_/ /_/
/_/


`

// CacheKeyFormat the format of cache store key
const CacheKeyFormat = "%s_%s_%s" // name, type, subnet

// servicePublicIP for server internal IP address
var servicePublicIP string

// statistics instance of statistics service
var statistics *ProtonStat // 统计数据
// cache instance of cache service
var cache *freecache.Cache // 解析缓存
