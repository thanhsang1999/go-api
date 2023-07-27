package appctx

import (
	"go-api/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db *gorm.DB
	ps pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, ps pubsub.Pubsub) *appCtx {
	return &appCtx{db: db, ps: ps}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.ps
}
