package cors

import (
	"net/http"

	staticContentType "ioyoa/static/shared/contentType"
	staticSymbols "ioyoa/static/shared/symbols"

	"github.com/gin-contrib/cors"
)

func BasicCorsConfig() cors.Config {
    return cors.Config{
        AllowOrigins:     []string{staticSymbols.Asterisk},
        AllowMethods:     []string{http.MethodGet, http.MethodPost},
        AllowHeaders:     []string{staticContentType.ContentType},
        AllowCredentials: true,
    }
}