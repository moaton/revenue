package postgres

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Option func(*config)

func TranslateError(enable bool) Option {
	return func(c *config) {
		c.translateError = enable
	}
}

func NowFunc(f func() time.Time) Option {
	return func(c *config) {
		c.nowFunc = f
	}
}

func MaxIdleConns(n int) Option {
	return func(c *config) {
		c.maxIdleConns = n
	}
}

func MaxOpenConns(n int) Option {
	return func(c *config) {
		c.maxOpenConns = n
	}
}

func (c *config) toGormConfig() *gorm.Config {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  c.logMode,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)
	return &gorm.Config{
		NowFunc:        c.nowFunc,
		TranslateError: c.translateError,
		Logger:         newLogger,
	}
}
