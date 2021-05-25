package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/winnix/go-apaleo-one-boilerplate/api/config"
	"github.com/winnix/go-apaleo-one-boilerplate/api/web"
)

func ApaleoAuthRequired(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := web.GetCurrentUser(c)
		if err != nil {
			redirectURL := fmt.Sprintf("%s%s", cfg.Host, c.Request.URL.String())
			web.StoreInSession(c, web.RedirectURLSessionKey, redirectURL)

			c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/auth/apaleo", cfg.Host))
			c.AbortWithStatus(http.StatusTemporaryRedirect)
		}

		c.Next()
	}
}
