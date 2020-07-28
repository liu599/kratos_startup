// Code generated by kratos tool genbts. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type Dao interface {
		Close()
		Ping(ctx context.Context) (err error)
		// bts: -nullcache=&model.Article{RegId:-1} -check_null_code=$!=nil&&$.RegId==-1
		Article(c context.Context, regId int64) (*model.Article, error)
	}
*/

package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache"
	"small-service/internal/model"
)

// Article get data from cache if miss will call source method, then add to cache.
func (d *dao) Article(c context.Context, regId int64) (res *model.Article, err error) {
	addCache := true
	res, err = d.CacheArticle(c, regId)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.RegId == -1 {
			res = nil
		}
	}()
	if res != nil {
		cache.MetricHits.Inc("bts:Article")
		return
	}
	cache.MetricMisses.Inc("bts:Article")
	res, err = d.RawArticle(c, regId)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.Article{RegId: -1}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheArticle(c, regId, miss)
	})
	return
}