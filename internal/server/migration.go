package server

import (
	"context"
	"go-xianyu/internal/model"
	"go-xianyu/pkg/log"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Migrate struct {
	db  *gorm.DB
	log *log.Logger
}

func NewMigrate(db *gorm.DB, log *log.Logger) *Migrate {
	return &Migrate{
		db:  db,
		log: log,
	}
}
func (m *Migrate) Start(ctx context.Context) error {
	if err := m.db.AutoMigrate(
		// 新增Model之后需要在这增加 Model结构
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.Message{},
	); err != nil {
		m.log.Error("user migrate error", zap.Error(err))
		return err
	}
	m.log.Info("AutoMigrate success")
	os.Exit(0)
	return nil
}
func (m *Migrate) Stop(ctx context.Context) error {
	m.log.Info("AutoMigrate stop")
	return nil
}
