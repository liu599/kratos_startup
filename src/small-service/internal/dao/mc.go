package dao

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/cache/memcache"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"small-service/internal/model"
)

//go:generate kratos tool genmc
type _mc interface {
	// mc: -key=keyArt -type=get
	CacheCardInfo(c context.Context, id int) (*model.CardInfo, error)
	// mc: -key=keyArt -expire=d.demoExpire
	AddCacheCardInfo(c context.Context, id int, art *model.CardInfo) (err error)
	// mc: -key=keyArt
	DeleteCardInfo(c context.Context, id int) (err error)
}

func NewMC() (mc *memcache.Memcache, cf func(), err error) {
	var (
		cfg memcache.Config
		ct paladin.TOML
	)
	if err = paladin.Get("memcache.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	mc =  memcache.New(&cfg)
	cf = func() {mc.Close()}
	return
}

func (d *dao) PingMC(ctx context.Context) (err error) {
	if err = d.mc.Set(ctx, &memcache.Item{Key: "ping", Value: []byte("pong"), Expiration: 0}); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

func keyArt(id int) string {
	return fmt.Sprintf("art_%d", id)
}
