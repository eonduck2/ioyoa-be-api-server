package httpmethod

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GinMethodHandler는 Gin 라우터와 HTTP 메서드, 경로, 핸들러를 받아 라우팅을 설정
//
// @params router, method, path, handler
// 지원되지 않는 메서드에 대해서는 핸들러를 등록하지 않고 함수가 종료
func GinMethodHandler(router *gin.Engine, method string, path string, handler gin.HandlerFunc) {
	method = strings.ToUpper(method)

	switch method {
	case http.MethodGet:
		router.GET(path, handler)
	case http.MethodPost:
		router.POST(path, handler)
	case http.MethodPut:
		router.PUT(path, handler)
	case http.MethodDelete:
		router.DELETE(path, handler)
	case http.MethodPatch:
		router.PATCH(path, handler)
	case http.MethodOptions:
		router.OPTIONS(path, handler)
	case http.MethodHead:
		router.HEAD(path, handler)
	default:
		return
	}
}
