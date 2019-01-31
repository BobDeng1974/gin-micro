package auth

import (
	"gin-micro/helpers/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// token 校验中间件
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		if strings.Contains(path, "swagger") {
			return
		}
		access_token := context.Request.Header.Get("ACCESS_TOKEN")
		if access_token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  -1,
				"message": "请求未携带token,无访问权限！",
			})
			context.Abort()
			return
		}
		j := token.NewJWT()
		// 解析token包含的信息
		claims, err := j.ResolveToken(access_token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  -1,
				"message": err.Error(),
			})
			context.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		context.Set("claims", claims)
	}
}

