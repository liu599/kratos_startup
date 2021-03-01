// Code generated by kratos tool genmc. DO NOT EDIT.

/*
  Package dao is a generated mc cache package.
  It is generated from:
  type _mc interface {
		// mc: -key=keyArt -type=get
		CacheCardInfo(c context.Context, id int) (*model.CardInfo, error)
		// mc: -key=keyArt -expire=d.demoExpire
		AddCacheCardInfo(c context.Context, id int, art *model.CardInfo) (err error)
		// mc: -key=keyArt
		DeleteCardInfo(c context.Context, id int) (err error)
	}
*/

package dao

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/memcache"
	"github.com/go-kratos/kratos/pkg/log"
	"small-service/internal/model"
)

var _ _mc

// CacheCardInfo get data from mc
func (d *dao) CacheCardInfo(c context.Context, id int) (res *model.CardInfo, err error) {
	key := keyArt(id)
	fmt.Println(id, "create-Id")
	res = &model.CardInfo{}
	if err = d.mc.Get(c, key).Scan(res); err != nil {
		res = nil
		if err == memcache.ErrNotFound {
			err = nil
		}
	}
	if err != nil {
		log.Errorv(c, log.KV("CacheCardInfo", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheCardInfo Set data to mc
func (d *dao) AddCacheCardInfo(c context.Context, id int, val *model.CardInfo) (err error) {
	if val == nil {
		return
	}
	key := keyArt(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheCardInfo", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DeleteCardInfo delete data from mc
func (d *dao) DeleteCardInfo(c context.Context, id int) (err error) {
	key := keyArt(id)
	if err = d.mc.Delete(c, key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("DeleteCardInfo", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}
