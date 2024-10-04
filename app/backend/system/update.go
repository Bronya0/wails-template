package system

import (
	"app/backend/common"
	"app/backend/types"
	"context"
	"github.com/go-resty/resty/v2"
)

type Update struct {
	ctx context.Context
}

func (obj *Update) Start(ctx context.Context) {
	obj.ctx = ctx
}
func (obj *Update) CheckUpdate() *types.Tag {
	client := resty.New()
	tag := &types.Tag{}
	_, _ = client.R().SetResult(tag).Get(common.UPDATE_URL)
	return tag
}
