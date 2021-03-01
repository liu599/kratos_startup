package dao

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/log"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"small-service/internal/model"
)

func NewDB() (db *sql.DB, cf func(), err error) {
	var (
		cfg sql.Config
		ct paladin.TOML
	)
	if err = paladin.Get("my-db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db = sql.NewMySQL(&cfg)
	cf = func() {db.Close()}
	return
}

func (d *dao) RawCardInfo(ctx context.Context, wid int) (art *model.CardInfo, err error) {
	// get data from db
	var arc model.CardInfo

	err = d.db.QueryRow(ctx, "SELECT * FROM raw WHERE wid = ?", wid).Scan(
		&arc.Wid, &arc.Wsid, &arc.CardName, &arc.CardCat,
		&arc.Color, &arc.Prop, &arc.Rare, &arc.Level, &arc.Cost,
		&arc.Judge, &arc.Soul, &arc.Attack, &arc.Series,
		&arc.Des1, &arc.Des2, &arc.Des3, &arc.Cover1, &arc.Cover2,
		&arc.Cover3, &arc.Rel1, &arc.Rel2, &arc.Cat)

	fmt.Println(arc.CardName, "   asdfasdfkjasldjflkasjdf")
	if err != nil && err != sql.ErrNoRows {
		log.Error("d.GetDemo.Query error(%v)", err)
		return nil, err
	}
	return &arc, nil
}

// Controller
// GetDemo
func (d *dao) GetDemo(c context.Context, did int) (demo model.CardInfo, err error) {
	err = d.db.QueryRow(c, "SELECT card_name FROM raw WHERE wid = ?", did).Scan(&demo)
	if err != nil && err != sql.ErrNoRows {
		log.Error("d.GetDemo.Query error(%v)", err)
		return
	}
	return demo, nil
}