package ctx

import (
	"gorm.io/gorm"
)

type AppContext struct {
	DB            *gorm.DB
	JWTSecret     string
	AccessTokenTTL  int64
	RefreshTokenTTL int64
}
