package middleware

import (
	"github.com/gin-gonic/gin"
	cfg "github.com/winnix/go-apaleo-one-boilerplate/api/config"
)

func Config(config *cfg.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}
