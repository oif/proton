package core

import (
	"github.com/coocood/freecache"
)

const PROTON_LOGO = `
		     _____
_______________________  /_____________
___  __ \_  ___/  __ \  __/  __ \_  __ \
__  /_/ /  /   / /_/ / /_ / /_/ /  / / /
_  .___//_/    \____/\__/ \____//_/ /_/
/_/


` // Logo

const CACHE_KEY_FORMAT = "%s_%s_%s" // name, type, subnet

var statistics *ProtonStat // 统计数据
var cache *freecache.Cache // 解析缓存
