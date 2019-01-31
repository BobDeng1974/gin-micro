package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 结构体
type Control struct {

}

// 响应成功
func (c *Control) Success( context *gin.Context,message string) {
	context.JSON(http.StatusOK, JsonObject{
		Code:    "0",
		Message: message,
		Content:nil,
	})
	context.Abort()
	return
}

// 返回状态和结果
func (c *Control) SuccessContent( context *gin.Context,message string,content interface{}) {
	context.JSON(http.StatusOK, JsonObject{
		Code:    "0",
		Message: message,
		Content: content,
	})
	context.Abort()
	return
}

// 无权限的访问
func (c *Control) RefusedError( context *gin.Context) {
	context.JSON(http.StatusUnauthorized, JsonObject{
		Code:    "401",
		Message: "Unauthorized",
		Content: nil,
	})
	context.Abort()
	return
}

// 无权限的访问
func (c *Control) BindingError( context *gin.Context) {
	context.JSON(http.StatusOK, JsonObject{
		Code:    "404",
		Message: "Binding Error",
		Content: nil,
	})
	context.Abort()
	return
}

// 服务器内部错误
func (c *Control) InternalError( context *gin.Context,err string) {
	context.JSON(http.StatusInternalServerError, JsonObject{
		Code:    "500",
		Message: err,
		Content: nil,
	})
	context.Abort()
	return
}

// 通用的统一返回结果
func (c *Control) ReturnResult( context *gin.Context,httpStatus int,code string,message string,content interface{}) {
	context.JSON(http.StatusOK, JsonObject{
		Code:    "0",
		Message: message,
		Content: content,
	})
	context.Abort()
	return
}






