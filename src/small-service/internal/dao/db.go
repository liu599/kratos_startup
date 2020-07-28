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

func (d *dao) RawArticle(ctx context.Context, regId int64) (art *model.Article, err error) {
	// get data from db
	var arc model.Article
	err = d.db.QueryRow(ctx, "SELECT * FROM regression WHERE reg_id = ?", regId).Scan(&arc.RegId, &arc.RegName, &arc.Author)
	fmt.Println(arc)
	if err != nil && err != sql.ErrNoRows {
		log.Error("d.GetDemo.Query error(%v)", err)
		return nil, err
	}
	art = &arc
	//art.RegName = arc.RegName
	//art.Author = arc.Author
	//art.RegId = arc.RegId
	return art, nil
}

// Controller
// GetDemo
func (d *dao) GetDemo(c context.Context, did int64) (demo model.Regression, err error) {
	err = d.db.QueryRow(c, "SELECT reg_name FROM regression WHERE reg_id = ?", did).Scan(&demo)
	if err != nil && err != sql.ErrNoRows {
		log.Error("d.GetDemo.Query error(%v)", err)
		return
	}
	return demo, nil
}