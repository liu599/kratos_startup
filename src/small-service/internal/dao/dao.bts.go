// Code generated by kratos tool genbts. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type Dao interface {
		Close()
		Ping(ctx context.Context) (err error)
		// bts: -nullcache=&model.CardInfo{Wid:1} -check_null_code=$!=nil&&$.Wid==1
		CardInfo(c context.Context, wid int) (*model.CardInfo, error)
	}
*/

package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache"
	"small-service/internal/model"
)

// CardInfo get data from cache if miss will call source method, then add to cache.
func (d *dao) CardInfo(c context.Context, wid int) (res *model.CardInfo, err error) {
	addCache := true
	res, err = d.CacheCardInfo(c, wid)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.Wid == 1 {
			res = nil
		}
	}()
	if res != nil {
		cache.MetricHits.Inc("bts:CardInfo")
		return
	}
	cache.MetricMisses.Inc("bts:CardInfo")
	res, err = d.RawCardInfo(c, wid)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.CardInfo{Wid: 1}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheCardInfo(c, wid, miss)
	})
	return
}
