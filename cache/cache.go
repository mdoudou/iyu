package cache

import (
	"github.com/lhlyu/iyu/common"
)

/**
LHLYU-BLOG:AUTHOR  -  存放作者信息【string】
LHLYU-BLOG:CATALOG   -   分类【Hash】
LHLYU-BLOG:TAGS    -   标签【Hash】
LHLYU-BLOG:NAIL    -   钉子【Hash】


LHLYU-BLOG:ARTICLE:LIST - 文章列表【list】
-- LHLYU-BLOG:ARTICLE:LIST:KEY  【string】 =》 LHLYU-BLOG:ARTICLE:LIST

LHLYU-BLOG:ARTICLE:MAP - 文章MAP【hash】
- field: 文章ID

LHLYU-BLOG:ARTICLE:IVEAW:id  - 文章浏览量【string】

LHLYU-BLOG:IVEAW      - 全站浏览量【string】
-- LHLYU-BLOG:IVEAW:KEY  【string】 =》 LHLYU-BLOG:IVEAW
*/

type Cache struct {
}

func (*Cache) hasRedis() bool {
	if common.Redis == nil {
		return false
	}
	return true
}

func (c *Cache) AddCatalogList(key string, v []interface{}) {
	if c.hasRedis() {
		common.Redis.RPush(key, v...)
	}

}