package authbox

import (
	"github.com/ak-the-noob-dev/authbox/authbox/ctx"
	"github.com/ak-the-noob-dev/authbox/handlers"
	"github.com/ak-the-noob-dev/authbox/models"
	"github.com/gin-gonic/gin"
)


var (
	appCtx *ctx.AppContext
)

func Init(cfg *Config) error {
	appCtx = &ctx.AppContext{
		DB:              cfg.DB,
		JWTSecret:       cfg.JWTSecret,
		AccessTokenTTL:  cfg.AccessTokenTTL,
		RefreshTokenTTL: cfg.RefreshTokenTTL,
	}

	// Migrate models
	if err := cfg.DB.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	// Register routes
	registerAuthRoutes(cfg.Router)

	return nil
}

func registerAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register(appCtx))
		auth.POST("/login", handlers.Login(appCtx))
	}
}
