package cache

import (
	"api/internal/consts"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type Token struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Key   string
}

// loginId 登录用户ID
// sceneCode 场景标识。注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneCode规避
func NewToken(ctx context.Context, loginId uint, sceneCode ...string) *Token {
	//可以做分库逻辑
	redis := g.Redis()
	return &Token{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(consts.CacheTokenFormat, utils.GetSceneCode(ctx, sceneCode...), loginId),
	}
}

func (cacheThis *Token) Set(value string, ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, value, ttl)
	return
}

func (cacheThis *Token) Get() (value string, err error) {
	valueTmp, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
