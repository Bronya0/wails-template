package system

import (
	"app/backend/common"
	"context"
	"github.com/go-resty/resty/v2"
)

type Update struct {
	ctx context.Context
}
type Tag struct {
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
}

func (obj *Update) Start(ctx context.Context) {
	obj.ctx = ctx
}
func (obj *Update) CheckUpdate() *Tag {
	client := resty.New()
	tag := &Tag{}
	_, _ = client.R().SetResult(tag).Get(common.UPDATE_URL)
	return tag
}
