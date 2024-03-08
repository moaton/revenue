package postgres

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	_defaultMaxIdleConns = 2
	_defaultMaxOpenConns = 0
	_defaultLogMode      = logger.Info
)

type Gorm struct {
	*gorm.DB
}

type config struct {
	maxIdleConns   int
	maxOpenConns   int
	nowFunc        func() time.Time
	translateError bool
	logMode        logger.LogLevel
}

func New(connString string, opts ...Option) (*Gorm, error) {
	defaultNowFunc := func() time.Time { return time.Now().UTC() }
	cfg := &config{
		maxIdleConns: _defaultMaxIdleConns,
		maxOpenConns: _defaultMaxOpenConns,
		logMode:      _defaultLogMode,
		nowFunc:      defaultNowFunc,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	db, err := gorm.Open(postgres.Open(connString), cfg.toGormConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to init db session %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve sqlDB %w", err)
	}
	sqlDB.SetMaxIdleConns(cfg.maxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.maxOpenConns)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db %w", err)
	}
	return &Gorm{db}, nil
}
