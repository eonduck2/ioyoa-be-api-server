package httpmethod

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GinMethodHandler는 Gin 라우터와 HTTP 메서드, 경로, 핸들러를 받아 라우팅을 설정합니다.
//
// @param router Gin 라우터 인스턴스입니다.
// @param method HTTP 메서드를 나타내는 문자열입니다. (예: "GET", "POST", "PUT" 등)
// @param path 요청 경로를 나타내는 문자열입니다.
// @param handler 요청을  처리하는 핸들러 함수입니다.
//
// 지원되지 않는 메서드에 대해서는 핸들러를 등록하지 않고 함수가 종료됩니다.
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
