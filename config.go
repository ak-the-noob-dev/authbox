package authbox

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EmailSender is an interface for sending emails (used for verification, etc.)
type EmailSender interface {
	Send(to string, subject string, body string) error
}

// Config holds all the settings for AuthBox
type Config struct {
	Router            *gin.Engine
	DB                *gorm.DB
	JWTSecret         string
	AccessTokenTTL    int    // in minutes
	RefreshTokenTTL   int    // in minutes
	AppBaseURL        string // used for email links
	EmailSender       EmailSender
	EnableRateLimit   bool
	VerificationEmail bool
}
