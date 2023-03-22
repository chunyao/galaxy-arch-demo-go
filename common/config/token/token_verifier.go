package token

import (
	"github.com/gin-gonic/gin"
	"mabang-arch-demo-go/common/config/cache"
	"mabang-arch-demo-go/dto"
	"strings"
)

func TokenVerify(c *gin.Context) {
	request := c.Request
	// 过滤不用token校验的url
	if noTokenVerify(TokenCfg.IgnorePaths, request.RequestURI) {
		return
	}

	// 获取token
	tokenStr := request.Header.Get("token")
	if len(tokenStr) == 0 {
		panic(NewTokenError(dto.Unauthorized, dto.GetResultMsg(dto.Unauthorized)))
	}

	if _, err := cache.BigCache.Get(tokenStr); err != nil {
		panic(NewTokenError(dto.Unauthorized, dto.GetResultMsg(dto.Unauthorized)))
	}

	c.Next()
}

// noTokenVerify 判断url是否不需要token校验
func noTokenVerify(ignorePaths []string, path string) bool {
	// 查询缓存
	if noVerify, err := cache.BigCache.Get(path); err == nil {
		return noVerify.(bool)
	}
	// 匹配url
	for _, ignorePath := range ignorePaths {
		// 路径尾通配符*过滤
		if strings.LastIndex(ignorePath, "*") == len(ignorePath)-1 {
			ignorePath = strings.Split(ignorePath, "*")[0]
			if endIndex := strings.LastIndex(path, "/"); strings.Compare(path[0:endIndex+1], ignorePath) == 0 {
				// 添加缓存
				cache.BigCache.Set(path, true)
				return true
			}
			// 无通配符*过滤
		} else if strings.Compare(path, ignorePath) == 0 {
			// 添加缓存
			cache.BigCache.Set(path, true)
			return true
		}
	}
	return false
}
